package cmd

import (
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "mv Files And Dirs",
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.PersistentFlags().StringP("source", "s", "", "Source Path")
	moveCmd.PersistentFlags().StringP("dest", "d", "", "Dest Path")
	moveCmd.PersistentFlags().BoolP("recursive", "r", false, "recursive flag")
	moveCmd.MarkPersistentFlagRequired("source")
	moveCmd.MarkPersistentFlagRequired("dest")
}
