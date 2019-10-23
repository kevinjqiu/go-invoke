package go_invoke

import (
	"bytes"
	"os"
	"sigs.k8s.io/kind/pkg/exec"
)

type CommandInvoker struct{}

func (ci *CommandInvoker) Run(cmd string, args []string, opts ...Option) ([]byte, []byte, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	o := options{
		stdout:    &stdout,
		stderr:    &stderr,
		stdin:     os.Stdin,
		envs:      []string{},
		teeStdout: false,
	}

	for _, opt := range opts {
		opt.apply(&o)
	}

	c := exec.Command(cmd, args...)
	c.SetEnv(o.envs...)
	c.SetStdin(o.stdin)
	c.SetStdout(o.stdout)
	c.SetStderr(o.stderr)

	err := c.Run()
	if err != nil {
		return stdout.Bytes(), stderr.Bytes(), err
	}

	if o.teeStdout {
		os.Stdout.Write(stdout.Bytes())
	}

	if o.teeStderr {
		os.Stderr.Write(stderr.Bytes())
	}

	return stdout.Bytes(), stderr.Bytes(), nil
}
