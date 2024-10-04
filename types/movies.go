package types

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Title       string         `json:"title"       validate:"required"`
	Description string         `json:"description"`
	Rating      float64        `json:"rating"`
	Image       string         `json:"image"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
