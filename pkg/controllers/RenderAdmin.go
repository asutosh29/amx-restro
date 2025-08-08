package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
)

func RenderAdminHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User
	fmt.Println("User :", User)

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
	categoryList := []string{"placed", "cooking", "served", "billed", "paid"}
	// Get Order details
	allOrders, _ := models.GetAllOrdersByOrder()
	// Package Data
	User := r.Context().Value("User")
	data := make(map[string]interface{})
	data["User"] = User
	data["Orders"] = allOrders
	data["Categories"] = categoryList

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
	data["User"] = User
	fmt.Println("User :", User)

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
