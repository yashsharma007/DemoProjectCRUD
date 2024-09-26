package repositories

import (
	"demoProject/models"
	"demoProject/utils"
)

// Get all users
func GetUsers() ([]models.User, error) {
	var users []models.User
	rows, err := utils.DB.Query("SELECT id, name, address FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Create a new user
func CreateUser(user models.User) error {
	_, err := utils.DB.Exec("INSERT INTO users (name, address) VALUES ($1, $2)", user.Name, user.Address)
	return err
}

// Get a single user by ID
func GetUserByID(id int) (models.User, error) {
	var user models.User
	err := utils.DB.QueryRow("SELECT id, name, address FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Address)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete a user by ID
func DeleteUser(id int) error {
	_, err := utils.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
