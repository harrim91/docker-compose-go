package client

import (
	"encoding/json"
)

type Version struct {
	Version string `json:"version"`
}

// docker compose version
//
// Show the Docker Compose version information
//
// https://docs.docker.com/engine/reference/commandline/compose_version/
func (c *ComposeClient) Version() (*Version, error) {
	res, err := c.RunQuery("version", "--format json")

	if err != nil {
		return nil, err
	}

	v := &Version{}

	err = json.Unmarshal(res, v)

	if err != nil {
		return nil, err
	}

	return v, nil
}
