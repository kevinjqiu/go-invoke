package go_invoke

import (
	"io"
	"os"
	"sigs.k8s.io/kind/pkg/exec"
)

type options struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Writer
	envs   []string
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

type CommandInvoker struct{}

func (ci *CommandInvoker) Run(cmd string, args []string, opts ...Option) ([]byte, error) {
	o := options{
		stdout: os.Stdout,
		stdin:  os.Stdin,
		stderr: os.Stderr,
		envs:   []string{},
	}

	for _, opt := range opts {
		opt.apply(&o)
	}

	c := exec.Command(cmd, args...)
	c.SetEnv()
	return []byte{}, nil
}
