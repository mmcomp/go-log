package log

import (
	"bytes"
	"testing"
)

func TestLogger_Log(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}

	tests := []struct {
		Input    []interface{}
		Excepted string
	}{
		{
			Input:    []interface{}{"test"},
			Excepted: "[test]\n",
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen"},
			Excepted: "[sometimes it is best to listen]\n",
		},

		{
			Input:    []interface{}{"Hello", "!!"},
			Excepted: "[Hello !!]\n",
		},
	}

	for testNumber, test := range tests {
		logger.Log(test.Input)
		out := b.String()
		if out != test.Excepted {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Excepted, out)
		}
		b.Reset()
	}
}

func TestLogger_Logf(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}

	tests := []struct {
		Input    []interface{}
		Excepted string
	}{
		{
			Input:    []interface{}{"test\n"},
			Excepted: "[test\n]",
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen\n"},
			Excepted: "[sometimes it is best to listen\n]",
		},

		{
			Input:    []interface{}{"Hello", "!!\n"},
			Excepted: "[Hello !!\n]",
		},
	}

	for testNumber, test := range tests {
		logger.Logf("%s", test.Input)
		out := b.String()
		if out != test.Excepted {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Excepted, out)
		}
		b.Reset()
	}
}

func TestLogger_Prefix_Log(t *testing.T) {
	var b bytes.Buffer
	tests := []struct {
		Input    []interface{}
		Excepted string
		Prefixes []string
	}{
		{
			Input:    []interface{}{"test\n"},
			Excepted: "Pr1: Pr2: [test\n]\n",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"The", "test\n"},
			Excepted: "Pr1: [The test\n]\n",
			Prefixes: []string{"Pr1"},
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen\n"},
			Excepted: "Pr1: Pr2: [sometimes it is best to listen\n]\n",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"Hello", "!!\n"},
			Excepted: "Pr1: [Hello !!\n]\n",
			Prefixes: []string{"Pr1"},
		},
	}

	for testNumber, test := range tests {
		logger := Default.Prefix(test.Prefixes...)
		logger.Output = &b
		logger.Log(test.Input)
		out := b.String()
		if out != test.Excepted {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Excepted, out)
		}
		b.Reset()
	}
}

func TestLogger_Prefix_Logf(t *testing.T) {
	var b bytes.Buffer
	tests := []struct {
		Input    []interface{}
		Excepted string
		Prefixes []string
	}{
		{
			Input:    []interface{}{"test\n"},
			Excepted: "Pr1: Pr2: [test\n]",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"The", "test\n"},
			Excepted: "Pr1: [The test\n]",
			Prefixes: []string{"Pr1"},
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen\n"},
			Excepted: "Pr1: Pr2: [sometimes it is best to listen\n]",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"Hello", "!!\n"},
			Excepted: "Pr1: [Hello !!\n]",
			Prefixes: []string{"Pr1"},
		},
	}
	for testNumber, test := range tests {
		logger := Default.Prefix(test.Prefixes...)
		logger.Output = &b
		logger.Logf("%s", test.Input)
		out := b.String()
		if out != test.Excepted {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Excepted, out)
		}
		b.Reset()
	}
}

func TestLogger_FuncNameOne(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.LogWithFuncName("test")
	out := b.String()
	if out != "github.com/mmcomp/go-log.TestLogger_FuncNameOne: test\n" {
		t.Errorf("`github.com/mmcomp/go-log.TestLogger_FuncNameOne: test\n` was expected but got %q", out)
	}
}

func TestLogger_FuncNameTwo(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.LogWithFuncName("test")
	out := b.String()
	if out != "github.com/mmcomp/go-log.TestLogger_FuncNameTwo: test\n" {
		t.Errorf("`github.com/mmcomp/go-log.TestLogger_FuncNameTwo: test\n` was expected but got %q", out)
	}
}

func TestLogger_fFuncNameOne(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.LogfWithFuncName("%s\n", "test")
	out := b.String()
	if out != "github.com/mmcomp/go-log.TestLogger_fFuncNameOne: test\n" {
		t.Errorf("`github.com/mmcomp/go-log.TestLogger_fFuncNameOne: test\n` was expected but got %q", out)
	}
}

func TestLogger_fFuncNameTwo(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.LogfWithFuncName("%s\n", "test")
	out := b.String()
	if out != "github.com/mmcomp/go-log.TestLogger_fFuncNameTwo: test\n" {
		t.Errorf("`github.com/mmcomp/go-log.TestLogger_fFuncNameTwo: test\n` was expected but got %q", out)
	}
}
