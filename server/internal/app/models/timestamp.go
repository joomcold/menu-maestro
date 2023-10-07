package models

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"     json:"deletedAt"`
}
