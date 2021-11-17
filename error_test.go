package errorstack

import (
	"log"
	"os"
	"testing"
)

func A() ErrorStack {
	err := New("call in A")
	return err
}

func B() ErrorStack {
	err := A()
	err.Trace("call in B")
	return err
}

func C() ErrorStack {
	err := B()
	err.Trace("call in C")
	return err
}

func PrintError(err error) {
	log.Println(err.Error())
}

func PrintStack(err ErrorStack) {
	log.Println(err.Stack())
}

func TestError(t *testing.T) {
	err := C()
	PrintError(err)
	PrintStack(err)
	// f, _ := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	err.Log(os.Stdout)
}
