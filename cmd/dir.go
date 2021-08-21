package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "CRUD Dir",
	Run: handleDir,
}

func init() {
	createCmd.AddCommand(dirCmd)
	removeCmd.AddCommand(dirCmd)
	moveCmd.AddCommand(dirCmd)
	copyCmd.AddCommand(dirCmd)
	dirCmd.MarkPersistentFlagRequired("name")
	dirCmd.MarkPersistentFlagRequired("recursive")
}

func handleDir(cmd *cobra.Command, args []string) {
	var credentials = config.GetCredentials("my-server")
	dir, _ := cmd.Flags().GetString("name")
	recursive, _ := cmd.Flags().GetBool("recursive")
	if cmd.Parent().Name() == removeCmd.Name() {
		force, _ := cmd.Flags().GetBool("force")
		ssh.Remove(credentials, dir, force, recursive)
	} else if cmd.Parent().Name() == createCmd.Name() {
		ssh.CreateDir(credentials, dir)
	} else if cmd.Parent().Name() == moveCmd.Name() {
		source, _ := cmd.Flags().GetString("source")
		dest, _ := cmd.Flags().GetString("dest")
		ssh.Move(credentials, source, dest)
	} else if cmd.Parent().Name() == copyCmd.Name() {
		source, _ := cmd.Flags().GetString("source")
		dest, _ := cmd.Flags().GetString("dest")
		ssh.Copy(credentials, source, dest, recursive)
	}
	viewLs(credentials)
}

func viewLs(credentials config.Credentials)  {
	var output = ssh.List(credentials, "")
	viewLsOutput(output)
}