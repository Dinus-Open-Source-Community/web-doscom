package service

import (
	"fmt"
	"mime/multipart"
	"web_doscom/internal/database/model"
)

type BlogService struct {
	BlogModel    *model.BlogModel
	BlogGallery  *model.BlogGalleryModel
	GalleryModel *GalleryService
}

func NewBlogService(n *model.BlogModel, m *model.BlogGalleryModel, g *GalleryService) *BlogService {
	return &BlogService{
		BlogModel:    n,
		BlogGallery:  m,
		GalleryModel: g,
	}
}

func (m *BlogService) CreateBlogImage(blogID int, existingID []int, newImages []*multipart.FileHeader) ([]*model.BlogGallery, error) {
	var result []*model.BlogGallery
	// kalau pake foto baru maka upload ke storage
	if len(newImages) > 0 {
		createGallery := &model.CreateGallery{
			GalleryType: "blog",
			Description: "foto untuk kepentingan blog",
			EventDate:   "",
		}

		uploadedFile, err := m.GalleryModel.UploadAndInsertGallery(newImages, createGallery)
		if err != nil {
			return nil, fmt.Errorf("Failed to upload and insert image")
		}

		var galleryIDS []int
		for _, gallery := range uploadedFile {
			galleryIDS = append(galleryIDS, gallery.ID)
		}

		// insert to database
		for _, id := range existingID {
			galleryBlog := &model.BlogGallery{
				BlogID:    blogID,
				GalleryID: int(id),
			}

			upload, err := m.BlogGallery.InsertBlogGallery(galleryBlog)
			if err != nil {
				return nil, fmt.Errorf("Failed to insert data")
			}
			result = append(result, upload)
		}
	}

	if len(existingID) > 0 {
		for _, id := range existingID {

			upload, err := m.BlogGallery.GetBlogGalleryByID(int(id))
			if err != nil {
				return nil, fmt.Errorf("Failed to insert data")
			}

			galleryBlog := &model.BlogGallery{
				BlogID:    blogID,
				GalleryID: upload.ID,
			}
			result = append(result, galleryBlog)
		}
	}

	return result, nil
}

func (s *BlogService) GetAllBlogs() ([]model.Blog, error) {
	return s.BlogModel.GetAllBlogs()
}

func (s *BlogService) GetBlogByID(id int) (*model.Blog, error) {
	return s.BlogModel.GetBlogById(id)
}

func (s *BlogService) UpdateBlog(id int, patch model.BlogPatch) (*model.Blog, error) {
	return s.BlogModel.UpdateBlog(id, patch)
}

func (s *BlogService) GetBlogsByKategori(kategori string) ([]model.Blog, error) {
	var blogs []model.Blog
	if err := s.BlogModel.DB.Where("kategori = ?", kategori).Order("created_at DESC").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func (s *BlogService) DeleteBlog(id int) error {
	return s.BlogModel.DeleteBlog(id)
}
