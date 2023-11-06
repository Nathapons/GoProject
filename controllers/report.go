package controllers

import (
	"main/db"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateReport(c *gin.Context) {
	var reporter models.Reporter
	var err error = c.ShouldBind(&reporter)

	if err == nil {
		c.JSON(http.StatusCreated, gin.H{"data": reporter}) 
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func GetReporters(c *gin.Context) {
	var repoters []models.Reporter
	db.GetDB().Find(&repoters)
	c.JSON(http.StatusOK, gin.H{"result": repoters})
}