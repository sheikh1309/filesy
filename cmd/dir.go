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
}

func handleDir(cmd *cobra.Command, args []string) {
	var credentials config.Credentials = config.GetCredentials("my-server")
	ssh.CreateDir(credentials)
}
