package migrate

import (
	"github.com/blakelead/vault-tool/internal/vault"
	"github.com/spf13/cobra"
)

// Run migrate command
func Run(cmd *cobra.Command, args []string) error {
	sourceClient, err := vault.NewClient(args[0], vault.GetSourceConfig())
	if err != nil {
		return err
	}
	defer sourceClient.Close()

	destinationClient, err := vault.NewClient(args[1], vault.GetDestinationConfig())
	if err != nil {
		return err
	}
	defer destinationClient.Close()

	secrets, err := sourceClient.ReadSecrets(args[0])
	if err != nil {
		return err
	}

	err = destinationClient.WriteSecrets(secrets, args[0], args[1])
	if err != nil {
		return err
	}

	return nil
}
