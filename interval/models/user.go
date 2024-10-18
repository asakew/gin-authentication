package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (User) TableName() string {
	return "users_db" // Table name
}
