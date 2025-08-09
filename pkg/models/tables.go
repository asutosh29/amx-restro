package models

import (
	"fmt"

	"github.com/asutosh29/amx-restro/pkg/types"
)

// Table
func AvailableTables() ([]types.Table, error) {
	rows, _ := DB.Query(`
    SELECT table_id, isAvailable
    FROM tables
    WHERE isAvailable = 1
`)
	var tables []types.Table
	for rows.Next() {
		var temp types.Table
		rows.Scan(&temp.Table_id, &temp.IsAvailable)
		tables = append(tables, temp)
	}

	return tables, nil
}

func SetTable(table_id int, IsAvailable int) error {
	_, err := DB.Exec(`
    UPDATE tables
    SET isAvailable = ?
    WHERE table_id = ?
`, IsAvailable, table_id)
	if err != nil {
		fmt.Println("Error Setting table")
		return err
	}
	return nil
}
