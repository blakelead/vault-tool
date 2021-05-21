package cmd

import (
	"github.com/blakelead/vault-tool/internal/config"
	"github.com/blakelead/vault-tool/internal/vault"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate secrets",
	Long:  "Migrate secrets from one Vault server to another or to the same Vault",
	RunE:  migrate,
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) error {
	sourceClient, err := vault.NewClient(args[0], config.GetSourceConfig())
	if err != nil {
		return err
	}
	defer sourceClient.Close()

	destinationClient, err := vault.NewClient(args[1], config.GetDestinationConfig())
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
