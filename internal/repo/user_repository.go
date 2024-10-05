package repo

import (
	"appGin/internal/models"
	"errors"
	"gorm.io/gorm"
)

var DB *gorm.DB

// CreateUser inserts a new user into the database
func CreateUser(user *models.User) error {
	return DB.Create(user).Error
}

// GetUserByUsername retrieves a user by their username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := DB.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	}
	return &user, result.Error
}
