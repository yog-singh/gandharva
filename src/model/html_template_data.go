package model

import "github.com/yog-singh/gandharva/src/entity"

type HTMLTemplateData struct {
	BaseURL   string
	Resources []entity.Resource
}
