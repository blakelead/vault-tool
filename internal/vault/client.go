package vault

import (
	"fmt"
	"strings"

	"github.com/blakelead/vault-tool/internal/config"
	"github.com/hashicorp/vault/api"
)

// Client is a wrapper around the Vault client
type Client struct {
	vaultClient   *api.Client
	engineVersion string
	config        *config.Config
}

// NewClient returns a new Vault Client
func NewClient(path string, config *config.Config) (*Client, error) {
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
		return nil, fmt.Errorf("invalid path")
	}
	if mount, ok := mounts[tokens[0]+"/"]; ok {
		if ver, ok := mount.Options["version"]; ok {
			version = ver
		}
	}
	if version == "" {
		return nil, fmt.Errorf("could not get engine version. This can happen when requested engine in not enabled, root path does not exist or configuration file is malformed")
	}

	return &Client{vaultClient, version, config}, nil
}

// Close closes client by clearing the token
func (c *Client) Close() {
	c.vaultClient.ClearToken()
}

// insert 'data' or 'metadata' key in secret path
func insert(path string, key string) string {
	pathParts := strings.Split(path, "/")
	if len(pathParts) > 1 {
		return fmt.Sprintf("%s/%s/%s", pathParts[0], key, strings.Join(pathParts[1:], "/"))
	} else {
		return fmt.Sprintf("%s/%s", pathParts[0], key)
	}
}
