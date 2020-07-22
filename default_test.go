package merr

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("TestCode", "I am message")
	fmt.Println(err)

}

func testReturnErr() error {
	return New(1000, "some error")
}

func TestM_Error(t *testing.T) {
	err := testReturnErr()
	fmt.Println(err.(*M))
}

func TestUnMarshal(t *testing.T) {
	var errJsonStr = "{\"code\":\"test_code\",\"text\":\"test message\"}"
	m := UnMarshal(errJsonStr)
	fmt.Println(m)
}

func TestUnMarshal2(t *testing.T) {
	var errJsonStr = "test message 2"
	m := UnMarshal(errJsonStr)
	fmt.Println(m)
}
