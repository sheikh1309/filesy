package cmd

import (
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/spf13/cobra"
)

var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Create/Remove Dir",
	Run: handleDir,
}

func init() {
	createCmd.AddCommand(dirCmd)
	removeCmd.AddCommand(dirCmd)
	dirCmd.PersistentFlags().StringP("name", "n", "", "Dir name")
	dirCmd.MarkPersistentFlagRequired("name")
}

func handleDir(cmd *cobra.Command, args []string) {
	var credentials = config.GetCredentials("my-server")
	dir, _ := cmd.Flags().GetString("name")
	if cmd.Parent().Name() == removeCmd.Name() {
		force, _ := cmd.Flags().GetBool("force")
		recursive, _ := cmd.Flags().GetBool("recursive")
		ssh.Remove(credentials, dir, force, recursive)
	} else {
		ssh.CreateDir(credentials, dir)
	}
	viewLs(credentials)
}

func viewLs(credentials config.Credentials)  {
	var output = ssh.List(credentials, "")
	viewLsOutput(output)
}