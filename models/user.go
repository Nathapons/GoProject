package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json: "username" gorm:"size(150)"`
	Password string `json: "password" gorm:"size(150)"` 
	FirstName string `json:"first_name" gorm:"size(150)"`
	LastName string `json:"last_name" gorm:"size(150)"`
	Email string `json:"email" gorm:"size(255)"`
	IsActive bool `default:"false"`
}

func(user User) ShowDetail() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}