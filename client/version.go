package client

import "encoding/json"

type Version struct {
	Version string `json:"version"`
}

func (c *ComposeClient) Version() (*Version, error) {
	res, err := c.RunQuery("version", "")

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