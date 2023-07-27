package handlers

func IsInAllowedColumns(s string, columns []string) bool {
	for _, column := range columns {
		if s == column {
			return true
		}
	}
	return false
}
