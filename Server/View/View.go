package View

import (
	"github.com/gorilla/mux"
	"github.com/mhthrh/Voiceflex/Controller"
	"github.com/mhthrh/Voiceflex/Model/User"
	"github.com/mhthrh/Voiceflex/Utils/ConfigUtil"
	"github.com/mhthrh/Voiceflex/Utils/DbUtil/DbPool"
	"github.com/mhthrh/Voiceflex/Utils/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunApiOnRouter(sm *mux.Router, log *logrus.Entry, db *DbPool.DBs, config *ConfigUtil.Config) {
	ph := Controller.New(log, ValidationUtil.NewValidation(), db, User.New(), config)
	sm.Use(ph.Middleware)
	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/signOut", ph.SignOut)
	postR.HandleFunc("/signUp", ph.SignUp)
	postR.HandleFunc("/signIn", ph.SignIn)
	postR.HandleFunc("/calc", ph.Calculate)
}
