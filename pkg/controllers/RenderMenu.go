package controllers

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderMenu(w http.ResponseWriter, r *http.Request) {
	var err error
	// Get Categories
	CategoryList, _ := models.GetAllCategories()
	CategoryList = append([]string{"All"}, CategoryList...)
	fmt.Println("Categories: ", CategoryList)

	query_paramas := r.URL.Query()

	category_query := query_paramas.Get("category")
	search_query := query_paramas.Get("search")

	var items []models.Item

	// Get items by category
	if category_query == "" {
		items, _ = models.GetAllItems()
	} else {
		items, err = models.GetAllItemsByCategory(category_query)
		if err != nil {
			// Invalid Category, Show All Items
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

	data := make(map[string]interface{})
	data["User"] = r.Context().Value("User")
	data["Items"] = items
	data["Categories"] = CategoryList

	views.Tpl.ExecuteTemplate(w, "menu.html", data)
}
