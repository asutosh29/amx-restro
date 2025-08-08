package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
)

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("See All order!")
}

func HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	var jsonResponse types.OrderRequest
	_ = json.NewDecoder(r.Body).Decode(&jsonResponse)
	fmt.Println("CART: ", jsonResponse)

	User := r.Context().Value("User").(types.User)
	fmt.Println(User)
	orderId, tableID := models.AddOrder(jsonResponse.Instructions, jsonResponse.Cart, User)
	// TODO: set orderID and tableID in context for showing order placed
	store := session_utils.Store
	session, _ := store.Get(r, "payments")
	session.Values["orderID"] = orderId
	session.Values["tableID"] = tableID
	session.Save(r, w)
	// ctx := context.WithValue(r.Context(), types.OrderID{}, orderId)
	// ctx = context.WithValue(ctx, types.TableID{}, tableID)
	// r = r.WithContext(ctx)

	type Res struct {
		Url string `json:"url"`
	}
	temp := Res{
		Url: "/payment",
	}
	json.NewEncoder(w).Encode(temp)
}
