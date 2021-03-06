package vault

import (
	"fmt"
	gopath "path"

	"github.com/hashicorp/vault/api"
)

func (c *Client) DeleteSecrets(path string) error {
	if c.config.ReadOnly {
		return fmt.Errorf("cannot delete. Vault is protected by the 'readonly' attribute")
	}
	return c.deleteRecurse(path)
}

func (c *Client) deleteRecurse(path string) error {
	_, err := c.getSecretData(path)
	if err == nil {
		_, err = c.Delete(path)
	} else {
		paths, err := c.getSecretPaths(path)
		for _, key := range paths {
			newPath := gopath.Join(path, key.(string))
			err = c.deleteRecurse(newPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Client) Delete(path string) (*api.Secret, error) {
	if c.engineVersion == "1" {
		return c.vaultClient.Logical().Delete(path)
	} else {
		return c.vaultClient.Logical().Delete(insert(path, "metadata"))
	}
}
