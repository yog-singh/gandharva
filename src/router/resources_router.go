package router

import (
	"github.com/gin-gonic/gin"

	"github.com/yog-singh/gandharva/src/controllers"
)

func SetResourceRoutes(router *gin.Engine) {
	resourceRouter := router.Group("/resources")
	{
		resourceRouter.POST("/", controllers.AddResource)
		resourceRouter.GET("/", controllers.GetAllResources)
		resourceRouter.GET("/ping", controllers.PingResources)
	}
}
