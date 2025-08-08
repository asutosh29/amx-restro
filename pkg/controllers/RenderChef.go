package controllers

import (
	"html/template"
	"net/http"
)

func RenderChef(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User

	templFiles := []string{
		"pkg/static/templates/chef.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}
