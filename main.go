package main

import (
	"log"
	"main/db"
	"main/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	router := gin.New()

	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("PORT")

	db.SetDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)
	
	routers.SetUpUserAPI(router)
	routers.SetReportRouter(router)

	router.Run(":" + PORT)
}
