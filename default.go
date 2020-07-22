package merr

import (
	"encoding/json"
	"fmt"
)

type M struct {
	Code interface{} `json:"code"`
	Text string      `json:"text"`
}

func (m *M) Error() string {
	return fmt.Sprintf("{\"code\":%#v,\"text\":\"%s\"}", m.Code, m.Text)
}

func New(code interface{}, text string) *M {
	return &M{Code: code, Text: text}
}

func UnMarshal(msg string) *M {
	m := new(M)
	err := json.Unmarshal([]byte(msg), &m)
	if err != nil {
		return &M{
			Code: "Unknown",
			Text: msg,
		}
	}
	return m
}
