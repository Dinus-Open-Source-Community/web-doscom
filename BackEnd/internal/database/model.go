package database

import (
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type Models struct {
	Users       model.UserModel
	Works       model.WorkModel
	Blogs       model.BlogModel
	Gallery     model.GalleryModel
	Pengurus    model.PengurusModel
	BlogGallery model.BlogGalleryModel
	Blog        model.BlogModel
	// Activities model.ActivityModel // Uncomment if you have ActivityModel
}

func NewModel(db *gorm.DB) Models {
	return Models{
		Users:       model.UserModel{DB: db},
		Works:       model.WorkModel{DB: db},
		Blogs:       model.BlogModel{DB: db},
		Gallery:     model.GalleryModel{DB: db},
		Pengurus:    model.PengurusModel{DB: db},
		BlogGallery: model.BlogGalleryModel{DB: db},
		Blog:        model.BlogModel{DB: db},
		// Activities: model.ActivityModel{DB: db}, // Uncomment if you have ActivityModel
	}
}
