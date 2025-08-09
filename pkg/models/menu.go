package models

import (
	"errors"
	"fmt"
	"slices"

	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/config"
)

func GetAllCategories() ([]string, error) {
	rows, err := DB.Query(`
    SELECT category_id, category_name
    FROM category
`)
	if err != nil {
		fmt.Println("Error fetching categories")
		return []string{}, err
	}
	var CategoryList []string

	for rows.Next() {
		var catg types.Category
		if err := rows.Scan(&catg.Category_id, &catg.Category_name); err != nil {
			fmt.Println("Error adding categories")
			return []string{}, err
		}
		CategoryList = append(CategoryList, catg.Category_name)
	}
	return CategoryList, nil
}

func GetAllItems() ([]types.Item, error) {
	rows, _ := DB.Query(`
    SELECT item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
    FROM items
    JOIN category ON items.category_id = category.category_id;
`)
	var ItemList []types.Item
	for rows.Next() {
		var temp types.Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item ")
			fmt.Println(err)
			return []types.Item{}, err
		}

		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}

func GetAllItemsByCategory(category_name string) ([]types.Item, error) {

	// Valid Categories
	category_list := config.ValidCategories

	IsValidCategory := slices.Contains(category_list, category_name)
	if !IsValidCategory {
		fmt.Println("Invalid category")
		return []types.Item{}, errors.New("invalid Category")
	}

	//
	rows, _ := DB.Query(`
    SELECT item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
    FROM items
    JOIN category ON items.category_id = category.category_id
    WHERE category_name = ?
`, category_name)
	var ItemList []types.Item
	for rows.Next() {
		var temp types.Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item by category")
			fmt.Println(err)
			return []types.Item{}, err
		}
		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}

func GetAllItemsBySearch(search_query string) ([]types.Item, error) {
	rows, _ := DB.Query(`
    SELECT item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
    FROM items
    JOIN category ON items.category_id = category.category_id
    WHERE item_name LIKE ?
`, "%"+search_query+"%")
	var ItemList []types.Item
	for rows.Next() {
		var temp types.Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error searching item ")
			fmt.Println(err)
			return []types.Item{}, err
		}

		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}
