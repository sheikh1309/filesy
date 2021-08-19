package view

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func Table(headers []string, rows [][]string)  {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetFooter([]string{"", "", "", "Total", strconv.Itoa(len(rows))})
	table.AppendBulk(rows)
	table.Render()
}