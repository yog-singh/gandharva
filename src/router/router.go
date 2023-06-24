package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yog-singh/gandharva/src/controllers"
)

func InitRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("src/view/*")
	router.GET("/ping", func(ctx *gin.Context) {
		message := "pong"
		ctx.JSON(http.StatusOK, message)
	})
	router.GET("/", controllers.GetStatusPage)

	SetResourceRoutes(router)
}
