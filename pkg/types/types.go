package types

type User struct {
	Username   string
	Email      string
	First_name string
	Last_name  string
	Contact    string
	Hashpwd    string
	Userole    string
}

//

type OrderRequest struct {
	Cart         []CartItem
	Instructions string
}

type CartItem struct {
	ID   string // Have to parse it into string later
	Info ItemInfo
	Qty  int
}

type ItemInfo struct {
	ImageURL        string
	IsVeg           bool
	ItemDescription string
	ItemID          int
	ItemName        string
	Price           float32
	Qty             int
}

type Order_table_id struct {
	OrderID int
	TableID int
}

type OrderID struct{}
type TableID struct{}
