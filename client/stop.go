package client

import (
	"fmt"
	"io"
	"strings"
)

// StopOptions represents the command line options for the `docker compose stop` command.
//
// https://docs.docker.com/compose/reference/stop/
type StopOptions struct {
	// Services to start
	Services []string

	// Specify a shutdown timeout in seconds
	Timeout *int
}

func stopFlags(opts *StopOptions) string {
	flags := ""

	if opts != nil {
		if opts.Timeout != nil {
			flags = fmt.Sprintf("%s --timeout %d", flags, *opts.Timeout)
		}

		for _, service := range opts.Services {
			flags = fmt.Sprintf("%s %s", flags, service)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose stop
//
// Stop services
//
// stderr is written to the given io.Writer
//
// https://docs.docker.com/compose/reference/stop/
func (c *ComposeClient) Stop(opts *StopOptions, w io.Writer, overrides ...*GlobalOptions) (<-chan error, error) {
	return c.RunCommand("stop", stopFlags(opts), nil, w, overrides...)
}
