package controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	var err error = c.ShouldBind(&user)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"result": user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"result": err})
	}
}
