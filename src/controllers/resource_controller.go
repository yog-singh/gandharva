package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yog-singh/gandharva/src/config"
	"github.com/yog-singh/gandharva/src/entity"
	"github.com/yog-singh/gandharva/src/model"
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

func GetStatusPage(ctx *gin.Context) {
	response, err := services.GetAllResources()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	config, _ := config.LoadConfig(".")
	htmlDataMap := model.HTMLTemplateData{BaseURL: config.BaseURL, Resources: response}
	ctx.HTML(http.StatusOK, "index.tmpl", htmlDataMap)
}

func PingResources(ctx *gin.Context) {
	err := services.CheckResourceHeartbeat()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
