package controllers

import (
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderChef(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	allOrders, _ := models.GetAllOrdersByOrder()
	data["Orders"] = allOrders
	data["User"] = User

	views.Tpl.ExecuteTemplate(w, "chef.html", data)
}
