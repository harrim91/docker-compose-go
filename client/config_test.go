package client_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
	"github.com/stretchr/testify/mock"
)

const configJSON string = "{\"services\": []}"

type mockConfigCmd struct {
	mock.Mock
	stdout io.Writer
	stderr io.Writer
}

func (o *mockConfigCmd) SetStdout(stdout io.Writer) {
	o.Called(stdout)
	o.stdout = stdout
}

func (o *mockConfigCmd) SetStderr(stderr io.Writer) {
	o.Called(stderr)
	o.stderr = stderr
}

func (o *mockConfigCmd) Run(cmd string) (<-chan error, error) {
	o.Called(cmd)

	if strings.Contains(cmd, runErrFlag) {
		return nil, errors.New(runErrFlag)
	}

	ch := make(chan error)

	go func() {
		if o.stdout != nil {
			o.stdout.Write([]byte(configJSON))
		}

		ch <- nil
	}()

	return ch, nil
}

func TestConfigCommand(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", "docker compose config --format json")

	config, err := c.Config(nil)

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)

	if string(config) != configJSON {
		t.Errorf("expected: %s, got: %s", configJSON, string(config))
	}
}

func TestConfigCommandResolveImageDigests(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", "docker compose config --format json --resolve-image-digests")

	_, err := c.Config(&client.ConfigOptions{
		ResolveImageDigests: true,
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}

func TestConfigCommandNoInterpolate(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)
	cmd.On("Run", "docker compose config --format json --no-interpolate")

	_, err := c.Config(&client.ConfigOptions{
		NoInterpolate: true,
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}

func TestConfigCommandQuiet(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)

	cmd.On("Run", "docker compose config --format json --quiet")

	_, err := c.Config(&client.ConfigOptions{
		Quiet: true,
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}

func TestConfigCommandServices(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)

	cmd.On("Run", "docker compose config --format json --services")

	_, err := c.Config(&client.ConfigOptions{
		Services: true,
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}

func TestConfigCommandVolumes(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)

	cmd.On("Run", "docker compose config --format json --volumes")

	_, err := c.Config(&client.ConfigOptions{
		Volumes: true,
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}

func TestConfigCommandHash(t *testing.T) {
	cmd := &mockConfigCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("SetStdout", mock.Anything)

	cmd.On("Run", "docker compose config --format json --hash=\"foo,bar,baz\"")

	_, err := c.Config(&client.ConfigOptions{
		Hash: "foo,bar,baz",
	})

	if err != nil {
		t.Error(err)
	}

	cmd.AssertExpectations(t)
}
