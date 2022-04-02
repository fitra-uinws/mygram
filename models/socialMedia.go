package models

type SocialMedia struct {
	ID uint `gorm:"PrimaryKey" json:"id"`
	DefaultModel
	Name           string `gorm:"not null" validate:"required" json:"name"`
	SocialMediaUrl string `gorm:"not null" validate:"required" json:"social_media_url"`
	UserID         uint   `json:"user_id"`
}
