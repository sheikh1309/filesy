package structs

type LsRow struct {
	Name string
	Size string
	Owner string
	Permission string
	LastModified string
}

func (lsRow *LsRow) Row() []string {
	return []string{lsRow.Name, lsRow.Size, lsRow.Owner, lsRow.Permission, lsRow.LastModified}
}