package controllers

import (
	"html/template"
	"net/http"
)

func RenderLogin(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	tpl := template.Must(template.ParseFiles("pkg/static/templates/login.html", "pkg/static/templates/partials/head.html", "pkg/static/templates/partials/message.html", "pkg/static/templates/partials/bootstrap.html"))
	tpl.Execute(w, data)
}
