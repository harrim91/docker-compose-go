package client_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
	"github.com/stretchr/testify/mock"
)

const (
	stdoutMsg      string = "hello from stdout"
	stderrMsg      string = "hello from stderr"
	errCommand     string = "test-err"
	runErrFlag     string = "run-err"
	processErrFlag string = "process-err"
)

type MockCmd struct {
	mock.Mock
	Stdout io.Writer
	Stderr io.Writer
}

func (o *MockCmd) SetStdout(stdout io.Writer) {
	o.Called(stdout)
	o.Stdout = stdout
}

func (o *MockCmd) SetStderr(stderr io.Writer) {
	o.Called(stderr)
	o.Stderr = stderr
}

func (o *MockCmd) Run(cmd string) (<-chan error, error) {
	o.Called(cmd)

	if strings.Contains(cmd, runErrFlag) {
		return nil, errors.New(runErrFlag)
	}

	ch := make(chan error)

	go func() {
		if o.Stdout != nil {
			o.Stdout.Write([]byte(stdoutMsg))
		}

		if o.Stderr != nil {
			o.Stderr.Write([]byte(stderrMsg))
		}

		if strings.Contains(cmd, processErrFlag) {
			ch <- errors.New(processErrFlag)
			return
		}

		ch <- nil
	}()

	return ch, nil
}

func TestRunsCommandWithFlags(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientConfigFiles(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Files: []string{
				"file1",
				"file2",
			},
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --file file1 --file file2 foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideFiles(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Files: []string{
				"file1",
				"file2",
			},
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --file file1 --file file2 --file file3 foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		Files: []string{
			"file3",
		},
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigProfiles(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Profiles: []string{
				"profile1",
				"profile2",
			},
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --profile profile1 --profile profile2 foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideProfiles(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Profiles: []string{
				"profile1",
				"profile2",
			},
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --profile profile1 --profile profile2 --profile profile3 foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		Profiles: []string{
			"profile3",
		},
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigProjectName(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			ProjectName: "my-project",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --project-name my-project foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideProjectName(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			ProjectName: "my-project",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --project-name override-project-name foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		ProjectName: "override-project-name",
	})

	cmd.AssertExpectations(t)
}
func TestClientConfigVerbose(t *testing.T) {
	cmd := &MockCmd{}

	v := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Verbose: &v,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --verbose foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideVerbose(t *testing.T) {
	cmd := &MockCmd{}

	v := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Verbose: &v,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	ov := false

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		Verbose: &ov,
	})

	cmd.AssertExpectations(t)
}
func TestClientConfigNoANSI(t *testing.T) {
	cmd := &MockCmd{}

	noAnsi := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			NoANSI: &noAnsi,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --no-ansi foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideNoANSI(t *testing.T) {
	cmd := &MockCmd{}

	noAnsi := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			NoANSI: &noAnsi,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	override := false

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		NoANSI: &override,
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigHost(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Host: "my-host",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --host my-host foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideHost(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Host: "my-host",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --host override-host foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		Host: "override-host",
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigTLS(t *testing.T) {
	cmd := &MockCmd{}

	tls := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLS: &tls,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tls foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideTLS(t *testing.T) {
	cmd := &MockCmd{}

	tls := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLS: &tls,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	override := false

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		TLS: &override,
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigTLSCACert(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSCACert: "my-tls-ca-cert",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlscacert my-tls-ca-cert foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideTLSCACert(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSCACert: "my-tls-ca-cert",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlscacert override-tls-ca-cert foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		TLSCACert: "override-tls-ca-cert",
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigTLSCert(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSCert: "my-tls-cert",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlscert my-tls-cert foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideTLSCert(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSCert: "my-tls-cert",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlscert override-tls-cert foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		TLSCert: "override-tls-cert",
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigTLSKey(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSKey: "my-tls-key",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlskey my-tls-key foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideTLSKey(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSKey: "my-tls-key",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlskey override-tls-key foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		TLSKey: "override-tls-key",
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigTLSVerify(t *testing.T) {
	cmd := &MockCmd{}

	tlsVerify := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSVerify: &tlsVerify,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --tlsverify foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideTLSVerify(t *testing.T) {
	cmd := &MockCmd{}

	tlsVerify := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			TLSVerify: &tlsVerify,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	override := false

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		TLSVerify: &override,
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigProjectDirectory(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			ProjectDirectory: "my-project-directory",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --project-directory my-project-directory foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideProjectDirectory(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			ProjectDirectory: "my-project-directory",
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --project-directory override-project-directory foo bar")

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		ProjectDirectory: "override-project-directory",
	})

	cmd.AssertExpectations(t)
}

func TestClientConfigCompatibility(t *testing.T) {
	cmd := &MockCmd{}

	compatibility := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Compatibility: &compatibility,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose --compatibility foo bar")

	c.RunCommand("foo", "bar", nil, nil)

	cmd.AssertExpectations(t)
}

func TestClientOverrideCompatibility(t *testing.T) {
	cmd := &MockCmd{}

	compatibility := true

	c := &client.ComposeClient{
		GlobalOptions: &client.GlobalOptions{
			Compatibility: &compatibility,
		},
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose foo bar")

	override := false

	c.RunCommand("foo", "bar", nil, nil, &client.GlobalOptions{
		Compatibility: &override,
	})

	cmd.AssertExpectations(t)
}

func TestRunCommandStdoutWriter(t *testing.T) {
	cmd := &MockCmd{}

	var buff bytes.Buffer

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", &buff)
	cmd.On("Run", "docker compose foo bar")

	c.RunCommand("foo", "bar", &buff, nil)

	cmd.AssertExpectations(t)
}

func TestRunCommandStderrWriter(t *testing.T) {
	cmd := &MockCmd{}

	var buff bytes.Buffer

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStderr", &buff)
	cmd.On("Run", "docker compose foo bar")

	c.RunCommand("foo", "bar", nil, &buff)

	cmd.AssertExpectations(t)
}

func TestRunCommandRunError(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", mock.Anything)

	_, err := c.RunCommand(errCommand, runErrFlag, nil, nil)

	if err == nil || err.Error() != runErrFlag {
		t.Errorf("expected error %s, got %v", runErrFlag, err)
	}
}

func TestRunCommandProcessError(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", mock.Anything)

	ch, err := c.RunCommand(errCommand, processErrFlag, nil, nil)

	if err != nil {
		t.Error(err)
	}

	err = <-ch

	if err == nil || err.Error() != processErrFlag {
		t.Errorf("expected error %s, got %v", processErrFlag, err)
	}
}

func TestRunQuery(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", "docker compose foo bar --format json")

	res, err := c.RunQuery("foo", "bar")

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)

	if string(res) != stdoutMsg {
		t.Errorf("expected: %s, got: %s", stdoutMsg, string(res))
	}
}

func TestRunQueryRunError(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", mock.Anything)

	_, err := c.RunQuery(errCommand, runErrFlag)

	if err == nil || err.Error() != runErrFlag {
		t.Errorf("expected error %s, got %v", runErrFlag, err)
	}
}

func TestRunQueryProcessError(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", mock.Anything)

	_, err := c.RunQuery(errCommand, processErrFlag)

	if err == nil || err.Error() != processErrFlag {
		t.Errorf("expected error %s, got %v", processErrFlag, err)
	}
}
