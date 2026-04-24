package service

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model"
	"web_doscom/internal/utils"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type BlogService struct {
	DB             *gorm.DB
	BlogModel      *model.BlogModel
	BlogGallery    *model.BlogGalleryModel
	GalleryService *GalleryService
}

func NewBlogService(db *gorm.DB, n *model.BlogModel, m *model.BlogGalleryModel, g *GalleryService) *BlogService {
	return &BlogService{
		DB:             db,
		BlogModel:      n,
		BlogGallery:    m,
		GalleryService: g,
	}
}

func (m *BlogService) ProcessBlogGalleries(
	ctx context.Context,
	newImages []*multipart.FileHeader,
	blogDetail *model.BlogPayload,
) ([]*model.GalleryResponse, []int, error) {

	var galleryIDS []int

	if len(blogDetail.ExistingID) > 0 {
		isValid, err := m.GalleryService.Model.CheckExistingGallery(blogDetail.ExistingID)
		if err != nil {
			return nil, nil, fmt.Errorf("failedt to check existing gallery %w", err)
		}

		if !isValid {
			return nil, nil, fmt.Errorf("invalid gallery ids")
		}

		for _, id := range blogDetail.ExistingID {
			if id != nil {
				galleryIDS = append(galleryIDS, *id)
			}
		}

	}

	const (
		maxGallery  = 5
		maxKategori = 3
	)
	totalGallery := len(galleryIDS) + len(newImages)
	if totalGallery > maxGallery {
		return nil, nil, fmt.Errorf("you can only use %d gallery", maxGallery)
	}
	if len(blogDetail.Kategori) > maxKategori {
		return nil, nil, fmt.Errorf("you can only tag %d kategori in one blog", maxKategori)
	}

	now := time.Now()
	galleryName := "foto untuk blog dengan judul " + blogDetail.Title
	galleryData := &model.GalleryInsert{
		IDUsers:     blogDetail.AuthorID,
		GalleryName: galleryName,
		GalleryType: "blog",
		Description: "foto untuk kepentingan blog",
		EventDate: time.Date(
			now.Year(),
			now.Month(),
			now.Day(), 0, 0, 0, 0, time.UTC,
		),
	}

	fileUpload := make([]*model.UploadFileRequest, len(newImages))
	for i, file := range newImages {
		fileContent, err := file.Open()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file")
		}

		fileUpload[i] = &model.UploadFileRequest{
			FileHeader: file,
			File:       fileContent,
			Folder:     "blog",
			UserID:     uint(blogDetail.AuthorID),
		}
	}

	defer func() {
		for _, file := range fileUpload {
			file.File.Close()
		}
	}()

	var newGalleryDataResponse []*model.GalleryResponse
	if len(newImages) > 0 {
		var err error
		newGalleryDataResponse, err = m.GalleryService.UploadAndInsertGalleryMultiple(
			ctx,
			galleryData,
			fileUpload,
		)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to upload new image and gallery %w", err)
		}
		for _, gallery := range newGalleryDataResponse {
			galleryIDS = append(galleryIDS, gallery.ID)
		}
	}

	return newGalleryDataResponse, galleryIDS, nil
}

