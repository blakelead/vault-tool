package cmd

import (
	"github.com/blakelead/vault-tool/internal/config"
	"github.com/blakelead/vault-tool/internal/vault"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete secrets",
	Long:    "Delete secret or entire path",
	RunE:    delete,
	Args:    cobra.ExactArgs(1),
	Example: "vault-tool delete secret/path",
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) error {
	client, err := vault.NewClient(args[0], config.GetSourceConfig())
	if err != nil {
		return err
	}
	defer client.Close()

	err = client.DeleteSecrets(args[0])
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
