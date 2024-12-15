package xargs

import (
	"bytes"
	"testing"
)

func TestExecMyArgs(t *testing.T) {
	t.Run("test just command without args", func(t *testing.T) {
		buffer := bytes.Buffer{}

		xargs := Xargs{
			command: "ls",
			args:    []string{},
		}

		xargs.Exec(&buffer)
		got := buffer.String()
		want := `xargs.go
xargs_test.go
`

		if got != want {
			t.Errorf("\ngot:%q\nwant:%q\n", got, want)
		}
	})

	t.Run("test just command with args", func(t *testing.T) {
		buffer := bytes.Buffer{}

		xargs := Xargs{
			command: "ls",
			args:    []string{"-a"},
		}

		xargs.Exec(&buffer)
		got := buffer.String()
		want := ".\n..\nxargs.go\nxargs_test.go\n"

		if got != want {
			t.Errorf("\ngot:%q\nwant:%q\n", got, want)
		}
	})
}
