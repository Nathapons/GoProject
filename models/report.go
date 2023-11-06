package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Reporter struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"size(128)"`
	LastName  string `json:"last_name" gorm:"size(128)"`
	Article []Article // One-to-many relationship with Article
}

func (r Reporter) GetFullName() string {
	return fmt.Sprintf("%s %s", r.FirstName, r.LastName)
}
