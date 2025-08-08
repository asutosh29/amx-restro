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

	// orderID := r.Context().Value("orderID")
	// tableID := r.Context().Value("tableID")
	store := session_utils.Store
	session, _ := store.Get(r, "payments")
	orderID, _ := session.Values["orderID"].(int)
	tableID, _ := session.Values["tableID"].(int)

	session.Values["orderID"] = -1
	session.Values["tableID"] = -1

	if orderID == -1 || tableID == -1 {
		fmt.Println("Bro pehle khaana order kro!")
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
	}

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
