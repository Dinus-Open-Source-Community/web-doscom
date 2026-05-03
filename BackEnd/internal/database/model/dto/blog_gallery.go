package dto

type BlogGalleryInsert struct {
	BlogID    int `gorm:"column:id_blog" json:"blog_id"`
	GalleryID int `gorm:"column:id_gallery" json:"gallery_id"`
}

type BlogGalleryResponse struct {
	ID        int `json:"id"`
	BlogID    int `json:"blog_id"`
	GalleryID int `json:"gallery_id"`
}
