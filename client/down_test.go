package client_test

import (
	"testing"

	"github.com/harrim91/docker-compose-go/client"
)

func TestDownCommand(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down")

	c.Down(nil, nil)

	cmd.AssertExpectations(t)
}

func TestDownRMILocal(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down --rmi local")

	c.Down(&client.DownOptions{
		RemoveImages: client.RemoveImageFlagLocal,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestDownRMIAll(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down --rmi all")

	c.Down(&client.DownOptions{
		RemoveImages: client.RemoveImageFlagAll,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestDownVolumes(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down --volumes")

	c.Down(&client.DownOptions{
		Volumes: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestDownRemoveOrphans(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down --remove-orphans")

	c.Down(&client.DownOptions{
		RemoveOrphans: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestDownTimeout(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose down --timeout 100")

	timeout := 100

	c.Down(&client.DownOptions{
		Timeout: &timeout,
	}, nil)

	cmd.AssertExpectations(t)
}
