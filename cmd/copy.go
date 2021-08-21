package cmd

import (
	"github.com/spf13/cobra"
)

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Cp files/dirs",
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.PersistentFlags().StringP("source", "s", "", "Source Path")
	copyCmd.PersistentFlags().StringP("dest", "d", "", "Dest Path")
	copyCmd.PersistentFlags().BoolP("recursive", "r", false, "recursive flag")
	copyCmd.MarkPersistentFlagRequired("source")
	copyCmd.MarkPersistentFlagRequired("dest")
}
