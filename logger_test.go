package log

import (
	"bytes"
	"testing"
)

func TestLog(t *testing.T) {
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

func TestLogf(t *testing.T) {
	var b bytes.Buffer
	logger := Logger{
		Output: &b,
	}
	logger.Logf("%s", "testf")
	out := b.String()
	if out != "testf" {
		t.Errorf("`testf` was expected but got %q", out)
	}
}
