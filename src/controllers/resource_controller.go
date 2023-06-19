package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yog-singh/gandharva/src/entity"
	"github.com/yog-singh/gandharva/src/services"
)

func AddResource(ctx *gin.Context) {
	var resource entity.Resource
	if err := ctx.BindJSON(&resource); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response, err := services.AddResource(resource)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusAccepted, response)
}

func GetAllResources(ctx *gin.Context) {
	response, err := services.GetAllResources()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func PingResources(ctx *gin.Context) {
	err := services.PingResources()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
