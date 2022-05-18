package client_test

import (
	"testing"

	"github.com/harrim91/docker-compose-go/client"
)

func TestUpCommand(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up")

	c.Up(nil, nil)

	cmd.AssertExpectations(t)
}

func TestUpDetached(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --detach")

	c.Up(&client.UpOptions{
		Detach: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpNoColor(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --no-color")

	c.Up(&client.UpOptions{
		NoColor: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpQuietPull(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --quiet-pull")

	c.Up(&client.UpOptions{
		QuietPull: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpNoDeps(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --no-deps")

	c.Up(&client.UpOptions{
		NoDeps: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpForceRecreate(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --force-recreate")

	c.Up(&client.UpOptions{
		ForceRecreate: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpAlwaysRecreateDeps(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --always-recreate-deps")

	c.Up(&client.UpOptions{
		AlwaysRecreateDeps: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpNoRecreate(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --no-recreate")

	c.Up(&client.UpOptions{
		NoRecreate: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpNoBuild(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --no-build")

	c.Up(&client.UpOptions{
		NoBuild: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpNoStart(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --no-start")

	c.Up(&client.UpOptions{
		NoStart: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpBuild(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --build")

	c.Up(&client.UpOptions{
		Build: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpAbortOnContainerExit(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --abort-on-container-exit")

	c.Up(&client.UpOptions{
		AbortOnContainerExit: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpAttachDependencies(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --attach-dependencies")

	c.Up(&client.UpOptions{
		AttachDependencies: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpTimeout(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --timeout 0")

	timeout := 0

	c.Up(&client.UpOptions{
		Timeout: &timeout,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpRenewAnonVolumes(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --renew-anon-volumes")

	c.Up(&client.UpOptions{
		RenewAnonVolumes: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpRemoveOrphans(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --remove-orphans")

	c.Up(&client.UpOptions{
		RemoveOrphans: true,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpExitCodeFrom(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up --exit-code-from my-service")

	c.Up(&client.UpOptions{
		ExitCodeFrom: "my-service",
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpScale(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	// services names are sorted alphabetically for testing predictability
	cmd.On("Run", "docker compose up --scale bar=2 --scale foo=1")

	c.Up(&client.UpOptions{
		Scale: map[string]int{
			"foo": 1,
			"bar": 2,
		},
	}, nil)

	cmd.AssertExpectations(t)
}

func TestUpServices(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose up foo bar baz")

	c.Up(&client.UpOptions{
		Services: []string{"foo", "bar", "baz"},
	}, nil)

	cmd.AssertExpectations(t)
}
