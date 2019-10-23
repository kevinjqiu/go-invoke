package go_invoke

import (
	"io"
	"io/ioutil"
)

type options struct {
	stdout    io.Writer
	stderr    io.Writer
	stdin     io.Reader
	envs      []string
	teeStdout bool
	teeStderr bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithEnvs(envs []string) Option {
	return optionFunc(func(o *options) {
		o.envs = envs
	})
}

func WithStdout(writer io.Writer) Option {
	return optionFunc(func(o *options) {
		o.stdout = writer
	})
}

func WithSuppressStdout() Option {
	return optionFunc(func(o *options) {
		o.stdout = ioutil.Discard
	})
}

func WithSuppressStderr() Option {
	return optionFunc(func(o *options) {
		o.stderr = ioutil.Discard
	})
}

func WithSupressOutput() Option {
	return optionFunc(func(o *options) {
		o.stdout = ioutil.Discard
		o.stderr = ioutil.Discard
	})
}

func WithTeeStdout(teeStdout bool) Option {
	return optionFunc(func(o *options) {
		o.teeStdout = teeStdout
	})
}

func WithTeeStderr(teeStderr bool) Option {
	return optionFunc(func(o *options) {
		o.teeStderr = teeStderr
	})
}

