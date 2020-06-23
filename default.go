package merr

import (
	"encoding/json"
	"errors"
	"fmt"
)

type C interface {
	MCode() interface{}
	MText() string
	All() string
}

type M struct {
	Code interface{}
	Text string
}

func (m *M) MCode() interface{} {
	return m.Code
}

func (m *M) MText() string {
	return m.Text
}

func (m *M) All() string {
	mJson, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("{\"Code\":\"ParseException\",\"Text\":\"%v\"}", err)
	}
	return string(mJson)
}

func New(code interface{}, text string) error {
	mJson, err := json.Marshal(&M{
		Code: code,
		Text: text,
	})
	if err != nil {
		return fmt.Errorf("{\"Code\":\"ParseException\",\"Text\":\"%v\"}", err)
	}
	return errors.New(string(mJson))
}

func Unmarshal(err error) C {
	newM := new(M)
	e := json.Unmarshal([]byte(err.Error()), newM)
	if e != nil {
		return &M{
			Code: "ParseException",
			Text: err.Error(),
		}
	}
	return newM
}
