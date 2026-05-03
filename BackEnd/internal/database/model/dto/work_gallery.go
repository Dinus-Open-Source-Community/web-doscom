package dto

type WorkGalleryInsert struct {
	IDWork    int `gorm:"column:id_work" json:"id_work" binding:"required"`
	IDGallery int `gorm:"column:id_gallery" json:"id_gallery" binding:"required"`
}

type WorkGalleryResponse struct {
	ID        int `json:"id"`
	IDWork    int `json:"id_work"`
	IDGallery int `json:"id_gallery"`
}
