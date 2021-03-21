package vault

import (
	"path"
	"strings"
)

func (c *Client) WriteSecrets(secrets map[string]interface{}, sourcePath, destinationPath string) error {
	for k, v := range secrets {
		mergedPath := mergePaths(sourcePath, k, destinationPath)
		c.Write(mergedPath, v.(map[string]interface{}))
	}
	return nil
}

func mergePaths(sourcePath, sourceSubpath, destinationPath string) string {
	relativePath := strings.TrimPrefix(sourceSubpath, sourcePath)
	return path.Join(destinationPath, relativePath)
}
