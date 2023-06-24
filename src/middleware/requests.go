package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptrace"

	"github.com/yog-singh/gandharva/src/entity"
)

func CheckHeartbeat(resource entity.Resource) *http.Response {
	var resp *http.Response
	var req *http.Request
	var err error
	clientTrace := &httptrace.ClientTrace{}

	switch resource.Type {
	case "GET":
		req, _ = http.NewRequest(http.MethodGet, resource.Url, nil)
	case "POST":
		req, _ = http.NewRequest(http.MethodPost, resource.Url, nil)
	case "PUT":
		req, _ = http.NewRequest(http.MethodPut, resource.Url, nil)
	case "DELETE":
		req, _ = http.NewRequest(http.MethodDelete, resource.Url, nil)
	}

	req.Header.Set("Content-Type", "application/json")
	clientTraceCtx := httptrace.WithClientTrace(req.Context(), clientTrace)
	req = req.WithContext(clientTraceCtx)
	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error checking heartbeat for: %s \n", resource.Name)
	}
	fmt.Printf("Success checking heartbeat for: %s \n", resource.Name)
	return resp
}
