package log

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	Output io.Writer
}

var Default Logger = Logger{
	Output: os.Stdout,
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	fmt.Fprintf(receiver.Output, format, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}

func Log(a ...interface{}) {
	Default.Log(a...)
}

func Logf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}
