package middleware

import (
	"fmt"
	"net/http"

	"github.com/yog-singh/gandharva/src/entity"
)

func CheckHeartbeat(resource entity.Resource) *http.Response {
	var resp *http.Response
	var req *http.Request
	var err error
	switch resource.Type {
	case "GET":
		resp, err = http.Get(resource.Url)
	case "POST":
		resp, err = http.Post(resource.Url, "application/json", nil)
	case "PUT":
		req, _ = http.NewRequest(http.MethodPut, resource.Url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err = client.Do(req)
	case "DELETE":
		req, _ = http.NewRequest(http.MethodDelete, resource.Url, nil)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err = client.Do(req)
	}

	if err != nil {
		fmt.Printf("Error checking heartbeat for: %s \n", resource.Name)
	}
	fmt.Printf("Success checking heartbeat for: %s \n", resource.Name)
	return resp
}
