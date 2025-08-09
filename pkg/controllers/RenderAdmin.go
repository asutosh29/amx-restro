package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"slices"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/config"
)

func RenderAdminHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User

	templFiles := []string{
		"pkg/static/templates/admin.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}

func RenderAdminOrders(w http.ResponseWriter, r *http.Request) {

	// TODO: Pagination Logic when writing frontend
	// TODO: Category
	var allOrders [][]types.OrderItem
	params := r.URL.Query()
	statusName := params.Get("category")
	statusList := config.ValidStatus

	IsValidStatus := slices.Contains(statusList, statusName)
	if !IsValidStatus {
		temp, err := models.GetAllOrdersByOrder()
		allOrders = temp
		if err != nil {
			fmt.Println("Error Fetching all orders")
			fmt.Println(err)
		}
	} else {
		temp, err := models.GetAllOrdersByOrderByStatus(statusName)
		allOrders = temp
		if err != nil {
			fmt.Println("Error Fetching all orders by status")
			fmt.Println(err)
		}
	}

	// Get Order details

	// Package Data
	User := r.Context().Value("User")
	data := make(map[string]interface{})
	data["User"] = User
	data["Orders"] = allOrders
	statusList = append([]string{"all"}, statusList...)
	data["Categories"] = statusList

	templFiles := []string{
		"pkg/static/templates/admin/orders.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
		"pkg/static/templates/partials/AdminOrderCategories.html",
		"pkg/static/templates/components/AdminOrderCard.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}

func RenderAdminUsers(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	AllUser, _ := models.GetAllUsers()
	data["User"] = User
	data["Users"] = AllUser

	templFiles := []string{
		"pkg/static/templates/admin/users.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}
