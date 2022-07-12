package client

import (
	"fmt"
	"io"
	"strings"
)

type RemoveImageFlag string

const (
	// Remove all images used by any service.
	RemoveImageFlagAll RemoveImageFlag = "all"

	// Remove only images that don't have a custom tag set by the `image` field.
	RemoveImageFlagLocal RemoveImageFlag = "local"
)

// DownOptions represents the command line options for the `docker compose down` command.
//
// https://docs.docker.com/compose/reference/down/
type DownOptions struct {
	// Remove images. Type must be one of:
	//
	// - all: Remove all images used by any service.
	//
	// - local: Remove only images that don't have a custom tag set by the `image` field.
	RemoveImages RemoveImageFlag

	// Remove named volumes declared in the `volumes` section of the Compose file and anonymous volumes attached to containers.
	Volumes bool

	// Remove containers for services not defined in the Compose file
	RemoveOrphans bool

	// Specify a shutdown timeout in seconds. (default: 10)
	Timeout *int
}

func downFlags(opts *DownOptions) string {
	flags := ""

	if opts != nil {
		if opts.RemoveImages != "" {
			flags = fmt.Sprintf("%s --rmi %s", flags, opts.RemoveImages)
		}

		if opts.Volumes {
			flags = fmt.Sprintf("%s --volumes", flags)
		}

		if opts.RemoveOrphans {
			flags = fmt.Sprintf("%s --remove-orphans", flags)
		}

		if opts.Timeout != nil {
			flags = fmt.Sprintf("%s --timeout %d", flags, *opts.Timeout)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose down
//
// Stops containers and removes containers, networks, volumes, and images created by `up`.
//
// stderr is written to the given io.Writer
//
// https://docs.docker.com/compose/reference/down/
func (client *ComposeClient) Down(opts *DownOptions, w io.Writer, overrides ...*GlobalOptions) (<-chan error, error) {
	return client.RunCommand("down", downFlags(opts), nil, w, overrides...)
}
