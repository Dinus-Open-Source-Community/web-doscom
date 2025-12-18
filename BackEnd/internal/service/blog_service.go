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
