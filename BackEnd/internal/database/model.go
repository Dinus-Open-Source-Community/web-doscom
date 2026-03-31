package database

import (
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type Models struct {
	Users        model.UserModel
	Works        model.WorkModel
	Blogs        model.BlogModel
	BlogGallery  model.BlogGalleryModel
	Gallery      model.GalleryModel
	Pengurus     model.PengurusModel
	FileUploads  model.FileUploadModel
	RefreshToken model.RefreshTokenModel
}

func NewModel(db *gorm.DB) Models {
	return Models{
		Users:       model.UserModel{DB: db},
		Works:       model.WorkModel{DB: db},
		Blogs:       model.BlogModel{DB: db},
		BlogGallery: model.BlogGalleryModel{DB: db},
		Gallery:     model.GalleryModel{DB: db},
		Pengurus:    model.PengurusModel{DB: db},
		FileUploads: model.FileUploadModel{DB: db},
		RefreshToken: model.RefreshTokenModel{
			DB: db,
		},
	}
}
