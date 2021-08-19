package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/sheikh1309/filesy/view"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

type listResult struct {
	permission string
	owner string
	size string
	lastModified string
	name string
}

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
	var output []byte = ssh.List(credentials, dir)
	viewListOutput(output)
}

func viewListOutput(output []byte)  {
	var reader io.Reader = bytes.NewReader(output)

	var scanner = bufio.NewScanner(reader)
	// remove first line (total ...)
	scanner.Scan()

	var listResultsRows = getRows(scanner)
	var rows [][]string
	for _, row := range listResultsRows {
		rowData := []string{row.name, row.size, row.owner, row.permission, row.lastModified}
		rows = append(rows, rowData)
	}

	headers := []string{"Name", "Size", "Owner", "Permission", "Last Modified"}

	view.Table(headers, rows)
}

func getRows(scanner *bufio.Scanner) []listResult {
	var listResults []listResult
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text)
		row := listResult {
			permission: words[0],
			owner: words[2],
			size: words[4],
			// Date => Day Month at h:m ->  3 Jul at 13:20
			lastModified: fmt.Sprintf("%v %v at %v", words[6], words[5], words[7]),
			name: words[8],
		}
		listResults = append(listResults, row)
	}
	return listResults
}