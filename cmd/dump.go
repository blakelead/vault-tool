package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/blakelead/vault-tool/internal/config"
	"github.com/blakelead/vault-tool/internal/vault"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:     "dump",
	Short:   "Dump secrets to stdout",
	Long:    "Print secrets to stdout in JSON format",
	RunE:    dump,
	Args:    cobra.ExactArgs(1),
	Example: "vault-tool dump secret/path",
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}

func dump(cmd *cobra.Command, args []string) error {
	client, err := vault.NewClient(args[0], config.GetSourceConfig())
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
