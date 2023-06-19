package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		message := "pong"
		ctx.JSON(http.StatusOK, message)
	})

	SetResourceRoutes(router)
}
