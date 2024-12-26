package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username" validate:"required,min=3,max=50"`
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	Verified bool   `gorm:"not null;default:false" json:"verified"`
	Urls     []Url  `json:"urls" gorm:"foreignKey:UserID"`
}

// Для запросов регистрации (чтобы не использовать основную модель)
type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// Для запросов логина
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
