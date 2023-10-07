package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"unique;not null"      json:"email"`
	Password string    `gorm:"not null"             json:"-"`
	Address  string    `json:"address"`
	Timestamp
}
