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
	logger.Log("test")
	out := b.String()
	if out != "test\n" {
		t.Errorf("`test` was expected but got %q", out)
	}
}

func TestLogger_Logf(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.Logf("%s", "testf\n")
	out := b.String()
	if out != "testf\n" {
		t.Errorf("`testf` was expected but got %q", out)
	}
}

func TestLogger_Prefix_Log(t *testing.T) {
	var b bytes.Buffer
	logger := Default.Prefix("Pr1", "Pr2")
	logger.Output = &b
	logger.Log("test")
	out := b.String()
	if out != "Pr1: Pr2: test\n" {
		t.Errorf("`Pr1: Pr2: test` was expected but got %q", out)
	}
}
