package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asutosh29/amx-restro/pkg/types"
)

type Cart struct {
	Id   int
	Qty  int
	Info Item
}

type Table struct {
	Table_id    int
	IsAvailable bool
}

func AddOrder(instruction string, cart []types.CartItem, user types.User) (int, int) {
	fmt.Println("user:", user)
	fmt.Println("cart:", cart)
	fmt.Println("instruction:", instruction)
	// Get User ID
	realUser, _ := GetUser(user)
	userID, _ := GertUserId(user)
	fmt.Println("Real User:", realUser)

	// Get Table ID
	tables, _ := AvailableTables()
	fmt.Println("Tables: ", tables)

	var table_id int
	if len(tables) != 0 {
		table_id = tables[0].Table_id
	}

	err := SetTable(table_id, 0)
	if err != nil {
		fmt.Println("Error setting table")
		fmt.Println(err)
	}
	// Format Time in Correct Format
	order_at_time := time.Now().Format("2006-01-02 15:04:05")

	// Extract item IDs from the input to query the database.
	var itemIDs []string
	for _, item := range cart {
		itemIDs = append(itemIDs, fmt.Sprintf("%v", item.ID))
	}
	idString := strings.Join(itemIDs, ",")
	fmt.Println("idString:", idString)

	// Fetch item prices from the database.
	items, err := GetItems(idString)
	if err != nil {
		fmt.Println("Error fetching items from DB")
		fmt.Println(err)
	}
	fmt.Println("Items ordered:", items)

	// Calculate total amount
	var totalAmount float32
	for _, e := range cart {
		for _, item := range items {
			fmt.Println("--- DEBUG ----")
			fmt.Println("cart e:", e)
			fmt.Println("items item:", item)
			ordered_qty := float32(e.Qty)
			var price float32
			ID, _ := strconv.Atoi(e.ID)
			if item.Item_id == ID {
				price = float32(item.Price)
				fmt.Println("Adding: ", ID, "Price:", price, "QTY: ", ordered_qty)
			} else {
				price = 0
			}
			totalAmount += float32(ordered_qty) * float32(price)
		}
	}

	fmt.Println("Total amount: ", totalAmount)
	fmt.Println("Time : ", order_at_time)
	// Insert into DB
	fmt.Println("User ID:", userID)
	result, err := DB.Exec(`insert into orders(customer_id, table_id, extra_instructions, total_amount, order_at_time)
                    values(?, ?, ?, ?, ?)`, userID, table_id, instruction, totalAmount, order_at_time)
	if err != nil {
		fmt.Println("Error placing order")
		fmt.Println(err)
		return -1, -1
	}

	orderID, _ := result.LastInsertId()

	for _, it_qt := range cart {
		DB.Exec(`insert into order_item(item_id, order_id, qty) value(?, ?, ?)`, it_qt.ID, orderID, it_qt.Qty)
	}
	fmt.Println("order Placed successfully: ", orderID, table_id)
	// Return {orderID, tableID}
	return int(orderID), table_id
}

func AvailableTables() ([]Table, error) {
	rows, _ := DB.Query(`select table_id, isAvailable
                from tables 
                where isAvailable=1`)
	var tables []Table
	for rows.Next() {
		var temp Table
		rows.Scan(&temp.Table_id, &temp.IsAvailable)
		tables = append(tables, temp)
	}

	return tables, nil
}

func SetTable(table_id int, IsAvailable int) error {
	_, err := DB.Exec(`update tables set isAvailable= ? where table_id=?`, IsAvailable, table_id)
	if err != nil {
		fmt.Println("Error Setting table")
		return err
	}
	return nil
}

func GetItems(idString string) ([]Item, error) {
	q := fmt.Sprintf(`select item_id, items.category_id, item_name, item_description, img_url, price, isVeg 
                    from items
                    where item_id in ( %s )`, idString)
	fmt.Println("Query: ", q)
	rows, _ := DB.Query(q)
	fmt.Println("Item string: ", idString)
	var ItemList []Item

	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item by category")
			fmt.Println(err)
			return []Item{}, err
		}
		fmt.Println("Adding item GET TIEMS:", temp)
		ItemList = append(ItemList, temp)
	}
	return ItemList, nil
}
