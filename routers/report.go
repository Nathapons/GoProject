package routers

import (
	"main/controllers"

	"github.com/gin-gonic/gin"
)

func SetReportRouter(router *gin.Engine) {
	reportAPI := router.Group("api/v1")

	reportAPI.POST("/reporter", controllers.CreateReport)
	reportAPI.GET("/reporter", controllers.GetReporters)
}
