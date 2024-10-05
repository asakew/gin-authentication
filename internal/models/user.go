package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return "users" // Change this to your table name
}
