package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
)

func RenderPayment(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User
	fmt.Println("User :", User)

	// orderID := r.Context().Value("orderID")
	// tableID := r.Context().Value("tableID")
	store := session_utils.Store
	session, _ := store.Get(r, "payments")
	orderID, _ := session.Values["orderID"].(int)
	tableID, _ := session.Values["tableID"].(int)

	fmt.Println(orderID, tableID)
	data["OrderID"] = orderID
	data["TableID"] = tableID

	templFiles := []string{
		"pkg/static/templates/payment.html",
		"pkg/static/templates/partials/head.html",
		"pkg/static/templates/partials/message.html",
		"pkg/static/templates/partials/bootstrap.html",
		"pkg/static/templates/partials/navbar.html",
	}
	tpl := template.Must(template.ParseFiles(templFiles...))
	tpl.Execute(w, data)
}
