package client

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type BuildProgressFlag string

const (
	BuildProgressFlagAuto  BuildProgressFlag = "auto"
	BuildProgressFlagPlain BuildProgressFlag = "plain"
	BuildProgressFlagTTY   BuildProgressFlag = "tty"
)

// BuildOptions represents the command line options for the `docker compose build` command.
//
// https://docs.docker.com/compose/reference/build/
type BuildOptions struct {
	// Set build-time variables for services.
	BuildArgs map[string]string

	// Compress the build context using gzip.
	Compress bool

	// Always remove intermediate containers.
	ForceRemove bool

	// Set memory limit for the build container.
	Memory string

	// Do not use cache when building the image.
	NoCache bool

	// Do not remove intermediate containers after a successful build.
	NoRemove bool

	// Build images in parallel
	Parallel bool

	// Set type of progress output (`auto`, `plain`, `tty`).
	Progress BuildProgressFlag

	// Always attempt to pull a newer version of the image.
	Pull bool

	// Don't print anything to `STDOUT`.
	Quiet bool

	// Services to build
	Services []string
}

func buildFlags(opts *BuildOptions) string {
	flags := ""

	if opts != nil {
		// Sort build args by name for predictable testing
		buildArgKeys := make([]string, 0, len(opts.BuildArgs))

		for key := range opts.BuildArgs {
			buildArgKeys = append(buildArgKeys, key)
		}

		sort.Strings(buildArgKeys)

		for _, key := range buildArgKeys {
			flags = fmt.Sprintf("%s --build-arg %s=%s", flags, key, opts.BuildArgs[key])
		}

		if opts.Compress {
			flags = fmt.Sprintf("%s --compress", flags)
		}

		if opts.ForceRemove {
			flags = fmt.Sprintf("%s --force-rm", flags)
		}

		if opts.Memory != "" {
			flags = fmt.Sprintf("%s --memory %s", flags, opts.Memory)
		}

		if opts.NoCache {
			flags = fmt.Sprintf("%s --no-cache", flags)
		}

		if opts.NoRemove {
			flags = fmt.Sprintf("%s --no-rm", flags)
		}

		if opts.Parallel {
			flags = fmt.Sprintf("%s --parallel", flags)
		}

		if opts.Progress != "" {
			flags = fmt.Sprintf("%s --progress %s", flags, opts.Progress)
		}

		if opts.Pull {
			flags = fmt.Sprintf("%s --pull", flags)
		}

		if opts.Quiet {
			flags = fmt.Sprintf("%s --quiet", flags)
		}

		for _, service := range opts.Services {
			flags = fmt.Sprintf("%s %s", flags, service)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose build
//
// Build or rebuild services
//
// stdout is written to the given io.Writer
//
// https://docs.docker.com/compose/reference/build/
func (c *ComposeClient) Build(opts *BuildOptions, w io.Writer, overrides ...*GlobalOptions) (<-chan error, error) {
	return c.RunCommand("build", buildFlags(opts), w, nil, overrides...)
}
