package log

import (
	"fmt"
	"io"
)

type Logger struct {
	OutputText bool
	Output     io.Writer
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.OutputText {
		fmt.Fprintln(receiver.Output, a...)
	}
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}
