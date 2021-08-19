package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "CRUD file",
	Run: handleFile,
}

func init() {
	createCmd.AddCommand(fileCmd)
}

func handleFile(cmd *cobra.Command, args []string)  {
	var credentials config.Credentials = config.GetCredentials("my-server")
	ssh.CreateFile(credentials)
}
