package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Create/Remove File",
	Run: handleFile,
}

func init() {
	createCmd.AddCommand(fileCmd)
	removeCmd.AddCommand(fileCmd)
	fileCmd.PersistentFlags().StringP("name", "n", "", "File name")
}

func handleFile(cmd *cobra.Command, args []string) {
	var credentials = config.GetCredentials("my-server")
	filename, _ := cmd.Flags().GetString("name")
	if cmd.Parent().Name() == removeCmd.Name() {
		ssh.Remove(credentials, filename, false, false)
	} else {
		ssh.CreateFile(credentials, filename)
	}
	viewLs(credentials)
}
