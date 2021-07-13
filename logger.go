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
	Black   = color("\033[1;30m%s\033[0m")
	Red     = color("\033[1;31m%s\033[0m")
	Green   = color("\033[1;32m%s\033[0m")
	Yellow  = color("\033[1;33m%s\033[0m")
	Purple  = color("\033[1;34m%s\033[0m")
	Magenta = color("\033[1;35m%s\033[0m")
	Teal    = color("\033[1;36m%s\033[0m")
	White   = color("\033[1;37m%s\033[0m")

	Blackf   = colorf("\033[1;30m%s\033[0m")
	Redf     = colorf("\033[1;31m%s\033[0m")
	Greenf   = colorf("\033[1;32m%s\033[0m")
	Yellowf  = colorf("\033[1;33m%s\033[0m")
	Purplef  = colorf("\033[1;34m%s\033[0m")
	Magentaf = colorf("\033[1;35m%s\033[0m")
	Tealf    = colorf("\033[1;36m%s\033[0m")
	Whitef   = colorf("\033[1;37m%s\033[0m")
)

func color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func colorf(colorString string) func(string, ...interface{}) string {
	sprint := func(format string, args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprintf(format, args...))
	}
	return sprint
}

type Logger struct {
	Output    io.Writer
	prfx      []string
	startTime time.Time
}

var Default Logger = Logger{
	Output:    os.Stdout,
	prfx:      nil,
	startTime: time.Now(),
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

func (receiver Logger) Begin(a ...interface{}) Logger {
	startTime = time.Now()
	receiver.output("BEGIN")
	logger := Logger{
		Output:    receiver.Output,
		prfx:      receiver.prfx,
		startTime: startTime,
	}
	return logger
}

func (receiver Logger) End(a ...interface{}) {
	endTime = time.Now()
	receiver.output("END", endTime.Sub(startTime))
}

func (receiver Logger) Prefix(newprefix ...string) Logger {
	logger := Logger{
		prfx:   newprefix,
		Output: Default.Output,
	}
	return logger
}

func (receiver Logger) Alert(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(Green(a...))
}

func (receiver Logger) Error(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(Red(a...))
}

func (receiver Logger) Highlight(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(Teal(a...))
}

func (receiver Logger) Inform(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(Magenta(a...))
}

func (receiver Logger) Trace(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(a...)
}

func (receiver Logger) Warn(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(Yellow(a...))
}

func (receiver Logger) output(a ...interface{}) {
	if receiver.prfx != nil {
		var prefixStr string = strings.Join(receiver.prfx[:], ": ") + ":"
		a = append([]interface{}{prefixStr}, a...)
	}
	var functionName string = ""
	{
		var pc uintptr
		var ok bool = true
		var found bool = false
		skip := 0
		for !found && ok {
			pc, _, _, ok = runtime.Caller(skip)
			if ok {
				fnc := runtime.FuncForPC(pc)
				foundedFunctionName := fnc.Name()
				if !strings.Contains(foundedFunctionName, "go-log") {
					functionName = foundedFunctionName
					found = true
				}
			}
			skip++
			if skip > 100 {
				found = true
			}
		}
	}
	a = append([]interface{}{functionName}, a...)
	fmt.Fprintln(receiver.Output, a...)
}

func (receiver Logger) outputf(format string, a ...interface{}) {
	var prefix strings.Builder
	{
		var pc uintptr
		var ok bool = true
		var found bool = false
		skip := 0
		var functionName string = ""
		for !found && ok {
			pc, _, _, ok = runtime.Caller(skip)
			if ok {
				fnc := runtime.FuncForPC(pc)
				foundedFunctionName := fnc.Name()
				if !strings.Contains(foundedFunctionName, "go-log") {
					functionName = foundedFunctionName
					found = true
				}
			}
			skip++
			if skip > 100 {
				found = true
			}
		}
		if functionName != "" {
			prefix.WriteString(functionName)
			prefix.WriteString(": ")
		}
	}
	if receiver.prfx != nil {
		var prefixStr string = strings.Join(receiver.prfx[:], ": ") + ": "
		prefix.WriteString(prefixStr)
	}
	format = fmt.Sprintf("%s%s%s", prefix.String(), format, "\n")
	fmt.Fprintf(receiver.Output, format, a...)
}

func (receiver Logger) Alertf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(Greenf(format, a...))
}

func (receiver Logger) Errorf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(Redf(format, a...))
}

func (receiver Logger) Highlightf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(Tealf(format, a...))
}

func (receiver Logger) Informf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(Magentaf(format, a...))
}

func (receiver Logger) Tracef(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(format, a...)
}

func (receiver Logger) Warnf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(Yellowf(format, a...))
}

func Log(a ...interface{}) {
	Default.Log(a...)
}

func Alert(a ...interface{}) {
	Default.Alert(a...)
}

func Error(a ...interface{}) {
	Default.Error(a...)
}

func Highlight(a ...interface{}) {
	Default.Highlight(a...)
}

func Inform(a ...interface{}) {
	Default.Inform(a...)
}

func Trace(a ...interface{}) {
	Default.Trace(a...)
}

func Warn(a ...interface{}) {
	Default.Warn(a...)
}

func Logf(format string, a ...interface{}) {
	Default.Logf(format, a...)
}

func Alertf(format string, a ...interface{}) {
	Default.Alertf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	Default.Errorf(format, a...)
}

func Highlightf(format string, a ...interface{}) {
	Default.Highlightf(format, a...)
}

func Informf(format string, a ...interface{}) {
	Default.Informf(format, a...)
}

func Tracef(format string, a ...interface{}) {
	Default.Tracef(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Default.Warnf(format, a...)
}

func Begin(a ...interface{}) Logger {
	return Default.Begin(a...)
}

func End(a ...interface{}) {
	Default.End(a...)
}
