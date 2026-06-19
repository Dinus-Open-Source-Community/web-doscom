package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"time"
	"web_doscom/internal/authorization"
	blogAuthorization "web_doscom/internal/authorization/blog"
	"web_doscom/internal/constants"
	"web_doscom/internal/database/model/dto"
	"web_doscom/internal/database/model/entity"
	"web_doscom/internal/utils"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type BlogService struct {
	DB             *gorm.DB
	BlogModel      *entity.BlogModel
	BlogGallery    *entity.BlogGalleryModel
	GalleryService *GalleryService
}

func NewBlogService(db *gorm.DB, n *entity.BlogModel, m *entity.BlogGalleryModel, g *GalleryService) *BlogService {
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
	blogDetail *dto.BlogPayload,
) ([]*dto.GalleryResponse, []int, error) {

	var allGalleryIDS []int
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
				allGalleryIDS = append(allGalleryIDS, *id)
			}
		}

	}

	const (
		maxGallery  = 5
		maxKategori = 3
	)
	totalGallery := len(allGalleryIDS) + len(newImages)
	if totalGallery > maxGallery {
		return nil, nil, fmt.Errorf("you can only use %d gallery", maxGallery)
	}
	if len(blogDetail.Kategori) > maxKategori {
		return nil, nil, fmt.Errorf("you can only tag %d kategori in one blog", maxKategori)
	}

	// if there is no new image return
	if len(newImages) == 0 {
		return nil, allGalleryIDS, nil
	}

	now := time.Now()
	galleryName := "foto untuk blog dengan judul " + blogDetail.Title
	galleryData := &dto.GalleryInsert{
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

	fileUpload := make([]*dto.UploadFileRequest, len(newImages))
	for i, file := range newImages {
		fileContent, err := file.Open()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open file")
		}

		fileUpload[i] = &dto.UploadFileRequest{
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

	var newGalleryDataResponse []*dto.GalleryResponse
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
			allGalleryIDS = append(allGalleryIDS, gallery.ID)
		}
	}

	return newGalleryDataResponse, allGalleryIDS, nil
}

