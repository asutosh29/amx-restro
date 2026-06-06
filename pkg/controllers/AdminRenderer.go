package controllers

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/config"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderAdminHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User

	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popus: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "admin.html", data)
}

func RenderAdminOrders(w http.ResponseWriter, r *http.Request) {

	// TODO: Pagination Logic when writing frontend
	params := r.URL.Query()
	statusName := params.Get("category")
	statusList := config.ValidStatus
	orderCounts := make(map[string]int, len(statusList)+1)

	allOrderGroups, err := models.GetAllOrdersByOrder()
	if err != nil {
		fmt.Println("Error Fetching all orders")
		fmt.Println(err)
	}

	orderCounts["all"] = len(allOrderGroups)
	for _, status := range statusList {
		orderCounts[status] = 0
	}
	for _, orderGroup := range allOrderGroups {
		if len(orderGroup) == 0 {
			continue
		}
		orderCounts[orderGroup[0].Order_status]++
	}

	// Sort: paid orders always at the bottom
	slices.SortStableFunc(allOrderGroups, func(a, b []types.OrderItem) int {
		if len(a) == 0 || len(b) == 0 {
			return 0
		}
		if a[0].Order_status == "paid" && b[0].Order_status != "paid" {
			return 1
		}
		if a[0].Order_status != "paid" && b[0].Order_status == "paid" {
			return -1
		}
		return 0
	})

	allOrders := allOrderGroups

	// For Invalid status show all orders
	IsValidStatus := slices.Contains(statusList, statusName)
	if IsValidStatus {
		temp, err := models.GetAllOrdersByOrderByStatus(statusName)
		allOrders = temp
		if err != nil {
			fmt.Println("Error Fetching all orders by status")
			fmt.Println(err)
		}
	}

	// Package Data
	User := r.Context().Value("User")
	data := make(map[string]interface{})
	if statusName != "" {
		data["Category"] = statusName
	} else {
		data["Category"] = "all"
	}
	data["User"] = User
	data["Orders"] = allOrders
	data["OrderCounts"] = orderCounts

	statusList = append([]string{"all"}, statusList...)
	data["Categories"] = statusList

	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popus: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "orders.html", data)
}

func RenderAdminUsers(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	AllUser, _ := models.GetAllUsers()
	data["User"] = User
	data["Users"] = AllUser

	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popus: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "adminUsers.html", data)
}
