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

var treeHeaders = []string{"Parent", "Name", "Size", "Owner", "Permission", "Last Modified"}

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "A brief description of your command",
	Run: handleTree,
}

func init() {
	rootCmd.AddCommand(treeCmd)
}

func handleTree(cmd *cobra.Command, args []string)  {
	profile, exists := os.LookupEnv("FILESY_PROFILE_NAME")
	if !exists {
		profile = "default"
	}
	var credentials config.Credentials = config.GetCredentials(profile)
	dir, _ := cmd.Flags().GetString("dir")
	var output []byte = ssh.Tree(credentials, dir)
	viewTreeOutput(output)
}

func viewTreeOutput(output []byte)  {
	var reader io.Reader = bytes.NewReader(output)
	var scanner = bufio.NewScanner(reader)
	var treeRows = getRows(scanner)
	var rows [][]string
	for _, row := range treeRows {
		rows = append(rows, row.Row())
	}
	footer := []string{"", "", "", "", "Total", strconv.Itoa(len(rows))}
	view.Table(treeHeaders, rows, footer)
}

func getRows(scanner *bufio.Scanner) []structs.TreeRow {
	var rows []string
	var indexes []int
	cnt := 0
	for scanner.Scan() {
		text := scanner.Text()
		rows = append(rows, text)
		if text == "" {
			indexes = append(indexes, cnt)
		}
		cnt++
	}
	treeParentMap := getTreeMap(rows, indexes)
	return getTreeRows(treeParentMap)
}

func getTreeMap(rows []string, indexes []int) map[string][]string {
	var treeParentMap map[string][]string = make(map[string][]string)
	var firstBlock []string = rows[0:indexes[0]]
	treeParentMap[firstBlock[0]] = firstBlock[1:]
	for i := 1; i < len(indexes); i++ {
		startIndex := indexes[i - 1] + 1
		endIndex := indexes[i]
		var block []string = rows[startIndex:endIndex]
		treeParentMap[block[0]] = block[1:]
	}
	return treeParentMap
}

func getTreeRows(treeParentMap map[string][]string) []structs.TreeRow {
	var treeRows []structs.TreeRow
	for key, value := range treeParentMap {
		for _, item := range value {
			columns := strings.Fields(item)
			date := fmt.Sprintf("%v %v at %v", columns[6], columns[5], columns[7])
			row := structs.TreeRow{
				Parent: key,
				LsRow: structs.LsRow {
					Permission: columns[0],
					Owner: columns[2],
					Size: columns[4],
					LastModified: date,
					Name: columns[8],
				},
			}
			treeRows = append(treeRows, row)
		}
	}
	return treeRows
}