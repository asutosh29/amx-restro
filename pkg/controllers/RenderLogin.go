package controllers

import (
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderLogin(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	views.Tpl.ExecuteTemplate(w, "login.html", data)
}
