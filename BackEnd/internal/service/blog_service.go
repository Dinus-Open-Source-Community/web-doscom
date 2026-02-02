package service

import (
	"web_doscom/internal/database/model"
)

type BlogService struct {
	Model *model.BlogModel
}

func NewBlogService(m *model.BlogModel) *BlogService {
	return &BlogService{Model: m}
}
