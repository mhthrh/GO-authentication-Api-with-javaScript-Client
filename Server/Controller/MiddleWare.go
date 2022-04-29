package Controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/mhthrh/Voiceflex/Model/Calc"
	"github.com/mhthrh/Voiceflex/Model/Input/SignIn"
	"github.com/mhthrh/Voiceflex/Model/Input/SignOut"
	"github.com/mhthrh/Voiceflex/Model/Input/SignUp"
	"github.com/mhthrh/Voiceflex/Model/Result"
	"github.com/mhthrh/Voiceflex/Utils/JsonUtil"
	"net/http"
	"strings"
)

func (b *Controller) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test")
		fNext := func(in interface{}) {
			r = r.WithContext(context.WithValue(r.Context(), Key{}, in))
			next.ServeHTTP(w, r)
		}
		b.l.WithFields(map[string]interface{}{
			"method": r.Method,
			"path":   r.URL,
			"status": nil,
		}).Info("request details")
		if r.Method != http.MethodPost {
			err := errors.New("invalid method")
			Result.New(1001, http.StatusForbidden, GenericError{Message: err.Error()}.Message).SendResponse(w)
			return
		}

		if r.Host != fmt.Sprintf("%s:%d", b.Conf.Server.IP, b.Conf.Server.Port) {
			err := errors.New("access denied")
			Result.New(1002, http.StatusForbidden, GenericError{Message: err.Error()}.Message).SendResponse(w)
			return
		}
		//teteet := r.Body
		//err := JsonUtil.New(nil, teteet).FromJSON(new(interface{}))
		//switch {
		//case err == io.EOF:
		//	Result.New(1003, http.StatusBadRequest, GenericError{Message: errors.New("empty request").Error()}.Message).SendResponse(w)
		//	return
		//case err != nil:
		//	Result.New(1004, http.StatusBadRequest, GenericError{Message: err.Error()}.Message).SendResponse(w)
		//	return
		//}
		switch strings.ToLower(r.URL.Path) {
		case "/signup":
			{
				var obj SignUp.SignUp
				err := JsonUtil.New(nil, r.Body).FromJSON(&obj)
				if err != nil {
					Result.New(1005, http.StatusBadRequest, GenericError{Message: err.Error()}.Message).SendResponse(w)
					return
				}
				errs := b.v.Validate(obj)
				if len(errs) != 0 {
					j := JsonUtil.New(nil, nil).Struct2Json(ValidationError{Messages: errs.Errors()}.Messages)
					Result.New(1006, http.StatusUnprocessableEntity, j).SendResponse(w)
					return
				}
				fNext(obj)
			}

		case "/signout":
			{
				var obj SignOut.SignOut
				err := JsonUtil.New(nil, r.Body).FromJSON(&obj)
				if err != nil {
					Result.New(1007, http.StatusBadRequest, GenericError{Message: err.Error()}.Message).SendResponse(w)
					return
				}
				errs := b.v.Validate(obj)
				if len(errs) != 0 {
					j := JsonUtil.New(nil, nil).Struct2Json(ValidationError{Messages: errs.Errors()}.Messages)
					Result.New(1008, http.StatusUnprocessableEntity, j).SendResponse(w)
					return
				}
				fNext(obj)
			}
		case "/signin":
			{
				obj := SignIn.SignIn{}
				err := JsonUtil.New(nil, r.Body).FromJSON(&obj)
				if err != nil {
					Result.New(1009, http.StatusBadRequest, GenericError{Message: err.Error()}.Message).SendResponse(w)
					return
				}
				errs := b.v.Validate(obj)
				if len(errs) != 0 {
					j := JsonUtil.New(nil, nil).Struct2Json(ValidationError{Messages: errs.Errors()}.Messages)
					Result.New(1010, http.StatusUnprocessableEntity, j).SendResponse(w)
					return
				}
				fNext(obj)
			}
		case "/calc":
			{
				var obj Calc.Calc
				err := JsonUtil.New(nil, r.Body).FromJSON(&obj)
				if err != nil {
					Result.New(1009, http.StatusBadRequest, GenericError{Message: err.Error()}.Message).SendResponse(w)
					return
				}
				errs := b.v.Validate(obj)
				if len(errs) != 0 {
					j := JsonUtil.New(nil, nil).Struct2Json(ValidationError{Messages: errs.Errors()}.Messages)
					Result.New(1010, http.StatusUnprocessableEntity, j).SendResponse(w)
					return
				}
				fNext(obj)
			}
		default:
			{
				err := errors.New("NotImplemented")
				Result.New(1020, http.StatusNotImplemented, GenericError{Message: err.Error()}.Message).SendResponse(w)
			}
		}

	})
}
