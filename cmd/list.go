package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/olekukonko/tablewriter"
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
	var credentials = config.GetCredentials(profile)
	dir, _ := cmd.Flags().GetString("dir")
	var output = ssh.List(credentials, dir)
	viewLsOutput(output)
}

func viewLsOutput(output []byte)  {
	var reader io.Reader = bytes.NewReader(output)
	var scanner = bufio.NewScanner(reader)
	var listResultsRows = getLsRows(scanner)
	footer := []string{"", "", "", "Total", strconv.Itoa(len(listResultsRows))}
	view.Table(lsHeaders, listResultsRows, footer)
}

func getLsRows(scanner *bufio.Scanner) []view.Row {
	var rows []view.Row
	for scanner.Scan() {
		text := scanner.Text()
		columns := strings.Fields(text)
		date := fmt.Sprintf("%v %v at %v", columns[6], columns[5], columns[7])
		data := structs.LsRow { Permission: columns[0], Owner: columns[2], Size: columns[4], LastModified: date, Name: columns[8] }
		row := view.Row { Data: data.Row(), Colors: getListRowColor(data.Name), }
		rows = append(rows, row)
	}
	return rows
}

func getListRowColor(name string) []tablewriter.Colors {
	var colors = []tablewriter.Colors{tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}}
	if strings.HasSuffix(name, "/") {
		colors[0] = tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor}
	}
	return colors
}