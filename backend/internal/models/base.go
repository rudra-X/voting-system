package models

import (
	"time"
	"gorm.io/gorm"
)

type Base struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `gorm:"is_deleted"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = b.CreatedAt
	return
}

func (b *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now()
	return
}
