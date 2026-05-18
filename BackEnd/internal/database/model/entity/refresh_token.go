package entity

import (
	"gorm.io/gorm"
	"time"
)

type RefreshTokenModel struct {
	DB *gorm.DB
}

type RefreshToken struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	UserId    int        `json:"user_id"`
	Token     string     `json:"token_hash"`
	Expires   *time.Time `json:"expires"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (RefreshToken) TableName() string {
	return "refresh_token"
}

func (r *RefreshTokenModel) CreateRefreshToken(refreshToken *RefreshToken) error {
	if err := r.DB.Create(refreshToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenModel) GetRefreshToken(refreshToken string) (*RefreshToken, error) {
	var refreshtoken RefreshToken
	if err := r.DB.Where("token = ?", &refreshToken).First(&refreshtoken).Error; err != nil {
		return nil, err
	}
	return &refreshtoken, nil
}

func (r *RefreshTokenModel) DeleteRefreshTokenByUserId(userId int) error {
	var refreshToken RefreshToken

	if err := r.DB.Where("user_id = ?", &userId).Delete(&refreshToken).Error; err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenModel) DeleteRefreshToken(tokenHash string) error {
	if err := r.DB.Where("token = ?", &tokenHash).Delete(&RefreshToken{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *RefreshTokenModel) UpdatePartial(refreshToken string, data map[string]any) error {
	// return p.DB.WithContext(ctx).
	// 	Model(&Payment{}).
	// 	Where("id = ?", paymentID).
	// 	Updates(data).
	// 	Error

	return r.DB.Model(&RefreshToken{}).
		Where("token_hash = ?", refreshToken).
		Updates(data).
		Error
}
