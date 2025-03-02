package services

import (
	"github.com/project-Grok3_Golang/internal/database"
	"github.com/project-Grok3_Golang/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
