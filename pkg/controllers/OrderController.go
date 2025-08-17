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
	type Res struct {
		Url string `json:"url"`
	}
	var jsonResponse types.OrderRequest
	fmt.Println("Request Body: ", r.Body)
	_ = json.NewDecoder(r.Body).Decode(&jsonResponse)

	User := r.Context().Value("User").(types.User)
	orderId, tableID := models.AddOrder(jsonResponse.Instructions, jsonResponse.Cart, User)
	if tableID == -1 {
		fmt.Println("No Valid Table")
		// TODO:
		session_utils.FlashMsgErr(w, r, "No tables empty currently. Please contact admin", true)
		fmt.Println("Redirecting to menu")
		temp := Res{
			Url: "/menu",
		}
		json.NewEncoder(w).Encode(temp)
		return
	}
	// TODO: set orderID and tableID in context for showing order placed
	store := session_utils.Store
	session, _ := store.Get(r, "payments")

	// Saving
	session.Values["orderID"] = orderId
	session.Values["tableID"] = tableID
	session.Save(r, w)

	temp := Res{
		Url: "/payment",
	}
	json.NewEncoder(w).Encode(temp)
}

func HandleOrderPlaced(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil || orderID <= 0 {
		fmt.Println("Invalid order ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid order ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if order exists
	orderExists, err := models.OrderExistsById(orderID)
	if err != nil {
		fmt.Println("Error checking order existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking order existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !orderExists {
		fmt.Println("Order not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "Order not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	currentStatus, err := models.GetOrderStatusById(orderID)
	if err != nil {
		fmt.Println("Error getting order status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error getting order status",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !models.ValidateOrderStatusTransition(currentStatus, "placed") {
		fmt.Println("Invalid status transition from", currentStatus, "to placed")
		res := map[string]interface{}{
			"error": fmt.Sprintf("Cannot mark order as placed. Current status: %s. Order cannot be placed if it's already paid.", currentStatus),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = models.MarkOrderPlacedById(orderID)
	if err != nil {
		fmt.Println("Error Marking order placed")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error marking order as placed",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is placed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderCooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil || orderID <= 0 {
		fmt.Println("Invalid order ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid order ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if order exists
	orderExists, err := models.OrderExistsById(orderID)
	if err != nil {
		fmt.Println("Error checking order existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking order existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !orderExists {
		fmt.Println("Order not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "Order not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	currentStatus, err := models.GetOrderStatusById(orderID)
	if err != nil {
		fmt.Println("Error getting order status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error getting order status",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !models.ValidateOrderStatusTransition(currentStatus, "cooking") {
		fmt.Println("Invalid status transition from", currentStatus, "to cooking")
		res := map[string]interface{}{
			"error": fmt.Sprintf("Cannot mark order as cooking. Current status: %s. Order cannot be cooking if it's already paid.", currentStatus),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = models.MarkOrderCookingById(orderID)
	if err != nil {
		fmt.Println("Error Marking order cooking")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error marking order as cooking",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is cooking`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderServed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil || orderID <= 0 {
		fmt.Println("Invalid order ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid order ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if order exists
	orderExists, err := models.OrderExistsById(orderID)
	if err != nil {
		fmt.Println("Error checking order existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking order existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !orderExists {
		fmt.Println("Order not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "Order not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	currentStatus, err := models.GetOrderStatusById(orderID)
	if err != nil {
		fmt.Println("Error getting order status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error getting order status",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !models.ValidateOrderStatusTransition(currentStatus, "served") {
		fmt.Println("Invalid status transition from", currentStatus, "to served")
		res := map[string]interface{}{
			"error": fmt.Sprintf("Cannot mark order as served. Current status: %s. Order cannot be served if it's already paid.", currentStatus),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = models.MarkOrderServedById(orderID)
	if err != nil {
		fmt.Println("Error Marking order served")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error marking order as served",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is served`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderBilled(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil || orderID <= 0 {
		fmt.Println("Invalid order ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid order ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if order exists before processing
	orderExists, err := models.OrderExistsById(orderID)
	if err != nil {
		fmt.Println("Error checking order existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking order existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !orderExists {
		fmt.Println("Order not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "Order not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	currentStatus, err := models.GetOrderStatusById(orderID)
	if err != nil {
		fmt.Println("Error getting order status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error getting order status",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !models.ValidateOrderStatusTransition(currentStatus, "billed") {
		fmt.Println("Invalid status transition from", currentStatus, "to billed")
		res := map[string]interface{}{
			"error": fmt.Sprintf("Cannot mark order as billed. Current status: %s. Order must be 'placed', 'cooking', or 'served' to be marked as 'billed'.", currentStatus),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = models.MarkOrderBilledById(orderID)
	if err != nil {
		fmt.Println("Error Marking order billed")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error marking order as billed",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is billed`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleOrderPaid(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID, err := strconv.Atoi(vars["id"])
	if err != nil || orderID <= 0 {
		fmt.Println("Invalid order ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid order ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if order exists before processing
	orderExists, err := models.OrderExistsById(orderID)
	if err != nil {
		fmt.Println("Error checking order existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking order existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !orderExists {
		fmt.Println("Order not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "Order not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	currentStatus, err := models.GetOrderStatusById(orderID)
	if err != nil {
		fmt.Println("Error getting order status:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error getting order status",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !models.ValidateOrderStatusTransition(currentStatus, "paid") {
		fmt.Println("Invalid status transition from", currentStatus, "to paid")
		res := map[string]interface{}{
			"error": fmt.Sprintf("Cannot mark order as paid. Current status: %s. Order must be 'billed' to be marked as 'paid'.", currentStatus),
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	err = models.MarkOrderPaidById(orderID)
	if err != nil {
		fmt.Println("Error Marking order paid")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error marking order as paid",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`order with Id %v is paid`, orderID),
	}
	json.NewEncoder(w).Encode(res)
}
