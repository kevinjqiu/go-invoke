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
			Name: "default options with env vars",
			Cmd: "sh",
			Args: []string{"-c", "echo hello $YOURNAME"},
			Options: []Option{WithEnvs([]string{"YOURNAME=john"})},
			Stdout: "hello john\n",
			Stderr: "",
			err: nil,
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
