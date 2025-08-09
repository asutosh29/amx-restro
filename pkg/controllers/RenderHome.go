package controllers

import (
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	User := r.Context().Value("User")

	data := make(map[string]interface{})
	data["User"] = User

	views.Tpl.ExecuteTemplate(w, "home.html", data)
}
