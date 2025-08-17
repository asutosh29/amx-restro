package models

import (
	"fmt"

	"github.com/asutosh29/amx-restro/pkg/types"
)

func GetItems(idString string) ([]types.Item, error) {
	q := fmt.Sprintf(`
    SELECT item_id, items.category_id, item_name, item_description, img_url, price, isVeg
    FROM items
    WHERE item_id IN (%s)
`, idString)
	rows, _ := DB.Query(q)
	var ItemList []types.Item

	for rows.Next() {
		var temp types.Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item by category")
			fmt.Println(err)
			return []types.Item{}, err
		}
		ItemList = append(ItemList, temp)
	}
	return ItemList, nil
}

// ItemExistsById checks if an item exists in the database
func ItemExistsById(itemID string) (bool, error) {
	var exists int
	err := DB.QueryRow(`
    SELECT 1
    FROM items
    WHERE item_id = ?
    LIMIT 1
`, itemID).Scan(&exists)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
