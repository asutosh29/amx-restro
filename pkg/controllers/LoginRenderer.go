package controllers

import (
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderLogin(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	popup, _ := session_utils.ExtractPopupFromFlash(w, r)
	data["Popup"] = popup
	views.Tpl.ExecuteTemplate(w, "login.html", data)
}
