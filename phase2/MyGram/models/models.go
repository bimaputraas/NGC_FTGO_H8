package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"Index" json:"username" binding:"required"`
	Email    string `gorm:"Index" json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,numeric,min=8"`
}
// type User struct {
// 	gorm.Model
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// 	Age      int    `json:"age"`
// }

type Photo struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" binding:"required"`
	UserID   uint   `json:"user_id"`
}

type SocialMedia struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
	UserID         uint   `json:"user_id"`
}

type Comment struct {
	gorm.Model
	Message string `json:"name" binding:"required"`
	UserID  uint   `json:"user_id"`
	PhotoID uint   `json:"photo_id"`
}
