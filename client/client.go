package client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/harrim91/docker-compose-go/cmd"
)

// New returns a new default ComposeClient
func New(config *Config) *ComposeClient {
	return &ComposeClient{
		Config: config,
		NewCmd: func() Cmd {
			return cmd.New()
		},
	}
}

// ComposeClient is used for executing Docker Compose commands. It should be created with `New`.
type ComposeClient struct {
	Config *Config
	NewCmd func() Cmd
}

type Cmd interface {
	SetStderr(stderr io.Writer)
	SetStdout(stdout io.Writer)
	Run(command string) (<-chan error, error)
}

// Config represents the global configuration options for the ComposeClient
//
// https://docs.docker.com/compose/reference/
type Config struct {
	// Specify alternate compose file(s) (default: docker-compose.yml)
	Files []string

	// Specify an alternate project name (default: directory name)
	ProjectName string

	// Specify a profile to enable
	Profiles []string

	// Show more output
	Verbose *bool

	// Do not print ANSI control characters
	NoANSI *bool

	// Daemon socket to connect to
	Host string

	// Use TLS; implied by TLSVerify
	TLS *bool

	// Trust certs signed only by this CA
	TLSCACert string

	// Path to TLS certificate file
	TLSCert string

	// Path to TLS key file
	TLSKey string

	// Use TLS and verify the remote
	TLSVerify *bool

	// Don't check the daemon's hostname against the name specified in the client certificate
	SkipHostnameCheck bool

	// Specify an alternate working directory (default: the path of the Compose file)
	ProjectDirectory string

	// If set, Compose will attempt to convert deploy keys in v3 files to their non-Swarm equivalent
	Compatibility *bool
}

func (c *ComposeClient) globalFlags(overrides ...*Config) string {
	flags := ""

	if c.Config != nil {
		for _, file := range c.Config.Files {
			flags = fmt.Sprintf("%s --file %s", flags, file)
		}
	}

	for _, override := range overrides {
		for _, file := range override.Files {
			flags = fmt.Sprintf("%s --file %s", flags, file)
		}
	}

	if c.Config != nil {
		for _, profile := range c.Config.Profiles {
			flags = fmt.Sprintf("%s --profile %s", flags, profile)
		}
	}

	for _, override := range overrides {
		for _, profile := range override.Profiles {
			flags = fmt.Sprintf("%s --profile %s", flags, profile)
		}
	}

	var (
		projectName      string
		verbose          *bool
		noANSI           *bool
		host             string
		tls              *bool
		tlsCACert        string
		tlsCert          string
		tlsKey           string
		tlsVerify        *bool
		projectDirectory string
		compatibility    *bool
	)

	if c.Config != nil {
		projectName = c.Config.ProjectName
		verbose = c.Config.Verbose
		noANSI = c.Config.NoANSI
		host = c.Config.Host
		tls = c.Config.TLS
		tlsCACert = c.Config.TLSCACert
		tlsCert = c.Config.TLSCert
		tlsKey = c.Config.TLSKey
		tlsVerify = c.Config.TLSVerify
		projectDirectory = c.Config.ProjectDirectory
		compatibility = c.Config.Compatibility
	}

	for _, override := range overrides {
		if override.ProjectName != "" {
			projectName = override.ProjectName
		}

		if override.Verbose != nil {
			verbose = override.Verbose
		}

		if override.NoANSI != nil {
			noANSI = override.NoANSI
		}

		if override.Host != "" {
			host = override.Host
		}

		if override.TLS != nil {
			tls = override.TLS
		}

		if override.TLSCACert != "" {
			tlsCACert = override.TLSCACert
		}

		if override.TLSCert != "" {
			tlsCert = override.TLSCert
		}

		if override.TLSKey != "" {
			tlsKey = override.TLSKey
		}

		if override.TLSVerify != nil {
			tlsVerify = override.TLSVerify
		}

		if override.ProjectDirectory != "" {
			projectDirectory = override.ProjectDirectory
		}

		if override.Compatibility != nil {
			compatibility = override.Compatibility
		}
	}

	if projectName != "" {
		flags = fmt.Sprintf("%s --project-name %s", flags, projectName)
	}

	if verbose != nil && *verbose {
		flags = fmt.Sprintf("%s --verbose", flags)
	}

	if noANSI != nil && *noANSI {
		flags = fmt.Sprintf("%s --no-ansi", flags)
	}

	if host != "" {
		flags = fmt.Sprintf("%s --host %s", flags, host)
	}

	if tls != nil && *tls {
		flags = fmt.Sprintf("%s --tls", flags)
	}

	if tlsCACert != "" {
		flags = fmt.Sprintf("%s --tlscacert %s", flags, tlsCACert)
	}

	if tlsCert != "" {
		flags = fmt.Sprintf("%s --tlscert %s", flags, tlsCert)
	}

	if tlsKey != "" {
		flags = fmt.Sprintf("%s --tlskey %s", flags, tlsKey)
	}

	if tlsVerify != nil && *tlsVerify {
		flags = fmt.Sprintf("%s --tlsverify", flags)
	}

	if projectDirectory != "" {
		flags = fmt.Sprintf("%s --project-directory %s", flags, projectDirectory)
	}

	if compatibility != nil && *compatibility {
		flags = fmt.Sprintf("%s --compatibility", flags)
	}

	return flags
}

// RunCommand executes the given docker compose command.
//
// stdout and stderr from the underlying docker compose processes are written to the given io.Writers
//
// Users would normally use of one of the specific command methods (e.g. Up, Down)
func (client *ComposeClient) RunCommand(command, flags string, stdout, stderr io.Writer, overrides ...*Config) (<-chan error, error) {
	cmd := client.NewCmd()

	if stdout != nil {
		cmd.SetStdout(stdout)
	}

	if stderr != nil {
		cmd.SetStderr(stderr)
	}

	return cmd.Run(strings.TrimSpace(fmt.Sprintf("docker compose%s %s %s", client.globalFlags(overrides...), command, flags)))
}

// RunQuery executes the given docker compose query, and returns the returned JSON byte array.
//
// Users would normally use of one of the specific query methods (e.g. Version)
func (client *ComposeClient) RunQuery(command, flags string, overrides ...*Config) ([]byte, error) {
	var stdout bytes.Buffer
	var result []byte

	ch, err := client.RunCommand(command, strings.TrimSpace(fmt.Sprintf("%s --format json", flags)), &stdout, nil, overrides...)

	if err != nil {
		return result, err
	}

	err = <-ch

	if err != nil {
		return result, err
	}

	scanner := bufio.NewScanner(&stdout)

	for scanner.Scan() {
		result = append(result, scanner.Bytes()...)
	}

	return result, nil
}
