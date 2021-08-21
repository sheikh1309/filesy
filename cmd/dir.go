package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "A brief description of your command",
	Run: handleDir,
}

func init() {
	createCmd.AddCommand(dirCmd)
	dirCmd.PersistentFlags().StringP("dir", "d", "~", "Dir name to list files")
}

func handleDir(cmd *cobra.Command, args []string) {
	var credentials = config.GetCredentials("my-server")
	dir, _ := cmd.Flags().GetString("dir")
	ssh.CreateDir(credentials, dir)
	var output = ssh.List(credentials, "")
	viewLsOutput(output)
}
