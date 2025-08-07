package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
)

func RenderMenu(w http.ResponseWriter, r *http.Request) {
	var err error
	// Get Categories
	CategoryList, _ := models.GetAllCategories()
	CategoryList = append([]string{"All"}, CategoryList...)
	fmt.Println("Category list: ", CategoryList)

	query_paramas := r.URL.Query()

	category_query := query_paramas.Get("category")
	search_query := query_paramas.Get("search")
	fmt.Println("search query", search_query)
	fmt.Println("category query", category_query)

	var items []models.Item

	// Get items by category
	if category_query == "" {
		items, _ = models.GetAllItems()
		// fmt.Println("items: ", items)
	} else {
		items, err = models.GetAllItemsByCategory(category_query)
		if err != nil {
			// Invalid Category
			// Show All Items
			items, _ = models.GetAllItems()
		}
	}

	// Get items by Search
	if search_query != "" {
		fmt.Println("Searching your food")
		items, err = models.GetAllItemsBySearch(search_query)
		if err != nil {
			fmt.Println("Error fetching by search")
			fmt.Println(err)
		}
	}
	// fmt.Println("items:", items)

	// Package the data

	// Render the templates

	templFiles := []string{
		"pkg/static/templates/menu.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
		"pkg/static/templates/partials/categories.html",
		"pkg/static/templates/components/MenuCard.html",
	}

	data := make(map[string]interface{})
	// TODO: Retrive this from the user object. Dummy for now
	data["User"] = types.User{
		Username:   "username",
		Email:      "email",
		First_name: "first_name",
		Last_name:  "last_name",
		Contact:    "contact",
		Hashpwd:    "hashpwd",
		Userole:    "customer",
	}
	data["Items"] = items
	data["Categories"] = CategoryList
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}
