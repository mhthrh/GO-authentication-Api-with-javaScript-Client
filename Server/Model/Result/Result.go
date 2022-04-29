package Result

import (
	"github.com/mhthrh/Voiceflex/Utils/JsonUtil"
	"net/http"
	"time"
)

type Response struct {
	Header struct {
		HStatus int
	}
	Body struct {
		BStatus int
		Time    time.Time
		Result  interface{}
	}
}

func New(BStatus, HStatus int, result interface{}) *Response {
	r := new(Response)
	r.Body.Time = time.Now()
	r.Header.HStatus = HStatus
	r.Body.BStatus = BStatus
	r.Body.Result = result
	return r
}
func (r *Response) SendResponse(w http.ResponseWriter) {
	w.Write([]byte(JsonUtil.New(nil, nil).Struct2Json(r.Body)))
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(r.Header.HStatus)
}
