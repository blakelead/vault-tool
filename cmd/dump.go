package cmd

import (
	"github.com/blakelead/vault-tool/internal/dump"
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:     "dump",
	Short:   "Dump secrets to stdout",
	Long:    "Print secrets to stdout in JSON format",
	RunE:    dump.Run,
	Args:    cobra.ExactArgs(1),
	Example: "vault-tool dump secret/path",
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
