package merr

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New("TestCode", "I am message")
	fmt.Println(err)
}

func TestUnmarshal(t *testing.T) {
	err := New("TestCode", "I am message")
	c := Unmarshal(err)
	fmt.Println(c.All())
	fmt.Println(c.MCode())
	fmt.Println(c.MText())

	fmt.Println("=============================")

	err2 := errors.New("errors without marshal")
	c2 := Unmarshal(err2)
	fmt.Println(c2.All())
	fmt.Println(c2.MCode())
	fmt.Println(c2.MText())
}
