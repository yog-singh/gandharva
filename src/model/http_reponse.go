package model

import "net/http"

type HTTPReponse struct {
	RequestCompletionTimeInMs int64
	Response                  *http.Response
}
