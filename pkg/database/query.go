package database

import "fmt"

func Search(columns []string, search string) string {
	conditions := "WHERE"
	for i := 0; i < len(columns); i++ {
		if i == (len(columns) - 1) {
			conditions += fmt.Sprintf(" %s LIKE '%s' ", columns[i], "%"+search+"%")
		} else {
			conditions += fmt.Sprintf(" %s LIKE '%s' OR ", columns[i], "%"+search+"%")
		}
	}

	return conditions
}

func OrderBy(sort_by string, sort string) string {
	return fmt.Sprintf(" ORDER BY %s %s", sort_by, sort)
}

func Limit(limit int, offset uint) string {
	return fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
}
