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
