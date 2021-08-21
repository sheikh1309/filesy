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
	fileCmd.PersistentFlags().StringP("filename", "f", "~", "Dir name to list files")
}

func handleFile(cmd *cobra.Command, args []string)  {
	var credentials = config.GetCredentials("my-server")
	filename, _ := cmd.Flags().GetString("filename")
	ssh.CreateFile(credentials, filename)
	var output = ssh.List(credentials, "")
	viewLsOutput(output)
}
