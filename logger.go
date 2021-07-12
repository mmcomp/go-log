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
	Output io.Writer
	Prfx   []string
}

var Default Logger = Logger{
	Output: os.Stdout,
	Prfx:   nil,
}

var startTime time.Time
var endTime time.Time

func (receiver Logger) log(a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.output(a...)
}

func (receiver Logger) logf(format string, a ...interface{}) {
	if receiver.Output == nil {
		return
	}
	receiver.outputf(format, a...)
}

func (receiver Logger) begin(a ...interface{}) {
	startTime = time.Now()
	receiver.output("BEGIN")
}

func (receiver Logger) end(a ...interface{}) {
	endTime = time.Now()
	receiver.output("END", endTime.Sub(startTime))
}

func (receiver Logger) prefix(newprefix ...string) Logger {
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
	if receiver.Prfx != nil {
		var prefixStr string = strings.Join(receiver.Prfx[:], ": ") + ": "
		prefix.WriteString(prefixStr)
	}
	format = fmt.Sprintf("%s%s%s", prefix.String(), format, "\n")
	fmt.Fprintf(receiver.Output, format, a...)
}

func Log(a ...interface{}) {
	Default.log(a...)
}

func Alert(a ...interface{}) {
	// TODO use alert method
	Default.log(Green(a...))
}

func Error(a ...interface{}) {
	Default.log(Red(a...))
}

func Highlight(a ...interface{}) {
	Default.log(Teal(a...))
}

func Inform(a ...interface{}) {
	Default.log(Magenta(a...))
}

func Trace(a ...interface{}) {
	Default.log(a...)
}

func Warn(a ...interface{}) {
	Default.log(Yellow(a...))
}

func Logf(format string, a ...interface{}) {
	Default.logf(format, a...)
}

func Alertf(format string, a ...interface{}) {
	Default.logf(Greenf(format, a...))
}

func Errorf(format string, a ...interface{}) {
	Default.logf(Redf(format, a...))
}

func Highlightf(format string, a ...interface{}) {
	Default.logf(Tealf(format, a...))
}

func Informf(format string, a ...interface{}) {
	Default.logf(Magentaf(format, a...))
}

func Tracef(format string, a ...interface{}) {
	Default.logf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Default.logf(Yellowf(format, a...))
}

func Begin(a ...interface{}) Logger {
	Default.begin(a...)
	return Default
}

func End(a ...interface{}) {
	Default.end(a...)
}
