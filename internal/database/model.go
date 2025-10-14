package database

import (
	"gorm.io/gorm"
)

type Models struct {
	Users UserModel
}

func NewModel(db *gorm.DB) Models {
	return Models{
		Users: UserModel{DB: db},
	}
}
