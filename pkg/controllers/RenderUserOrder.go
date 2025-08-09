package controllers

import (
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderUserOrder(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User

	views.Tpl.ExecuteTemplate(w, "userOrder.html", data)
}
