package models

import (
	"errors"
	"fmt"
	"slices"
)

type Category struct {
	Category_id   int
	Category_name string
}

type Item struct {
	Item_id          int
	Category_id      int
	Category_name    string
	Item_name        string
	Item_description string
	Img_url          string
	Price            float32
	IsVeg            bool
	IsAvailable      bool
}

func GetAllCategories() ([]string, error) {
	rows, err := DB.Query("select category_id, category_name from category")
	if err != nil {
		fmt.Println("Error fetching categories")
		return []string{}, err
	}
	var CategoryList []string

	for rows.Next() {
		var catg Category
		if err := rows.Scan(&catg.Category_id, &catg.Category_name); err != nil {
			fmt.Println("Error adding categories")
			return []string{}, err
		}
		CategoryList = append(CategoryList, catg.Category_name)
	}
	return CategoryList, nil
}

func GetAllItems() ([]Item, error) {
	rows, _ := DB.Query(`select item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
from items 
JOIN category
ON items.category_id = category.category_id;`)
	var ItemList []Item
	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item ")
			fmt.Println(err)
			return []Item{}, err
		}

		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}

func GetAllItemsByCategory(category_name string) ([]Item, error) {

	// Valid Categories
	category_list := []string{
		"Appetizers",
		"Main Course",
		"Desserts",
		"Beverages",
		"Salads",
		"Soups",
		"Snacks",
		"Combos",
	}

	IsValidCategory := slices.Contains(category_list, category_name)
	if !IsValidCategory {
		fmt.Println("Invalid category")
		return []Item{}, errors.New("invalid Category")
	}

	//
	rows, _ := DB.Query(`select item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
from items 
JOIN category
ON items.category_id = category.category_id
WHERE category_name= ? `, category_name)
	var ItemList []Item
	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item by category")
			fmt.Println(err)
			return []Item{}, err
		}
		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}

func GetAllItemsBySearch(search_query string) ([]Item, error) {
	rows, _ := DB.Query(`select item_id, items.category_id, category_name, item_name, item_description, img_url, price, isVeg
from items 
JOIN category
ON items.category_id = category.category_id
WHERE item_name like ? `, "%"+search_query+"%")
	var ItemList []Item
	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Category_name, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error searching item ")
			fmt.Println(err)
			return []Item{}, err
		}

		ItemList = append(ItemList, temp)
	}

	return ItemList, nil
}
