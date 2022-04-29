package Controller

import (
	"fmt"
	"github.com/mhthrh/Voiceflex/Model/Calc"
	"github.com/mhthrh/Voiceflex/Model/Input/SignIn"
	"github.com/mhthrh/Voiceflex/Model/Input/SignOut"
	"github.com/mhthrh/Voiceflex/Model/Input/SignUp"
	"github.com/mhthrh/Voiceflex/Model/Result"
	"github.com/mhthrh/Voiceflex/Model/User"
	"github.com/mhthrh/Voiceflex/Utils/ConfigUtil"
	"github.com/mhthrh/Voiceflex/Utils/DbUtil/DbPool"
	"github.com/mhthrh/Voiceflex/Utils/JsonUtil"
	"github.com/mhthrh/Voiceflex/Utils/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Key struct{}

type Controller struct {
	l    *logrus.Entry
	v    *ValidationUtil.Validation
	db   *DbPool.DBs
	Usr  *User.User
	Conf *ConfigUtil.Config
}

var InvalidPath = fmt.Errorf("invalid Path, path must be /ViewControler/[id]")

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func New(l *logrus.Entry, v *ValidationUtil.Validation, db *DbPool.DBs, u *User.User, c *ConfigUtil.Config) *Controller {
	return &Controller{l, v, db, u, c}
}

func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {
	var i SignIn.SignIn
	i = r.Context().Value(Key{}).(SignIn.SignIn)
	user := User.New()
	user.UserName = i.UserName
	user.Password = i.Password

	d := c.db.Pull()
	u, err := user.SignIn(d.Db)
	if err != nil {
		Result.New(1011, http.StatusInternalServerError, err.Error()).SendResponse(w)
		return
	}
	c.db.Push(d)
	Result.New(1, http.StatusOK, JsonUtil.New(nil, nil).Struct2Json(u)).SendResponse(w)

}
func (c *Controller) SignOut(w http.ResponseWriter, r *http.Request) {
	var i SignOut.SignOut
	i = r.Context().Value(Key{}).(SignOut.SignOut)
	user := User.New()
	user.Ticket = i.Ticket
	d := c.db.Pull()
	if err := user.SignOut(d.Db); err != nil {
		Result.New(1012, http.StatusInternalServerError, err.Error()).SendResponse(w)
		return
	}
	c.db.Push(d)
	Result.New(1, http.StatusOK, "OK").SendResponse(w)

}
func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	var i SignUp.SignUp
	i = r.Context().Value(Key{}).(SignUp.SignUp)
	user := User.New()
	user.UserName = i.UserName
	user.Password = i.Password
	user.FirstName = i.FirstName
	user.LastName = i.LastName
	user.Email = i.Email
	user.CellNo = i.CellNo
	d := c.db.Pull()
	if err := user.SignUp(d.Db); err != nil {
		Result.New(1013, http.StatusInternalServerError, err.Error()).SendResponse(w)
		return
	}
	c.db.Push(d)
	Result.New(1, http.StatusOK, "OK").SendResponse(w)
}
func (c *Controller) Calculate(w http.ResponseWriter, r *http.Request) {
	var i Calc.Calc
	i = r.Context().Value(Key{}).(Calc.Calc)
	calc := Calc.New(i.Number1, i.Number2, i.Ticket)
	d := c.db.Pull()
	s, err := calc.Sum(d.Db)
	c.db.Push(d)
	if err != nil {
		Result.New(1019, http.StatusInternalServerError, err.Error()).SendResponse(w)
		return
	}
	Result.New(1, http.StatusOK, s).SendResponse(w)

}
