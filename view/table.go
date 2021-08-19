package view

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func Table(headers []string, rows [][]string)  {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
}