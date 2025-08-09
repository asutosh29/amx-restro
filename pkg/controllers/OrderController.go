package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/types"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/gorilla/mux"
)

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("See All order!")
}

func HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	var jsonResponse types.OrderRequest
	_ = json.NewDecoder(r.Body).Decode(&jsonResponse)

	User := r.Context().Value("User").(types.User)
	orderId, tableID := models.AddOrder(jsonResponse.Instructions, jsonResponse.Cart, User)
	// TODO: set orderID and tableID in context for showing order placed
	store := session_utils.Store
	session, _ := store.Get(r, "payments")

	// Saving
	session.Values["orderID"] = orderId
	session.Values["tableID"] = tableID
	session.Save(r, w)

	type Res struct {
		Url string `json:"url"`
	}
	temp := Res{
		Url: "/payment",
	}
	json.NewEncoder(w).Encode(temp)
}

func HandleOrderPlaced(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	err := models.MarkOrderPlacedById(orderID)
	if err != nil {
		fmt.Println("Error Marking order placed")
		fmt.Println(err)
	}
	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderCooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	_ = models.MarkOrderCookingById(orderID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderServed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	_ = models.MarkOrderServedById(orderID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderBilled(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	_ = models.MarkOrderBilledById(orderID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderPaid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, _ := strconv.Atoi(vars["id"])
	_ = models.MarkOrderPaidById(orderID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}
