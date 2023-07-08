package model

import (
	"net/http"

	"github.com/yog-singh/gandharva/src/entity"
)

type HTTPReponse struct {
	RequestCompletionTimeInMs int64
	Response                  *http.Response
	ResponseBody              []byte
	HTTPTiming                entity.HTTPTiming
}
