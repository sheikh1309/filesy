package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sheikh1309/filesy/cmd/structs"
	"github.com/sheikh1309/filesy/config"
	"github.com/sheikh1309/filesy/ssh"
	"github.com/sheikh1309/filesy/view"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strconv"
	"strings"
)

var lsHeaders = []string{"Name", "Size", "Owner", "Permission", "Last Modified"}

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
	viewLsOutput(output)
}

func viewLsOutput(output []byte)  {
	var reader io.Reader = bytes.NewReader(output)
	var scanner = bufio.NewScanner(reader)
	var listResultsRows = getLsRows(scanner)
	var rows [][]string
	for _, row := range listResultsRows {
		rows = append(rows, row.Row())
	}
	footer := []string{"", "", "", "Total", strconv.Itoa(len(rows))}
	view.Table(lsHeaders, rows, footer)
}

func getLsRows(scanner *bufio.Scanner) []structs.LsRow {
	var lsRows []structs.LsRow
	for scanner.Scan() {
		text := scanner.Text()
		columns := strings.Fields(text)
		date := fmt.Sprintf("%v %v at %v", columns[6], columns[5], columns[7])
		row := structs.LsRow {
			Permission: columns[0],
			Owner: columns[2],
			Size: columns[4],
			LastModified: date, // Date => Day Month at h:m ->  3 Jul at 13:20
			Name: columns[8],
		}
		lsRows = append(lsRows, row)
	}
	return lsRows
}