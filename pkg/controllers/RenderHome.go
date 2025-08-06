package controllers

import (
	"html/template"
	"net/http"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	tpl := template.Must(template.ParseFiles("pkg/static/templates/home.html"))
	tpl.Execute(w, data)
}
