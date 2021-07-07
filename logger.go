package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

type Logger struct {
	Output io.Writer
	Prfx   []string
}

var Default Logger = Logger{
	Output: os.Stdout,
	Prfx:   nil,
}

var startTime time.Time
var endTime time.Time

func (receiver Logger) Log(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(a...)
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(format, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	startTime = time.Now()
	receiver.output("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	endTime = time.Now()
	receiver.output("END", endTime.Sub(startTime))
}

func (receiver Logger) Prefix(newprefix ...string) Logger {
	logger := Logger{
		Prfx:   newprefix,
		Output: Default.Output,
	}
	return logger
}

func (receiver Logger) output(a ...interface{}) {
	if receiver.Prfx != nil {
		var prefixStr string = strings.Join(receiver.Prfx[:], ": ") + ":"
		a = append([]interface{}{prefixStr}, a...)
	}
	var functionName string = ""
	{
		pc, _, _, ok := runtime.Caller(2)
		if ok {
			fnc := runtime.FuncForPC(pc)
			functionName = fnc.Name()
		}
	}
	a = append([]interface{}{functionName}, a...)
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) outputf(format string, a ...interface{}) {
	if receiver.Prfx != nil {
		var prefixStr string = strings.Join(receiver.Prfx[:], ": ") + ":"
		a = append([]interface{}{prefixStr}, a...)
	}
	var functionName string = ""
	{
		pc, _, _, ok := runtime.Caller(2)
		if ok {
			fnc := runtime.FuncForPC(pc)
			functionName = fnc.Name()
		}
	}
	a = append([]interface{}{functionName}, a...)
	a = append(a, "\n")
	fmt.Fprintf(receiver.Output, format, a...)
}

func Log(a ...interface{}) {
	Default.Log(a...)
}

func Alert(a ...interface{}) {
	// TODO use alert method
	Default.Log(Green(a...))
}

func Error(a ...interface{}) {
	Default.Log(Red(a...))
}

func Highlight(a ...interface{}) {
	Default.Log(Teal(a...))
}

func Inform(a ...interface{}) {
	Default.Log(Magenta(a...))
}

func Trace(a ...interface{}) {
	Default.Log(a...)
}

func Warn(a ...interface{}) {
	Default.Log(Yellow(a...))
}

func Logf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Alertf(format string, a ...interface{}) {
	Default.Logf(format, Green(a...))
}

func Errorf(format string, a ...interface{}) {
	Default.Logf(format, Red(a...))
}

func Highlightf(format string, a ...interface{}) {
	Default.Logf(format, Teal(a...))
}

func Informf(format string, a ...interface{}) {
	Default.Logf(format, Magenta(a...))
}

func Tracef(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Default.Logf(format, Yellow(a...))
}

func Begin(a ...interface{}) {
	Default.Begin(a...)
}

func End(a ...interface{}) {
	Default.End(a...)
}
