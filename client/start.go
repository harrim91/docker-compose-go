package client

import (
	"fmt"
	"io"
	"strings"
)

// StartOptions represents the command line options for the `docker compose start` command.
//
// https://docs.docker.com/compose/reference/start/
type StartOptions struct {
	// Services to start
	Services []string
}

func startFlags(opts *StartOptions) string {
	flags := ""

	if opts != nil {
		for _, service := range opts.Services {
			flags = fmt.Sprintf("%s %s", flags, service)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose start
//
// Start services
//
// stderr is written to the given io.Writer
//
// https://docs.docker.com/compose/reference/start/
func (c *ComposeClient) Start(opts *StartOptions, w io.Writer, overrides ...*GlobalOptions) (<-chan error, error) {
	return c.RunCommand("start", startFlags(opts), nil, w, overrides...)
}
