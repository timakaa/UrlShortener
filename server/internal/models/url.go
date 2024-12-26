package models

import "gorm.io/gorm"


type Url struct {
	gorm.Model
	OriginalUrl string `json:"original_url" gorm:"not null"`
	ShortUrl    string `json:"short_url" gorm:"not null"`
	Visits      int    `json:"visits" gorm:"default:0"`
	User        User   `json:"user" gorm:"foreignKey:UserID"`
	UserID      uint   `json:"user_id" gorm:"not null"`
}
