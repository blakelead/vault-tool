package vault

import (
	"fmt"
	gopath "path"

	"github.com/hashicorp/vault/api"
	log "github.com/sirupsen/logrus"
)

func (c *Client) ReadSecrets(path string) (map[string]interface{}, error) {
	secrets := make(map[string]interface{})
	err := c.readRecurse(secrets, path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return secrets, nil
}

// readRecurse checks if a path corresponds to a secret. If it does it adds it
// to the secrets map. Otherwise it goes one step deeper into the path.
func (c *Client) readRecurse(secrets map[string]interface{}, path string) error {
	data, err := c.getSecretData(path)
	if err == nil {
		secrets[path] = data
	} else {
		paths, err := c.getSecretPaths(path)
		for _, key := range paths {
			newPath := gopath.Join(path, key.(string))
			err = c.readRecurse(secrets, newPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// getSecretData returns an error when underlying Read fails or if the path
// doesn't contain secrets. Otherwise a map of key/value pairs is returned.
func (c *Client) getSecretData(path string) (map[string]interface{}, error) {
	data, err := c.Read(path)
	if err != nil {
		log.Fatal(err)
		return map[string]interface{}{}, err
	}
	if data == nil || c.ExtractData(data.Data) == nil {
		return map[string]interface{}{}, fmt.Errorf("no secret data")
	}
	secrets := make(map[string]interface{})
	for key, val := range c.ExtractData(data.Data) {
		secrets[key] = val
	}
	return secrets, nil
}

// getSecretPaths returns a list of paths directly under the provided path.
func (c *Client) getSecretPaths(path string) ([]interface{}, error) {
	data, err := c.List(path)
	if err != nil || data == nil || data.Data == nil {
		return []interface{}{}, err
	}
	return data.Data["keys"].([]interface{}), nil
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

func (c *Client) ExtractData(data map[string]interface{}) map[string]interface{} {
	if c.engineVersion == "1" {
		return data
	} else {
		if data["data"] == nil {
			return nil
		}
		return data["data"].(map[string]interface{})
	}
}
