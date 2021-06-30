package log

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Logger struct {
	Output io.Writer
	Prfx   []string
}

var Default Logger = Logger{
	Output: os.Stdout,
	Prfx:   nil,
}

func (receiver Logger) Log(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	if receiver.Prfx != nil {
		fmt.Fprintf(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	if receiver.Prfx != nil {
		fmt.Fprintf(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	fmt.Fprintf(receiver.Output, format, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	receiver.Log("END")
}

func (receiver Logger) Prefix(newprefix ...string) Logger {
	logger := Logger{
		Prfx:   newprefix,
		Output: Default.Output,
	}
	return logger
}

func Log(a ...interface{}) {
	Default.Log(a...)
}

func Logf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}
