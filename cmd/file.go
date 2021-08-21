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
	moveCmd.AddCommand(fileCmd)
}

func handleFile(cmd *cobra.Command, args []string) {
	var credentials = config.GetCredentials("my-server")
	filename, _ := cmd.Flags().GetString("name")
	if cmd.Parent().Name() == removeCmd.Name() {
		ssh.Remove(credentials, filename, false, false)
	} else if cmd.Parent().Name() == createCmd.Name() {
		ssh.CreateFile(credentials, filename)
	} else if cmd.Parent().Name() == moveCmd.Name() {
		source, _ := cmd.Flags().GetString("source")
		dest, _ := cmd.Flags().GetString("dest")
		ssh.Move(credentials, source, dest)
	}
	viewLs(credentials)
}
