package cmd

import (
	"io"
	"os/exec"
)

type executor func(name string, arg ...string) *exec.Cmd

// Cmd is used for executing shell commands
type Cmd struct {
	Exec   executor
	stdout io.Writer
	stderr io.Writer
}

// New returns a new Cmd
func New() *Cmd {
	return &Cmd{
		Exec: exec.Command,
	}
}

// Sets the stdout writer
func (c *Cmd) SetStderr(stderr io.Writer) {
	c.stderr = stderr
}

// Sets the stderr writer
func (c *Cmd) SetStdout(stdout io.Writer) {
	c.stdout = stdout
}

// Run a shell command. The returned channel will emit a single message and then close once the command has completed.
func (c *Cmd) Run(command string) (<-chan error, error) {
	execcmd := c.Exec("/bin/sh", "-c", command)

	if c.stdout != nil {
		execcmd.Stdout = c.stdout
	}

	if c.stderr != nil {
		execcmd.Stderr = c.stderr
	}

	if err := execcmd.Start(); err != nil {
		return nil, err
	}

	ch := make(chan error)

	go func() {
		defer close(ch)
		ch <- execcmd.Wait()
	}()

	return ch, nil
}
