package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove Command",
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().BoolP("force", "f", false, "force flag")
	removeCmd.PersistentFlags().BoolP("recursive", "r", false, "recursive flag")
}
