package controllers

import (
	"html/template"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/types"
)

func RenderMenu(w http.ResponseWriter, r *http.Request) {

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

	templFiles := []string{
		"pkg/static/templates/menu.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
		"pkg/static/templates/partials/categories.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}
