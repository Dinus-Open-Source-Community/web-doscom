package database

import (
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type Models struct {
<<<<<<< HEAD
	Users    model.UserModel
	Works    model.WorkModel
	Blogs    model.BlogModel
	Gallery  model.GalleryModel
	Pengurus model.PengurusModel
=======
	Users model.UserModel
	Works model.WorkModel
<<<<<<< Updated upstream
	// Activities model.ActivityModel
=======
>>>>>>> Stashed changes
	Blogs model.BlogModel
>>>>>>> 858e57d (add : update kategori & list by kategori)
}

func NewModel(db *gorm.DB) Models {
	return Models{
<<<<<<< HEAD
		Users:    model.UserModel{DB: db},
		Works:    model.WorkModel{DB: db},
		Blogs:    model.BlogModel{DB: db},
		Gallery:  model.GalleryModel{DB: db},
		Pengurus: model.PengurusModel{DB: db},
=======
		Users: model.UserModel{DB: db},
		Works: model.WorkModel{DB: db},
<<<<<<< Updated upstream
		// Activities: model.ActivityModel{DB: db},
=======
>>>>>>> Stashed changes
		Blogs: model.BlogModel{DB: db},
>>>>>>> 858e57d (add : update kategori & list by kategori)
	}
}
