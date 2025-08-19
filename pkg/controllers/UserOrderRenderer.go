package controllers

import (
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func RenderUserOrder(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	User := r.Context().Value("User").(types.User)
	data["User"] = User

	UserOrders, _ := models.GetOrderByUserId(User.UserId)
	data["UserOrders"] = UserOrders

	popup, err := session_utils.ExtractPopupFromFlash(w, r)
	if err != nil {
		fmt.Println("Error Loading Popus: ", err)
	}
	data["Popup"] = popup

	views.Tpl.ExecuteTemplate(w, "userOrder.html", data)
}
