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
			Excepted: "testing.tRunner test\n",
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen"},
			Excepted: "testing.tRunner sometimes it is best to listen\n",
		},

		{
			Input:    []interface{}{"Hello", "!!"},
			Excepted: "testing.tRunner Hello !!\n",
		},
	}

	for testNumber, test := range tests {
		logger.Log(test.Input...)
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
		Format   string
	}{
		{
			Input:    []interface{}{"test"},
			Excepted: "testing.tRunner test \n",
			Format:   "%s %s %s",
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen"},
			Excepted: "testing.tRunner sometimes it is best to listen \n",
			Format:   "%s %s %s %s %s %s %s %s",
		},

		{
			Input:    []interface{}{"Hello", "!!"},
			Excepted: "testing.tRunner Hello !! \n",
			Format:   "%s %s %s %s",
		},
	}

	for testNumber, test := range tests {
		logger.Logf(test.Format, test.Input...)
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
			Input:    []interface{}{"test"},
			Excepted: "testing.tRunner Pr1: Pr2: test\n",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"The", "test"},
			Excepted: "testing.tRunner Pr1: The test\n",
			Prefixes: []string{"Pr1"},
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen"},
			Excepted: "testing.tRunner Pr1: Pr2: sometimes it is best to listen\n",
			Prefixes: []string{"Pr1", "Pr2"},
		},

		{
			Input:    []interface{}{"Hello", "!!"},
			Excepted: "testing.tRunner Pr1: Hello !!\n",
			Prefixes: []string{"Pr1"},
		},
	}

	for testNumber, test := range tests {
		logger := Default.Prefix(test.Prefixes...)
		logger.Output = &b
		logger.Log(test.Input...)
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
		Format   string
	}{
		{
			Input:    []interface{}{"test"},
			Excepted: "testing.tRunner Pr1: Pr2: test \n",
			Prefixes: []string{"Pr1", "Pr2"},
			Format:   "%s %s %s %s",
		},

		{
			Input:    []interface{}{"The", "test"},
			Excepted: "testing.tRunner Pr1: The test \n",
			Prefixes: []string{"Pr1"},
			Format:   "%s %s %s %s %s",
		},

		{
			Input:    []interface{}{"sometimes", "it", "is", "best", "to", "listen"},
			Excepted: "testing.tRunner Pr1: Pr2: sometimes it is best to listen \n",
			Prefixes: []string{"Pr1", "Pr2"},
			Format:   "%s %s %s %s %s %s %s %s %s",
		},

		{
			Input:    []interface{}{"Hello", "!!"},
			Excepted: "testing.tRunner Pr1: Hello !! \n",
			Prefixes: []string{"Pr1"},
			Format:   "%s %s %s %s %s",
		},
	}
	for testNumber, test := range tests {
		logger := Default.Prefix(test.Prefixes...)
		logger.Output = &b
		logger.Logf(test.Format, test.Input...)
		out := b.String()
		if out != test.Excepted {
			t.Errorf("Test %d :  %q was expected but got %q", testNumber, test.Excepted, out)
		}
		b.Reset()
	}
}
