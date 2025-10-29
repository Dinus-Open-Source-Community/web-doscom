package database

import (
	"web_doscom/internal/database/model"

	"gorm.io/gorm"
)

type Models struct {
	Users model.UserModel
	Works model.WorkModel
	Blogs model.BlogModel
}

func NewModel(db *gorm.DB) Models {
	return Models{
		Users: model.UserModel{DB: db},
		Works: model.WorkModel{DB: db},
		Blogs: model.BlogModel{DB: db},
	}
}
