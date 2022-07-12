package client_test

import (
	"bytes"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
)

func TestStopCommand(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose stop")

	c.Stop(nil, nil)

	cmd.AssertExpectations(t)
}

func TestStopCommandTimeout(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose stop --timeout 7")

	timeout := 7

	c.Stop(&client.StopOptions{
		Timeout: &timeout,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestStopCommandServices(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose stop foo bar baz")

	c.Stop(&client.StopOptions{
		Services: []string{"foo", "bar", "baz"},
	}, nil)

	cmd.AssertExpectations(t)
}

func TestStopCommandIOWriter(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	var buff bytes.Buffer

	cmd.On("Run", "docker compose stop")

	// Stop command writes to Stderr
	cmd.On("SetStderr", &buff)

	c.Stop(nil, &buff)

	cmd.AssertExpectations(t)
}
