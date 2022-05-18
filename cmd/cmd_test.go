package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/harrim91/docker-compose-go/cmd"
)

const (
	stdoutMessage = "stdout message"
	stderrMessage = "stderr message"
)

const (
	successExitCode = 0
	errExitCode     = 1
)

func TestCommandRunSuccess(t *testing.T) {
	cmd := cmd.Cmd{
		Exec: func(command string, args ...string) *exec.Cmd {
			cs := []string{"-test.run=TestShellProcessSuccess", "--", command}
			cs = append(cs, args...)
			cmd := exec.Command(os.Args[0], cs...)
			cmd.Env = []string{"GO_TEST_PROCESS=1"}
			return cmd
		},
	}

	ch, err := cmd.Run("echo hello")

	if err != nil {
		t.Errorf("expected no error from cmd.Run, got: %v", err)
		return
	}

	err = <-ch

	if err != nil {
		t.Errorf("expected no error from channel, got: %v", err)
		return
	}
}

func TestCommandRunError(t *testing.T) {
	cmd := cmd.Cmd{
		Exec: func(command string, args ...string) *exec.Cmd {
			cs := []string{"-test.run=TestShellProcessError", "--", command}
			cs = append(cs, args...)
			cmd := exec.Command(os.Args[0], cs...)
			cmd.Env = []string{"GO_TEST_PROCESS=1"}
			return cmd
		},
	}

	ch, err := cmd.Run("echo hello")

	if err != nil {
		t.Error(err)
		return
	}

	err = <-ch

	if err == nil {
		t.Errorf("expected error, got: %v", err)
		return
	}

	expected := "exit status 1"

	if err.Error() != expected {
		t.Errorf("expected error '%s', got: '%v'", expected, err)
		return
	}
}

func TestCommandRunInput(t *testing.T) {
	input := "echo hello"

	cmd := cmd.Cmd{
		Exec: func(command string, args ...string) *exec.Cmd {
			expectedCommand := "/bin/sh"

			if command != expectedCommand {
				t.Errorf("expected command '%s', got: %v", expectedCommand, command)
			}

			expectedArg0 := "-c"

			if args[0] != expectedArg0 {
				t.Errorf("expected args[0] '%s', got: %v", expectedArg0, args[0])
			}

			if args[1] != input {
				t.Errorf("expected args[1] '%s', got: %v", input, args[1])
			}

			cs := []string{"-test.run=TestShellProcessSuccess", "--", command}
			cs = append(cs, args...)
			cmd := exec.Command(os.Args[0], cs...)
			cmd.Env = []string{"GO_TEST_PROCESS=1"}
			return cmd
		},
	}

	ch, err := cmd.Run(input)

	if err != nil {
		t.Errorf("expected no error from cmd.Run, got: %v", err)
		return
	}

	<-ch
}

func TestCommandStdout(t *testing.T) {
	cmd := cmd.Cmd{
		Exec: func(command string, args ...string) *exec.Cmd {
			cs := []string{"-test.run=TestShellProcessSuccess", "--", command}
			cs = append(cs, args...)
			cmd := exec.Command(os.Args[0], cs...)
			cmd.Env = []string{"GO_TEST_PROCESS=1"}
			return cmd
		},
	}

	var buff bytes.Buffer

	cmd.SetStdout(&buff)

	ch, err := cmd.Run("echo hello")

	if err != nil {
		t.Errorf("expected no error from cmd.Run, got: %v", err)
		return
	}

	<-ch

	got := buff.String()
	want := stdoutMessage

	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestCommandStderr(t *testing.T) {
	cmd := cmd.Cmd{
		Exec: func(command string, args ...string) *exec.Cmd {
			cs := []string{"-test.run=TestShellProcessError", "--", command}
			cs = append(cs, args...)
			cmd := exec.Command(os.Args[0], cs...)
			cmd.Env = []string{"GO_TEST_PROCESS=1"}
			return cmd
		},
	}

	var buff bytes.Buffer

	cmd.SetStderr(&buff)

	ch, err := cmd.Run("echo hello")

	if err != nil {
		t.Errorf("expected no error from cmd.Run, got: %v", err)
		return
	}

	<-ch

	got := buff.String()
	want := stderrMessage

	if got != want {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

// TestShellProcessSuccess is a method that is called as a substitute for a shell command.
// It writes a predetermined message to STDOUT and returns an exit code of 0
// The GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is skipped.
func TestShellProcessSuccess(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	fmt.Fprint(os.Stdout, stdoutMessage)

	os.Exit(successExitCode)
}

// TestShellProcessError is a method that is called as a substitute for a shell command.
// It writes a predetermined message to STDOUT and returns an exit code of 0
// The GO_TEST_PROCESS flag ensures that if it is called as part of the test suite, it is skipped.
func TestShellProcessError(t *testing.T) {
	if os.Getenv("GO_TEST_PROCESS") != "1" {
		return
	}

	fmt.Fprint(os.Stderr, stderrMessage)

	os.Exit(errExitCode)
}