func (m *BlogService) CreateBlogImage(
	ctx context.Context,
	blogDetail *model.BlogPayload,
	newImages []*multipart.FileHeader,
	userRole string,
) (*model.BlogResponse, error) {

	if err := utils.CheckRolePermission(
		userRole,
		constants.RoleAdmin,
		constants.RoleKeyKoorMedcrev,
	); err != nil {
		return nil, fmt.Errorf("you are not allowed to access this resource %w", err)
	}
	if blogDetail.Status == "" {
		blogDetail.Status = constants.StatusDraft
	}
	if userRole == constants.RoleKeyKoorMedcrev &&
		(blogDetail.Status != constants.StatusDraft || blogDetail.Status != constants.StatusPending) {
		return nil, fmt.Errorf("you are can only set status to draft or pending")
	}

	newGalleryDataResponse, allGalleryIDS, err := m.ProcessBlogGalleries(
		ctx,
		newImages,
		blogDetail,
	)
	if err != nil {
		return nil, fmt.Errorf("failed while process gallery %w", err)
	}

	newGalleryIDS := make([]int, 0, len(newGalleryDataResponse))
	for _, idGallery := range newGalleryDataResponse {
		newGalleryIDS = append(newGalleryIDS, idGallery.ID)
	}

	var blogDataResponse model.BlogResponse
	err = m.DB.Transaction(func(tx *gorm.DB) error {

		var txFailed bool
		defer func() {
			if txFailed {
				m.GalleryService.DeleteGalleryMultiple(ctx, newGalleryIDS)
			}
		}()

		modelBlog := m.BlogModel.WithTx(tx)
		modelBlogGallery := m.BlogGallery.WithTx(tx)
		var err error

		var thumbnailURL string
		if len(newGalleryDataResponse) > 0 {
			thumbnailURL = newGalleryDataResponse[0].FileURL
		}

		blog := &model.Blog{
			AuthorID:     blogDetail.AuthorID,
			Title:        blogDetail.Title,
			Slug:         blogDetail.Slug,
			Content:      blogDetail.Content,
			Kategori:     blogDetail.Kategori,
			ThumbnailURL: thumbnailURL,
			Status:       blogDetail.Status,
			PublishedAt:  nil,
		}

		if err := modelBlog.InsertBlog(ctx, blog); err != nil {
			txFailed = true
			log.Printf("Failed to insert blog %v: %v", blog, err)
			return fmt.Errorf("failed to insert blog %w", err)
		}

		blogGallery := make([]*model.BlogGallery, len(allGalleryIDS))
		for i, id := range blogDetail.ExistingID {
			if id == nil {
				txFailed = true
				return fmt.Errorf("gallery id is nil/empty, issue on backend side")
			}
			blogGallery[i] = &model.BlogGallery{
				BlogID:    blog.ID,
				GalleryID: *id,
			}
		}

		blogGalleryResponse, err := modelBlogGallery.InsertBlogGalleryMultiple(blogGallery)
		if err != nil {
			txFailed = true
			log.Printf("Failed to insert blog gallery %v: %v", blogGallery, err)
			return fmt.Errorf("failed to insert blog gallery %w", err)
		}

		blogDataResponse = model.BlogResponse{
			ID:           blog.ID,
			AuthorID:     blog.AuthorID,
			Title:        blog.Title,
			Slug:         blog.Slug,
			Content:      blog.Content,
			Kategori:     blog.Kategori,
			ThumbnailURL: blog.ThumbnailURL,
			PublishedAt:  blog.PublishedAt,
			BlogImage:    blogGalleryResponse,
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert database and Transaction failed %w", err)
	}

	return &blogDataResponse, nil
}

func (m *BlogService) GetAllBlogsForAdmin(ctx context.Context, kategori []string, offset, limit int) ([]model.BlogThumbnail, int, error) {
	// validasi kategory max 3
	if len(kategori) > 3 {
		return nil, 0, fmt.Errorf("max filter is 3 kategori")
	}
	blogs, totalData, err := m.BlogModel.GetBlogs(ctx, kategori, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return blogs, int(totalData), nil
}

func (m *BlogService) GetAllBlogs(ctx context.Context, page, limit int) ([]model.BlogThumbnail, int, error) {

	offset := (page - 1) * limit

	blogs, totalData, err := m.BlogModel.GetAllBlogs(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return blogs, totalData, nil
}

func (m *BlogService) GetBlogByID(id int) (*model.BlogResponse, error) {

	// get blog by id
	blog, err := m.BlogModel.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (m *BlogService) UpdateBlog(
	ctx context.Context,
	idBlog, authorID int,
	blogDetail *model.BlogPatch,
	newImages []*multipart.FileHeader,
) (*model.BlogResponse, error) {
	var GalleryIDS []*int
	// check if there is existing id image
	if len(blogDetail.ExistingID) > 0 {
		for _, id := range blogDetail.ExistingID {
			GalleryIDS = append(GalleryIDS, id)
		}

		// check if existing id gallery is valid
		isValid, err := m.GalleryModel.Model.CheckExistingGallery(GalleryIDS)
		if err != nil {
			return nil, err
		}
		if !isValid {
			return nil, fmt.Errorf("invalid gallery ids")
		}
	}

	// check len of new images and old images
	allImagesID := len(GalleryIDS) + len(newImages)
	if allImagesID > maxGallery {
		return nil, fmt.Errorf("You can only use %d gallery", maxGallery)
	}

	fileUpload := make([]*model.UploadFileRequest, len(newImages))
	for i, file := range newImages {
		fileContent, err := file.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open file")
		}

		fileUpload[i] = &model.UploadFileRequest{
			FileHeader: file,
			File:       fileContent,
			Folder:     "blog",
			UserID:     uint(authorID),
		}
	}

	defer func() {
		for _, file := range fileUpload {
			file.File.Close()
		}
	}()
	// upload foto baru and insert to database
	var gallery []*model.GalleryResponse
	if len(newImages) > 0 {
		now := time.Now()
		galleryName := "foto untuk blog dengan judul " + blogDetail.Title
		galleryData := &model.GalleryInsert{
			IDUsers:     authorID,
			GalleryName: galleryName,
			GalleryType: "blog",
			Description: "foto untuk kepentingan blog",
			EventDate: time.Date(
				now.Year(),
				now.Month(),
				now.Day(), 0, 0, 0, 0, time.UTC,
			),
		}

		var err error
		gallery, err = m.GalleryModel.UploadAndInsertGalleryMultiple(
			ctx,
			galleryData,
			fileUpload,
		)
		if err != nil {
			return nil, err
		}

		for _, gallery := range gallery {
			GalleryIDS = append(GalleryIDS, &gallery.ID)
		}
	}

	if blogDetail.Status == "" {
		blogDetail.Status = "draft"
	}

	// update blog
	var publishedAt *time.Time
	if blogDetail.Status == "draft" {
		publishedAt = nil
	} else if blogDetail.Status == "published" {
		now := time.Now()
		publishedAt = &now
	} else {
		publishedAt = blogDetail.PublishedAt
	}

	oldData := &model.Blog{
		AuthorID:    authorID,
		Title:       blogDetail.Title,
		Slug:        blogDetail.Slug,
		Content:     blogDetail.Content,
		Kategori:    pq.StringArray(blogDetail.Kategori),
		PublishedAt: nil,
		Status:      blogDetail.Status,
		UpdatedAt:   time.Now(),
	}
	updatedBlog := utils.StructToMap(oldData)
	if len(gallery) > 0 {
		updatedBlog["thumbnail_url"] = gallery[0].FileURL
	}

	blogUpdate, err := m.BlogModel.UpdateBlogPartial(idBlog, updatedBlog)
	if err != nil {
		return nil, fmt.Errorf("failed to insert data %w", err)
	}

	// update blog_gallery -> jika terdapat terdapat perubahan foto
	blogGallery, err := m.BlogGallery.UpdateBlogGallery(GalleryIDS, blogUpdate.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update gallery image %w", err)
	}

	result := &model.BlogResponse{
		ID:           blogUpdate.ID,
		AuthorID:     blogUpdate.AuthorID,
		Title:        blogUpdate.Title,
		Slug:         blogUpdate.Slug,
		Content:      blogUpdate.Content,
		Kategori:     blogUpdate.Kategori,
		ThumbnailURL: blogUpdate.ThumbnailURL,
		PublishedAt:  blogUpdate.PublishedAt,
		BlogImage:    blogGallery,
	}

	return result, nil
}

func (m *BlogService) GetBlogByKategori(
	ctx context.Context,
	kategori []string,
	limit, offset int,
) ([]model.BlogThumbnail, int, error) {

	// check whether there is a filter sent or not
	var (
		blog      []model.BlogThumbnail
		totalData int
		err       error
	)
	if len(kategori) == 0 {
		blog, totalData, err = m.BlogModel.GetAllBlogs(ctx, limit, offset)
		if err != nil {
			return nil, 0, err
		}
	} else {
		blog, totalData, err = m.BlogModel.GetBlogsByKategori(
			ctx,
			kategori,
			limit,
			offset,
		)
		if err != nil {
			return nil, 0, err
		}
	}

	return blog, totalData, nil
}

func (m *BlogService) DeleteBlogByID(
	ctx context.Context,
	blogId int,
) error {

	tx := m.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// delete blog_gallery first
	if err := m.BlogGallery.DeleteBlogGalleryByBlogID(ctx, tx, blogId); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete blog gallery: %w", err)
	}

	// delete blog then
	if err := m.BlogModel.DeleteBlog(ctx, tx, blogId); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete blog : %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
