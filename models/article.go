package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	HeadLine time.Time `json:"head_line" gorm:"size(200)"`
	Content string `json:"content" gorm:"type:text"`
	ReporterID uint      // Foreign key to link with Reporter's ID
	Reporter Reporter `gorm:"foreignkey:ReporterId;references:Id`
}
