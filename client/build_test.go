package client_test

import (
	"bytes"
	"testing"

	"github.com/harrim91/docker-compose-go/client"
)

func TestBuildCommand(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build")

	c.Build(nil, nil)

	cmd.AssertExpectations(t)
}
func TestBuildCommandBuildArgs(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	// Args are sorted alphabetically for testing predictability
	cmd.On("Run", "docker compose build --build-arg baz=qux --build-arg foo=bar")

	c.Build(&client.BuildOptions{
		BuildArgs: map[string]string{
			"foo": "bar",
			"baz": "qux",
		},
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandCompress(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --compress")

	compress := true

	c.Build(&client.BuildOptions{
		Compress: &compress,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandForceRemove(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --force-rm")

	forceRemove := true

	c.Build(&client.BuildOptions{
		ForceRemove: &forceRemove,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandMemory(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --memory 50M")

	c.Build(&client.BuildOptions{
		Memory: "50M",
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandNoCache(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --no-cache")

	noCache := true

	c.Build(&client.BuildOptions{
		NoCache: &noCache,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandNoRemove(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --no-rm")

	noRemove := true

	c.Build(&client.BuildOptions{
		NoRemove: &noRemove,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandParallel(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --parallel")

	parallel := true

	c.Build(&client.BuildOptions{
		Parallel: &parallel,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandProgressAuto(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --progress auto")

	c.Build(&client.BuildOptions{
		Progress: client.BuildProgressFlagAuto,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandProgressPlain(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --progress plain")

	c.Build(&client.BuildOptions{
		Progress: client.BuildProgressFlagPlain,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandProgressTTY(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --progress tty")

	c.Build(&client.BuildOptions{
		Progress: client.BuildProgressFlagTTY,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandPull(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --pull")

	pull := true

	c.Build(&client.BuildOptions{
		Pull: &pull,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandQuiet(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build --quiet")

	quiet := true

	c.Build(&client.BuildOptions{
		Quiet: &quiet,
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandServices(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	cmd.On("Run", "docker compose build foo bar baz")

	c.Build(&client.BuildOptions{
		Services: []string{"foo", "bar", "baz"},
	}, nil)

	cmd.AssertExpectations(t)
}

func TestBuildCommandIOWriter(t *testing.T) {
	cmd := &MockCmd{}

	c := &client.ComposeClient{
		NewCmd: func() client.Cmd {
			return cmd
		},
	}

	var buff bytes.Buffer

	cmd.On("Run", "docker compose build")

	// Build command writes to Stdout
	cmd.On("SetStdout", &buff)

	c.Build(nil, &buff)

	cmd.AssertExpectations(t)
}
