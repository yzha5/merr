package merr

import (
	"fmt"
)

type M struct {
	Code interface{}
	Text string
}

func (m *M) Error() string {
	return fmt.Sprintf("{\"code\":%#v,\"text\":\"%s\"}", m.Code, m.Text)
}

func New(code interface{}, text string) *M {
	return &M{Code: code, Text: text}
}
