package middleware

import (
	"fmt"
	"io"
	"log"
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
	if err != nil {
		fmt.Printf("Error checking heartbeat for: %s \n", resource.Name)
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	endTime = time.Now()
	requestStatistics.ReadResponseDoneTime = endTime

	requestCompletionTime = (endTime.UnixNano() / int64(time.Millisecond)) - (startTime.UnixNano() / int64(time.Millisecond))
	httpTiming := getHTTPTiming(requestStatistics)
	response := model.HTTPReponse{RequestCompletionTimeInMs: requestCompletionTime, Response: resp, ResponseBody: responseBody, HTTPTiming: httpTiming}
	return response
}

func getHTTPTiming(rs model.RequestStatistics) entity.HTTPTiming {
	httpTiming := entity.HTTPTiming{}
	httpTiming.DNSLookupTime = (rs.DNSDoneTime.UnixNano() / int64(time.Millisecond)) - (rs.DNSStartTime.UnixNano() / int64(time.Millisecond))
	httpTiming.ConnectionTime = (rs.ConnectDoneTime.UnixNano() / int64(time.Millisecond)) - (rs.ConnectStartTime.UnixNano() / int64(time.Millisecond))
	httpTiming.TLSHandshakeTime = (rs.TLSHandshakeDoneTime.UnixNano() / int64(time.Millisecond)) - (rs.TLSHandshakeStartTime.UnixNano() / int64(time.Millisecond))
	httpTiming.FirstByteWaitTime = (rs.GotFirstResponseByteTime.UnixNano() / int64(time.Millisecond)) - (rs.RequestStartTime.UnixNano() / int64(time.Millisecond))
	httpTiming.TotalTime = (rs.ReadResponseDoneTime.UnixNano() / int64(time.Millisecond)) - (rs.RequestStartTime.UnixNano() / int64(time.Millisecond))
	return httpTiming
}
