package views

import (
	"html/template"
	"path/filepath"
)

var Tpl *template.Template

func InitViews() {
	// Makes all the templates on server start. Doesn't rebuild in each route, serves the required one
	var initFiles []string
	layouts, _ := filepath.Glob(filepath.Join("pkg", "static", "templates", "*.html"))
	initFiles = append(initFiles, layouts...)
	adminLayouts, _ := filepath.Glob(filepath.Join("pkg", "static", "templates", "admin", "*.html"))
	initFiles = append(initFiles, adminLayouts...)
	partials, _ := filepath.Glob(filepath.Join("pkg", "static", "templates", "partials", "*.html"))
	initFiles = append(initFiles, partials...)
	components, _ := filepath.Glob(filepath.Join("pkg", "static", "templates", "components", "*.html"))
	initFiles = append(initFiles, components...)

	Tpl = template.Must(template.ParseFiles(initFiles...))

}