func (m *BlogService) CreateBlog(
	ctx context.Context,
	blogDetail *dto.BlogPayload,
	newImages []*multipart.FileHeader,
	userRole string,
) (*dto.BlogResponse, error) {

	if err := authorization.CheckRolePermission(
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

	if len(newImages) == 0 && len(blogDetail.ExistingID) == 0 {
		return nil, fmt.Errorf("at least one gallery image is required")
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

	var blogDataResponse dto.BlogResponse
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

		blog := &entity.Blog{
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

		blogGallery := make([]*entity.BlogGallery, len(allGalleryIDS))
		for i, id := range blogDetail.ExistingID {
			if id == nil {
				txFailed = true
				return fmt.Errorf("gallery id is nil/empty, issue on backend side")
			}
			blogGallery[i] = &entity.BlogGallery{
				BlogID:    blog.ID,
				GalleryID: *id,
			}
		}

		blogGalleryResponse, err := modelBlogGallery.InsertBlogGalleryMultiple(ctx, blogGallery)
		if err != nil {
			txFailed = true
			log.Printf("Failed to insert blog gallery %v: %v", blogGallery, err)
			return fmt.Errorf("failed to insert blog gallery %w", err)
		}

		blogImageJson, err := json.Marshal(blogGalleryResponse)
		if err != nil {
			return fmt.Errorf("failed to marshal blog gallery response %w", err)
		}
		blogDataResponse = dto.BlogResponse{
			ID:           blog.ID,
			AuthorID:     blog.AuthorID,
			Title:        blog.Title,
			Slug:         blog.Slug,
			Content:      blog.Content,
			Kategori:     blog.Kategori,
			ThumbnailURL: blog.ThumbnailURL,
			PublishedAt:  blog.PublishedAt,
			BlogImage:    blogImageJson,
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to insert database and Transaction failed %w", err)
	}

	return &blogDataResponse, nil
}

func (m *BlogService) GetAllBlogsForAdmin(ctx context.Context, kategori []string, offset, limit int) ([]dto.BlogThumbnail, int, error) {
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

func (m *BlogService) GetAllBlogs(ctx context.Context, page, limit int) ([]dto.BlogThumbnail, int, error) {

	offset := (page - 1) * limit

	blogs, totalData, err := m.BlogModel.GetAllBlogs(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	return blogs, totalData, nil
}

func (m *BlogService) GetBlogByID(id int) (*dto.BlogResponse, error) {

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
	blogDetail *dto.BlogPatch,
	newImages []*multipart.FileHeader,
	userRole string,
) (*dto.BlogResponse, error) {

	_, err := m.BlogModel.GetBlogById(idBlog)
	if err != nil {
		return nil, fmt.Errorf("blog not found you can't do update %w", err)
	}

	if err := authorization.CheckRolePermission(
		userRole,
		constants.RoleKeySuperAdmin,
		constants.RoleKeyKoorMedcrev,
	); err != nil {
		return nil, err
	}

	status, err := blogAuthorization.CanSetStatusBlog(userRole, blogDetail.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to set status %w", err)
	}
	var (
		newGalleryDataResponse []*dto.GalleryResponse
		allGalleryIDS          []int
	)

	galleryTouch := len(newImages) != 0 || len(blogDetail.ExistingID) != 0
	if galleryTouch {
		blogUpdate := &dto.BlogPayload{
			AuthorID:    authorID,
			Content:     blogDetail.Content,
			ExistingID:  blogDetail.ExistingID,
			Kategori:    blogDetail.Kategori,
			PublishedAt: blogDetail.PublishedAt,
			Slug:        blogDetail.Slug,
			Status:      status,
			Title:       blogDetail.Title,
		}
		newGalleryDataResponse, allGalleryIDS, err = m.ProcessBlogGalleries(
			ctx,
			newImages,
			blogUpdate,
		)
	}

	requestData := &entity.Blog{
		AuthorID:  authorID,
		Title:     blogDetail.Title,
		Slug:      blogDetail.Slug,
		Content:   blogDetail.Content,
		Kategori:  pq.StringArray(blogDetail.Kategori),
		Status:    status,
		UpdatedAt: time.Now(),
	}

	updatedBlog := utils.StructToMap(requestData)
	if len(newGalleryDataResponse) > 0 {
		updatedBlog["thumbnail_url"] = newGalleryDataResponse[0].FileURL
	}

	blogUpdate, err := m.BlogModel.UpdateBlogPartial(idBlog, updatedBlog)
	if err != nil {
		return nil, fmt.Errorf("failed to insert data %w", err)
	}

	// update blog_gallery -> jika terdapat terdapat perubahan foto
	var blogGallery []*dto.BlogGalleryResponse
	if galleryTouch {
		blogGalleryInsert, err := m.BlogGallery.UpdateBlogGallery(
			allGalleryIDS,
			blogUpdate.ID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to update gallery image %w", err)
		}
		for _, data := range blogGalleryInsert {
			blogGallery = append(blogGallery, &dto.BlogGalleryResponse{
				ID:        data.ID,
				BlogID:    data.BlogID,
				GalleryID: data.GalleryID,
			})
		}
	}
	blogImageJSON, err := json.Marshal(blogGallery)
	if err != nil {
		return nil, fmt.Errorf("failed to encode blog gallery: %w", err)
	}
	result := &dto.BlogResponse{
		ID:           blogUpdate.ID,
		AuthorID:     blogUpdate.AuthorID,
		Title:        blogUpdate.Title,
		Slug:         blogUpdate.Slug,
		Content:      blogUpdate.Content,
		Kategori:     blogUpdate.Kategori,
		ThumbnailURL: blogUpdate.ThumbnailURL,
		PublishedAt:  blogUpdate.PublishedAt,
		BlogImage:    blogImageJSON,
	}

	return result, nil
}

func (m *BlogService) GetAllBlogOrByKategori(
	ctx context.Context,
	kategori []string,
	limit, offset int,
) ([]dto.BlogThumbnail, int, error) {

	var (
		blog      []dto.BlogThumbnail
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
