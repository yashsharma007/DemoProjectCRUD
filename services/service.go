package services

import (
	"demoProject/models"
	"demoProject/repositories"
)

func GetUsers() ([]models.User, error) {
	return repositories.GetUsers()
}

func CreateUser(user models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByID(id int) (models.User, error) {
	return repositories.GetUserByID(id)
}

func DeleteUser(id int) error {
	return repositories.DeleteUser(id)
}
