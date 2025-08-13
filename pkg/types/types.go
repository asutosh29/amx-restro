package types

type User struct {
	UserId    int
	Username  string
	Email     string
	FirstName string
	LastName  string
	Contact   string
	Hashpwd   string
	Userole   string
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
	CategoryID      int
	ItemName        string
	Price           float32
	Qty             int
	IsAvailable     int
}

type OrderItem struct {
	OrderID            int
	Order_status       string
	CustomerID         int
	TableID            int
	Extra_instructions string
	Total_amount       float32
	Order_at_time      string
	ItemInfo
}

//

type Order_table_id struct {
	OrderID int
	TableID int
}

type OrderID struct{}
type TableID struct{}

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

type Cart struct {
	Id   int
	Qty  int
	Info Item
}

type Table struct {
	Table_id    int
	IsAvailable bool
}

type Popup struct {
	Msg     string
	IsError bool
}
