package client_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
	"github.com/stretchr/testify/mock"
)

const (
	version     string = "v2.3.3"
	versionJSON string = "{\"version\": \"v2.3.3\"}"
)

type mockVersionCmd struct {
	mock.Mock
	stdout io.Writer
	stderr io.Writer
}

func (o *mockVersionCmd) SetStdout(stdout io.Writer) {
	o.Called(stdout)
	o.stdout = stdout
}

func (o *mockVersionCmd) SetStderr(stderr io.Writer) {
	o.Called(stderr)
	o.stderr = stderr
}

func (o *mockVersionCmd) Run(cmd string) (<-chan error, error) {
	o.Called(cmd)

	if strings.Contains(cmd, runErrFlag) {
		return nil, errors.New(runErrFlag)
	}

	ch := make(chan error)

	go func() {
		if o.stdout != nil {
			o.stdout.Write([]byte(versionJSON))
		}

		ch <- nil
	}()

	return ch, nil
}

func TestVersion(t *testing.T) {
	cmd := &mockVersionCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", "docker compose version --format json")

	res, err := c.Version()

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)

	if res.Version != version {
		t.Errorf("expected: %s, got: %s", version, res)
	}
}
