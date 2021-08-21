package view

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type Row struct {
	Data []string
	Colors []tablewriter.Colors
}

func Table(headers []string, rows []Row, footer []string)  {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetFooter(footer)
	for _, row := range rows {
		table.Rich(row.Data, row.Colors)
	}
	table.SetAutoMergeCellsByColumnIndex([]int{0})
	table.SetRowLine(true)
	table.Render()
}