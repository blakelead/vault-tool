package dump

import (
	"encoding/json"
	"fmt"

	"github.com/blakelead/vault-tool/internal/vault"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	client, err := vault.NewClient(args[0], vault.GetSourceConfig())
	if err != nil {
		return err
	}
	defer client.Close()

	secrets, err := client.ReadSecrets(args[0])
	if err != nil {
		return err
	}

	jsonSecrets, err := json.MarshalIndent(secrets, "", "  ")
	if err != nil {
		return err
	}
	fmt.Print(string(jsonSecrets))

	return nil
}
