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
	// Get User ID
	userID, _ := GertUserId(user)

	// Get Table ID
	tables, _ := AvailableTables()

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

	// Fetch item prices from the database.
	items, err := GetItems(idString)
	if err != nil {
		fmt.Println("Error fetching items from DB")
		fmt.Println(err)
	}

	// Calculate total amount
	var totalAmount float32
	for _, e := range cart {
		for _, item := range items {
			ordered_qty := float32(e.Qty)
			var price float32
			ID, _ := strconv.Atoi(e.ID)
			if item.Item_id == ID {
				price = float32(item.Price)
			} else {
				price = 0
			}
			totalAmount += float32(ordered_qty) * float32(price)
		}
	}

	// Insert into DB
	result, err := DB.Exec(`
    INSERT INTO orders (customer_id, table_id, extra_instructions, total_amount, order_at_time)
    VALUES (?, ?, ?, ?, ?)
`, userID, table_id, instruction, totalAmount, order_at_time)
	if err != nil {
		fmt.Println("Error placing order")
		fmt.Println(err)
		return -1, -1
	}

	orderID, _ := result.LastInsertId()

	for _, it_qt := range cart {
		DB.Exec(`
    INSERT INTO order_item (item_id, order_id, qty)
    VALUES (?, ?, ?)
`, it_qt.ID, orderID, it_qt.Qty)
	}
	return int(orderID), table_id
}

func GetAllOrdersByOrder() ([][]types.OrderItem, error) {
	var IdList []int
	var OrderList [][]types.OrderItem
	// var OrderList []any
	// Create Unique ID List
	ids, _ := DB.Query(`
    SELECT DISTINCT order_id
    FROM orders
    ORDER BY order_id ASC
`)
	for ids.Next() {
		var temp int
		ids.Scan(&temp)
		IdList = append(IdList, temp)
	}

	// Fetch Orders by ID
	for _, orderId := range IdList {
		var tempOrderItem []types.OrderItem
		tempOrderItem, _ = GetOrder(orderId)
		OrderList = append(OrderList, tempOrderItem)
	}
	// return the Slice of orders
	return OrderList, nil
}

func GetAllOrdersByOrderByStatus(statusName string) ([][]types.OrderItem, error) {
	var IdList []int
	var OrderList [][]types.OrderItem
	// var OrderList []any
	// Create Unique ID List
	ids, _ := DB.Query(`
    SELECT DISTINCT order_id
    FROM orders
    WHERE order_status = ?
    ORDER BY order_id ASC;
`, statusName)
	for ids.Next() {
		var temp int
		ids.Scan(&temp)
		IdList = append(IdList, temp)
	}

	// Fetch Orders by ID
	for _, orderId := range IdList {
		var tempOrderItem []types.OrderItem
		tempOrderItem, _ = GetOrder(orderId)
		OrderList = append(OrderList, tempOrderItem)
	}
	// return the Slice of orders
	return OrderList, nil
}

func GetOrder(orderId int) ([]types.OrderItem, error) {
	rows, err := DB.Query(`
    SELECT orders.order_id, customer_id, tables.table_id, extra_instructions, order_status, total_amount, order_at_time,
           items.item_id, qty, category_id, item_name, item_description, img_url, price, isVeg, items.isAvailable
    FROM orders
    JOIN tables ON tables.table_id = orders.table_id
    JOIN users ON orders.customer_id = users.id
    JOIN order_item ON orders.order_id = order_item.order_id
    JOIN items ON items.item_id = order_item.item_id
    WHERE orders.order_id = ?
    ORDER BY orders.order_id DESC;
`, orderId)
	var OrderList []types.OrderItem
	if err != nil {
		fmt.Println("Error Fetching Order")
		fmt.Println(err)
		return OrderList, err
	}
	for rows.Next() {
		var tempOrderItem types.OrderItem
		err := rows.Scan(&tempOrderItem.OrderID, &tempOrderItem.CustomerID, &tempOrderItem.TableID, &tempOrderItem.Extra_instructions, &tempOrderItem.Order_status, &tempOrderItem.Total_amount, &tempOrderItem.Order_at_time, &tempOrderItem.ItemID, &tempOrderItem.Qty, &tempOrderItem.CategoryID, &tempOrderItem.ItemName, &tempOrderItem.ItemDescription, &tempOrderItem.ImageURL, &tempOrderItem.Price, &tempOrderItem.IsVeg, &tempOrderItem.IsAvailable)
		if err != nil {
			fmt.Println("Error Fetching Order")
			fmt.Println(err)
			return OrderList, err
		}
		OrderList = append(OrderList, tempOrderItem)
	}

	return OrderList, nil
}

