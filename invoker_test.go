package go_invoke

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Name    string
	Cmd     string
	Args    []string
	Options []Option
	Stdout  string
	Stderr  string
	err     error
}

func TestInvoker_Run(t *testing.T) {
	testCases := []testCase{
		{
			Name:   "default options",
			Cmd:    "echo",
			Args:   []string{"hello", "world"},
			Stdout: "hello world\n",
			Stderr: "",
			err:    nil,
		},
		{
			Name:   "echo to stderr",
			Cmd:    "sh",
			Args:   []string{"-c", "echo foo >&2"},
			Stdout: "",
			Stderr: "foo\n",
		},
		{
			Name:    "default options with env vars",
			Cmd:     "sh",
			Args:    []string{"-c", "echo hello $YOURNAME"},
			Options: []Option{WithEnvs([]string{"YOURNAME=john"})},
			Stdout:  "hello john\n",
			Stderr:  "",
			err:     nil,
		},
		{
			Name:    "suppress stdout",
			Cmd:     "echo",
			Args:    []string{"foobar"},
			Options: []Option{WithSuppressStdout()},
			Stdout:  "",
			Stderr:  "",
			err:     nil,
		},
		{
			Name:    "suppress stderr",
			Cmd:     "sh",
			Args:    []string{"-c", "echo foo >&2"},
			Options: []Option{WithSuppressStderr()},
			Stdout:  "",
			Stderr:  "",
		},
		{
			Name:    "suppress all",
			Cmd:     "sh",
			Args:    []string{"-c", "echo foo >&2"},
			Options: []Option{WithSupressOutput()},
			Stdout:  "",
			Stderr:  "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			invoker := CommandInvoker{}
			stdout, stderr, err := invoker.Run(tc.Cmd, tc.Args, tc.Options...)
			assert.Nil(t, err)
			assert.Equal(t, tc.Stdout, string(stdout))
			assert.Equal(t, tc.Stderr, string(stderr))
		})
	}
}
