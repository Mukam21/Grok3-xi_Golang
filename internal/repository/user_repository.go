package repository

import (
	"github.com/project-Grok3_Golang/internal/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

func CreateUser(user models.User) error {
	result := db.Create(&user)
	return result.Error
}
