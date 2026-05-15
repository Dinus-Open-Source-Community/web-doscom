package database

import (
	"web_doscom/internal/database/model/entity"

	"gorm.io/gorm"
)

type Models struct {
	Users          entity.UserModel
	Works          entity.WorkModel
	WorkGallery    entity.WorkGalleryModel
	Blogs          entity.BlogModel
	BlogGallery    entity.BlogGalleryModel
	Gallery        entity.GalleryModel
	Pengurus       entity.PengurusModel
	PengurusSosmed entity.PengurusSosmedModel
	FileUploads    entity.FileUploadModel
	RefreshToken   entity.RefreshTokenModel
}

func NewModel(db *gorm.DB) Models {
	return Models{
		Users:       entity.UserModel{DB: db},
		Works:       entity.WorkModel{DB: db},
		WorkGallery: entity.WorkGalleryModel{DB: db},
		Blogs:       entity.BlogModel{DB: db},
		BlogGallery: entity.BlogGalleryModel{DB: db},
		Gallery:     entity.GalleryModel{DB: db},
		Pengurus:    entity.PengurusModel{DB: db},
		PengurusSosmed: entity.PengurusSosmedModel{
			DB: db,
		},
		FileUploads: entity.FileUploadModel{DB: db},
		RefreshToken: entity.RefreshTokenModel{
			DB: db,
		},
	}
}
