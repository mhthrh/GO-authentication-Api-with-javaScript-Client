package JsonUtil

import (
	"encoding/json"
	"io"
)

type JsonFace interface {
	ToJSON()
	FromJSON()
}

type Json struct {
	w io.Writer
	r io.Reader
}

func New(wr io.Writer, re io.Reader) *Json {
	return &Json{
		w: wr,
		r: re,
	}
}

func (j *Json) ToJSON(i interface{}) error {
	e := json.NewEncoder(j.w)
	return e.Encode(i)
}

func (j *Json) FromJSON(i interface{}) error {
	d := json.NewDecoder(j.r)
	return d.Decode(i)
}
func (j *Json) Struct2Json(i interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func (j *Json) Json2Struct(b []byte, i interface{}) error {
	return json.Unmarshal(b, &i)
}