// Order Status Change
func MarkOrderPlacedById(orderID int) error {
	_, err := DB.Exec(`
    UPDATE orders
    SET order_status = 'placed'
    WHERE order_id = ?;
`, orderID)
	if err != nil {
		fmt.Println("Couldn't change state of Order")
		return err
	}
	var tableID int
	DB.QueryRow(`
    SELECT table_id
    FROM orders
    WHERE order_id = ?
`, orderID).Scan(&tableID)
	// Update Table Status
	SetTable(tableID, 0)
	return nil
}
func MarkOrderCookingById(orderID int) error {
	_, err := DB.Exec(`
    UPDATE orders
    SET order_status = 'cooking'
    WHERE order_id = ?;
`, orderID)
	if err != nil {
		fmt.Println("Couldn't change state of Order")
		return err
	}
	var tableID int
	DB.QueryRow(`
    SELECT table_id
    FROM orders
    WHERE order_id = ?
`, orderID).Scan(&tableID)
	// Update Table Status
	SetTable(tableID, 0)
	return nil
}
func MarkOrderServedById(orderID int) error {
	_, err := DB.Exec(`
    UPDATE orders
    SET order_status = 'served'
    WHERE order_id = ?;
`, orderID)
	if err != nil {
		fmt.Println("Couldn't change state of Order")
		return err
	}
	var tableID int
	DB.QueryRow(`
    SELECT table_id
    FROM orders
    WHERE order_id = ?
`, orderID).Scan(&tableID)
	// Update Table Status
	SetTable(tableID, 1)
	return nil
}
func MarkOrderBilledById(orderID int) error {
	_, err := DB.Exec(`
    UPDATE orders
    SET order_status = 'billed'
    WHERE order_id = ?;
`, orderID)
	if err != nil {
		fmt.Println("Couldn't change state of Order")
		return err
	}
	var tableID int
	DB.QueryRow(`
    SELECT table_id
    FROM orders
    WHERE order_id = ?
`, orderID).Scan(&tableID)
	// Update Table Status
	SetTable(tableID, 0)
	return nil
}
func MarkOrderPaidById(orderID int) error {
	_, err := DB.Exec(`
    UPDATE orders
    SET order_status = 'paid'
    WHERE order_id = ?;
`, orderID)
	if err != nil {
		fmt.Println("Couldn't change state of Order")
		return err
	}
	var tableID int
	DB.QueryRow(`
    SELECT table_id
    FROM orders
    WHERE order_id = ?
`, orderID).Scan(&tableID)
	// Update Table Status
	SetTable(tableID, 1)
	return nil
}

// Table
func AvailableTables() ([]Table, error) {
	rows, _ := DB.Query(`
    SELECT table_id, isAvailable
    FROM tables
    WHERE isAvailable = 1
`)
	var tables []Table
	for rows.Next() {
		var temp Table
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

// Items
func GetItems(idString string) ([]Item, error) {
	q := fmt.Sprintf(`
    SELECT item_id, items.category_id, item_name, item_description, img_url, price, isVeg
    FROM items
    WHERE item_id IN (%s)
`, idString)
	rows, _ := DB.Query(q)
	var ItemList []Item

	for rows.Next() {
		var temp Item
		err := rows.Scan(&temp.Item_id, &temp.Category_id, &temp.Item_name, &temp.Item_description, &temp.Img_url, &temp.Price, &temp.IsVeg)
		if err != nil {
			fmt.Println("Error adding item by category")
			fmt.Println(err)
			return []Item{}, err
		}
		ItemList = append(ItemList, temp)
	}
	return ItemList, nil
}
