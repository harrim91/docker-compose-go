package client

import (
	"fmt"
	"strings"
)

type ConfigOptions struct {
	// Pin image tags to digests.
	ResolveImageDigests bool

	// Don't interpolate environment variables.
	NoInterpolate bool

	// Only validate the configuration, don't print anything.
	Quiet bool

	// Print the service names, one per line. Takes precedence over `Volumes` flag.
	// This will make the function return a []byte representing a JSON string array of service names.
	Services bool

	// Print the volume names, one per line. Takes precedence over `Hash` flag.
	// This will make the function return a []byte representing a JSON string array of volume names.
	Volumes bool

	// Print the service config hash, one per line.
	// Set "service1,service2" for a list of specified services or use the wildcard symbol "*" to display all services.
	Hash string
}

func configFlags(opts *ConfigOptions) string {
	flags := "--format json"

	if opts != nil {
		if opts.ResolveImageDigests {
			flags = fmt.Sprintf("%s --resolve-image-digests", flags)
		}

		if opts.NoInterpolate {
			flags = fmt.Sprintf("%s --no-interpolate", flags)
		}

		if opts.Quiet {
			flags = fmt.Sprintf("%s --quiet", flags)
		}

		if opts.Services {
			flags = fmt.Sprintf("%s --services", flags)
		}

		if opts.Volumes {
			flags = fmt.Sprintf("%s --volumes", flags)
		}

		if opts.Hash != "" {
			flags = fmt.Sprintf("%s --hash=\"%s\"", flags, opts.Hash)
		}
	}

	return strings.TrimSpace(flags)
}

// docker compose config
//
// Validate and view the Compose file.

// Returns a byte array representing the Compose file in JSON format.
//
// If Services, Volumes or Hash options are specified, returns a byte array representing the list of services/volumes/hashes (one per line)
//
// https://docs.docker.com/compose/reference/config/
func (c *ComposeClient) Config(opts *ConfigOptions) ([]byte, error) {
	return c.RunQuery("config", configFlags(opts))
}
