package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: handleList,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("dir", "d", "~", "Dir name to list files")
}

func handleList(cmd *cobra.Command, args []string)  {
	profile, exists := os.LookupEnv("FILESY_PROFILE_NAME")
	if !exists {
		profile = "default"
	}
	var credentials config.Credentials = config.GetCredentials(profile)
	dir, _ := cmd.Flags().GetString("dir")
	ssh.List(credentials, dir)
}