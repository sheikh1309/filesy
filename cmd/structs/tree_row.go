package structs

type TreeRow struct {
	Parent string
	LsRow
}

func (treeRow *TreeRow) Row() []string {
	return []string{treeRow.Parent, treeRow.Name, treeRow.Size, treeRow.Owner, treeRow.Permission, treeRow.LastModified}
}