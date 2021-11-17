package errorstack

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"strings"
)

type ErrorStack interface {
	error
	Log(w io.Writer)
	Stack() string
	Trace(msg string)
}

type Error struct {
	Msg    string
	Stacks []string
}

func (err Error) Error() string {
	return err.Msg
}

func (err Error) Log(w io.Writer) {
	if w != nil {
		log.SetOutput(w)
	}
	log.Println(err.Stack())
}

func (err Error) Stack() string {
	return fmt.Sprintf("[ERROR]: %s\nTraceback:\n%s", err.Msg, strings.Join(err.Stacks, "\n"))
}

func (err *Error) Trace(msg string) {
	err.Msg = msg
	if len(err.Stacks) == 0 {
		err.Stacks = append(err.Stacks, fmt.Sprintf("%s %s", err.runtime(), msg))
	} else {
		stack := ""
		for i := 0; i < len(err.Stacks); i++ {
			stack += " "
		}
		stack += "|- " + err.runtime() + " " + msg
		err.Stacks = append(err.Stacks, stack)
	}
}

func (err Error) runtime() string {
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf("%s:%d", file, line)
}

func New(msg string) ErrorStack {
	err := &Error{}
	err.Trace(msg)
	return err
}
