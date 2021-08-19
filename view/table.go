package view

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

func Table(headers []string, rows [][]string, footer []string)  {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetFooter(footer)
	table.AppendBulk(rows)
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.Render()
}