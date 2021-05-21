package vault

import (
	"fmt"
	"path"
	"strings"

	"github.com/hashicorp/vault/api"
)

func (c *Client) WriteSecrets(secrets map[string]interface{}, sourcePath, destinationPath string) error {
	if c.config.ReadOnly {
		return fmt.Errorf("cannot write. Vault is protected by the 'readonly' attribute")
	}
	for k, v := range secrets {
		mergedPath := mergePaths(sourcePath, k, destinationPath)
		_, err := c.Write(mergedPath, v.(map[string]interface{}))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) Write(path string, data map[string]interface{}) (*api.Secret, error) {
	if c.engineVersion == "1" {
		return c.vaultClient.Logical().Write(path, data)
	} else {
		secret := map[string]interface{}{"data": data}
		return c.vaultClient.Logical().Write(insert(path, "data"), secret)
	}
}

func mergePaths(sourcePath, sourceSubpath, destinationPath string) string {
	relativePath := strings.TrimPrefix(sourceSubpath, sourcePath)
	return path.Join(destinationPath, relativePath)
}
