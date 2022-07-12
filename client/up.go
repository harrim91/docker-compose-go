package client

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// UpOptions represents the command line options for the `docker compose up` command.
//
// https://docs.docker.com/compose/reference/up/
type UpOptions struct {
	// Detached mode: Run containers in the background, print new container names. Detached mode is incompatible with --abort-on-container-exit.
	Detach bool

	// Produce monochrome output.
	NoColor bool

	// Pull without printing progress information.
	QuietPull bool

	// Don't start linked services.
	NoDeps bool

	// Recreate containers even if their configuration and image haven't changed.
	ForceRecreate bool

	// Recreate dependent containers. Incompatible with --no-recreate.
	AlwaysRecreateDeps bool

	// If containers already exist, don't recreate them. Incompatible with --force-recreate and --renew-anon-volumes
	NoRecreate bool

	// Don't build an image, even if it's missing.
	NoBuild bool

	// Don't start the services after creating them.
	NoStart bool

	// Build images before starting containers.
	Build bool

	// Stops all containers if any container was stopped. Incompatible with --detach.
	AbortOnContainerExit bool

	// Attach to dependent containers.
	AttachDependencies bool

	// Use this timeout in seconds for container shutdown when attached or when containers are already running. (default: 10)
	Timeout *int

	// Recreate anonymous volumes instead of retrieving data from the previous containers.
	RenewAnonVolumes bool

	// Remove containers for services not defined in the Compose file.
	RemoveOrphans bool

	// Return the exit code of the selected service container. Implies --abort-on-container-exit.
	ExitCodeFrom string

	// Scale SERVICE to NUM instances. Overrides the `scale` setting in the Compose file if present.
	Scale map[string]int

	// Defines the services to interact with
	Services []string
}

func upFlags(opts *UpOptions) string {
	flags := ""

	if opts != nil {
		if opts.Detach {
			flags = fmt.Sprintf("%s --detach", flags)
		}

		if opts.NoColor {
			flags = fmt.Sprintf("%s --no-color", flags)
		}

		if opts.QuietPull {
			flags = fmt.Sprintf("%s --quiet-pull", flags)
		}

		if opts.NoDeps {
			flags = fmt.Sprintf("%s --no-deps", flags)
		}

		if opts.ForceRecreate {
			flags = fmt.Sprintf("%s --force-recreate", flags)
		}

		if opts.AlwaysRecreateDeps {
			flags = fmt.Sprintf("%s --always-recreate-deps", flags)
		}

		if opts.NoRecreate {
			flags = fmt.Sprintf("%s --no-recreate", flags)
		}

		if opts.NoBuild {
			flags = fmt.Sprintf("%s --no-build", flags)
		}

		if opts.NoStart {
			flags = fmt.Sprintf("%s --no-start", flags)
		}

		if opts.Build {
			flags = fmt.Sprintf("%s --build", flags)
		}

		if opts.AbortOnContainerExit {
			flags = fmt.Sprintf("%s --abort-on-container-exit", flags)
		}

		if opts.AttachDependencies {
			flags = fmt.Sprintf("%s --attach-dependencies", flags)
		}

		if opts.Timeout != nil {
			flags = fmt.Sprintf("%s --timeout %d", flags, *opts.Timeout)
		}

		if opts.RenewAnonVolumes {
			flags = fmt.Sprintf("%s --renew-anon-volumes", flags)
		}

		if opts.RemoveOrphans {
			flags = fmt.Sprintf("%s --remove-orphans", flags)
		}

		if opts.ExitCodeFrom != "" {
			flags = fmt.Sprintf("%s --exit-code-from %s", flags, opts.ExitCodeFrom)
		}

		// Sort scaled services by service name for predictable testing
		scaleServices := make([]string, 0, len(opts.Scale))

		for service := range opts.Scale {
			scaleServices = append(scaleServices, service)
		}

		sort.Strings(scaleServices)

		for _, service := range scaleServices {
			flags = fmt.Sprintf("%s --scale %s=%d", flags, service, opts.Scale[service])
		}

		for _, service := range opts.Services {
			flags = fmt.Sprintf("%s %s", flags, service)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose up
//
// Builds, (re)creates, starts, and attaches to containers for a service.
//
// stderr is written to the given io.Writer
//
// https://docs.docker.com/compose/reference/up/
func (client *ComposeClient) Up(opts *UpOptions, w io.Writer, overrides ...*GlobalOptions) (<-chan error, error) {
	return client.RunCommand("up", upFlags(opts), nil, w, overrides...)
}
