package db

import (
	"fmt"
	"main/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT string) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrateDB(database)

	db = database
}

func migrateDB(database *gorm.DB) {
	database.AutoMigrate(&models.Reporter{})
	database.AutoMigrate(&models.Article{})
}


func GetDB() *gorm.DB{
	return db
}

