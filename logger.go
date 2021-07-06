package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

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
	if receiver.Prfx != nil {
		fmt.Fprint(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) Logf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	if receiver.Prfx != nil {
		fmt.Fprint(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	format = format + "\n"
	fmt.Fprintf(receiver.Output, format, a...)
}

func (receiver Logger) LogWithFuncName(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fnc := runtime.FuncForPC(pc)
		functionName := fnc.Name()
		fmt.Fprintf(receiver.Output, "%s: ", functionName)
	}
	if receiver.Prfx != nil {
		fmt.Fprint(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) LogfWithFuncName(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fnc := runtime.FuncForPC(pc)
		functionName := fnc.Name()
		fmt.Fprintf(receiver.Output, "%s: ", functionName)
	}
	if receiver.Prfx != nil {
		fmt.Fprint(receiver.Output, strings.Join(receiver.Prfx[:], ": "))
		fmt.Fprintf(receiver.Output, ": ")
	}
	fmt.Fprintf(receiver.Output, format, a...)
}

func (receiver Logger) Begin(a ...interface{}) {
	startTime = time.Now()
	receiver.Log("BEGIN")
}

func (receiver Logger) End(a ...interface{}) {
	endTime = time.Now()
	receiver.Log("END", endTime.Sub(startTime))
}

func (receiver Logger) Prefix(newprefix ...string) Logger {
	logger := Logger{
		Prfx:   newprefix,
		Output: Default.Output,
	}
	return logger
}

func (receiver Logger) output(a ...interface{}) {
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

func Log(a ...interface{}) {
	Default.Log(a...)
}

func Alert(a ...interface{}) {
	Default.Log(a...)
}

func Error(a ...interface{}) {
	Default.Log(a...)
}

func Highlight(a ...interface{}) {
	Default.Log(a...)
}

func Inform(a ...interface{}) {
	Default.Log(a...)
}

func Trace(a ...interface{}) {
	Default.Log(a...)
}

func Warn(a ...interface{}) {
	Default.Log(a...)
}

func Logf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Alertf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Highlightf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Informf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Tracef(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Begin(a ...interface{}) {
	Default.Begin(a...)
}

func End(a ...interface{}) {
	Default.End(a...)
}

func OutputFn(a ...interface{}) {
	Default.output(a...)
}
