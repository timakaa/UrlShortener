package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string
	ShortUrl    string
	Visits      int
}
