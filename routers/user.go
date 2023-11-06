package routers

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)


func SetUpUserAPI(router *gin.Engine) {
	userAPI := router.Group("api/v1")

	userAPI.POST("/user", controllers.CreateUser)
}