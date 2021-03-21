package vault

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
)

// Client is a wrapper around the Vault client
type Client struct {
	vaultClient   *api.Client
	engineVersion string
}

// NewClient returns a new Vault Client
func NewClient(path string, config *Config) (*Client, error) {
	vaultConfig := api.DefaultConfig()
	vaultConfig.Address = config.Address
	vaultConfig.ConfigureTLS(&api.TLSConfig{
		Insecure: config.TLSInsecure,
	})

	vaultClient, err := api.NewClient(vaultConfig)
	if err != nil {
		return nil, err
	}

	vaultClient.SetToken(config.Token)

	// list of mounts will be used to get the engine version
	// of the provided path
	mounts, err := vaultClient.Sys().ListMounts()
	if err != nil {
		return nil, err
	}

	// extract the secret engine name and get its version
	var version string
	tokens := strings.Split(path, "/")
	if len(tokens) == 0 {
		return nil, fmt.Errorf("Invalid path")
	}
	if mount, ok := mounts[tokens[0]+"/"]; ok {
		if ver, ok := mount.Options["version"]; ok {
			version = ver
		}
	}
	if version == "" {
		return nil, fmt.Errorf("Could not get engine version")
	}

	return &Client{vaultClient, version}, nil
}

// Close closes client by clearing the token
func (c *Client) Close() {
	c.vaultClient.ClearToken()
}

func (c *Client) Read(path string) (*api.Secret, error) {
	if c.engineVersion == "1" {
		return c.vaultClient.Logical().Read(path)
	} else {
		return c.vaultClient.Logical().Read(insert(path, "data"))
	}
}

func (c *Client) List(path string) (*api.Secret, error) {
	if c.engineVersion == "1" {
		return c.vaultClient.Logical().List(path)
	} else {
		return c.vaultClient.Logical().List(insert(path, "metadata"))
	}
}

func (c *Client) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	if c.engineVersion == "1" {
		return c.vaultClient.Logical().Write(path, data)
	} else {
		secret := map[string]interface{}{"data": data}
		return c.vaultClient.Logical().Write(insert(path, "data"), secret)
	}
}

func (c *Client) ExtractData(data map[string]interface{}) map[string]interface{} {
	if c.engineVersion == "1" {
		return data
	} else {
		return data["data"].(map[string]interface{})
	}
}

func insert(path string, token string) string {
	tokens := strings.Split(path, "/")
	if len(tokens) > 1 {
		return fmt.Sprintf("%s/%s/%s", tokens[0], token, strings.Join(tokens[1:], "/"))
	} else {
		return fmt.Sprintf("%s/%s", tokens[0], token)
	}
}
