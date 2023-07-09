package model

import (
	"github.com/yog-singh/gandharva/src/entity"
)

type StatusPageHTMLPayload struct {
	BaseURL                    string
	ResourcesUpCount           int32
	ResourcesDownCount         int32
	ResourcesInaccessibleCount int32
	ResourcesUp                []entity.Resource
	ResourcesDown              []entity.Resource
	ResourcesInaccessible      []entity.Resource
}

func GenerateStatusPageHTMLPayload(baseUrl string, resources []entity.Resource) StatusPageHTMLPayload {
	var resourcesUpCount, resourcesDownCount, resourcesInaccessibleCount int32
	var resourcesUp, resourcesDown, resourcesInaccessible []entity.Resource
	for _, resource := range resources {
		switch resource.Status {
		case entity.RESOURCE_UP, entity.PENDING:
			resourcesUpCount = resourcesUpCount + 1
			resourcesUp = append(resourcesUp, resource)
		case entity.RESOURCE_DOWN:
			resourcesDownCount = resourcesDownCount + 1
			resourcesDown = append(resourcesDown, resource)
		case entity.RESOURCE_INACCESSIBLE:
			resourcesInaccessibleCount = resourcesInaccessibleCount + 1
			resourcesInaccessible = append(resourcesInaccessible, resource)
		}
	}
	statusPagePayload := StatusPageHTMLPayload{
		BaseURL:                    baseUrl,
		ResourcesUpCount:           resourcesUpCount,
		ResourcesDownCount:         resourcesDownCount,
		ResourcesInaccessibleCount: resourcesInaccessibleCount,
		ResourcesUp:                resourcesUp,
		ResourcesDown:              resourcesDown,
		ResourcesInaccessible:      resourcesInaccessible,
	}
	return statusPagePayload
}
