package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/yog-singh/gandharva/src/entity"
	"github.com/yog-singh/gandharva/src/model"
)

func CheckHeartbeat(resource entity.Resource) model.HTTPReponse {
	var resp *http.Response
	var req *http.Request
	var err error
	var startTime, endTime time.Time
	var requestCompletionTime int64

	requestStatistics := model.RequestStatistics{}
	model.InitRequestStatistics(&requestStatistics)

	switch resource.RequestMethod {
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
	clientTraceCtx := httptrace.WithClientTrace(req.Context(), requestStatistics.ClientTrace)
	req = req.WithContext(clientTraceCtx)
	startTime = time.Now()
	resp, err = http.DefaultTransport.RoundTrip(req)
	endTime = time.Now()
	if err != nil {
		fmt.Printf("Error checking heartbeat for: %s \n", resource.Name)
	}

	requestCompletionTime = (endTime.UnixNano() / int64(time.Millisecond)) - (startTime.UnixNano() / int64(time.Millisecond))
	fmt.Printf("Latency(ms): %d \n", requestCompletionTime)
	fmt.Printf("Response stats: %+v\n", requestStatistics)
	response := model.HTTPReponse{RequestCompletionTimeInMs: requestCompletionTime, Response: resp}
	return response
}
