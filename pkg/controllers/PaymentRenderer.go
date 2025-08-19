package controllers

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderPayment(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User")
	data["User"] = User

	store := session_utils.Store
	session, err := store.Get(r, "payments")

	// TODO: Fix Session Issue
	if err != nil {
		fmt.Println("Error Loading session")
		fmt.Println(err)
	}
	fmt.Println("session", session)
	orderID := session.Values["orderID"]
	tableID := session.Values["tableID"]

	session.Values["orderID"] = -1
	session.Values["tableID"] = -1
	session.Save(r, w)

	if orderID == -1 || tableID == -1 {
		session_utils.FlashMsgErr(w, r, "Please Place an Order first!", true)
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
	}

	data["OrderID"] = orderID
	data["TableID"] = tableID

	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popup: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "payment.html", data)
}
