package models

import (
	"time"
)

type VerificationCode struct {
  ID        uint      `gorm:"primarykey"`
	Email     string    `gorm:"index"`
	Code      string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (v *VerificationCode) IsExpired() bool {
	return time.Now().After(v.ExpiresAt)
}
