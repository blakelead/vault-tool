package cmd

import (
	"github.com/blakelead/vault-tool/internal/migrate"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate secrets from one Vault server to another",
	Long:  "Migrate secrets from one Vault server to another",
	RunE:  migrate.Run,
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
