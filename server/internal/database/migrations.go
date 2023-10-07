package database

import (
	"github.com/joomcold/go-next-docker/internal/app/models"
	"github.com/joomcold/go-next-docker/internal/initializers"
)

func Migrations() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate models")
	}
}
