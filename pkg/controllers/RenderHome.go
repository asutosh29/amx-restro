package controllers

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderHome(w http.ResponseWriter, r *http.Request) {
	User := r.Context().Value("User")

	data := make(map[string]interface{})
	data["User"] = User
	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popus: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "home.html", data)
}
