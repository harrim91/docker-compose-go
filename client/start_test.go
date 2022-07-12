package client_test

import (
	"bytes"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
)

func TestStartCommand(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose start")

	c.Start(nil, nil)

	cmd.AssertExpectations(t)
}

func TestStartCommandServices(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose start foo bar baz")

	c.Start(&client.StartOptions{
		Services: []string{"foo", "bar", "baz"},
	}, nil)

	cmd.AssertExpectations(t)
}

func TestStartCommandIOWriter(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	var buff bytes.Buffer

	cmd.On("Run", "docker compose start")

	// Start command writes to Stderr
	cmd.On("SetStderr", &buff)

	c.Start(nil, &buff)

	cmd.AssertExpectations(t)
}
