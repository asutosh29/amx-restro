package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/gorilla/mux"
)

func HandleMakeAdmin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil || userID <= 0 {
		fmt.Println("Invalid user ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid user ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if user exists before processing
	userExists, err := models.UserExistsById(userID)
	if err != nil {
		fmt.Println("Error checking user existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking user existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !userExists {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "User not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	_, err = models.MakeAdminById(userID)
	if err != nil {
		fmt.Println("Error making user admin")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error making user admin",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`user id %v made admin successfully`, userID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleMakeCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["id"])
	if err != nil || userID <= 0 {
		fmt.Println("Invalid user ID format")
		w.WriteHeader(http.StatusBadRequest)
		res := map[string]interface{}{
			"error": "Invalid user ID format",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if user exists before processing
	userExists, err := models.UserExistsById(userID)
	if err != nil {
		fmt.Println("Error checking user existence:", err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error checking user existence",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	if !userExists {
		fmt.Println("User not found")
		w.WriteHeader(http.StatusNotFound)
		res := map[string]interface{}{
			"error": "User not found",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	_, err = models.MakeCustomerById(userID)
	if err != nil {
		fmt.Println("Error making user customer")
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		res := map[string]interface{}{
			"error": "Error making user customer",
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	res := map[string]interface{}{
		"message": fmt.Sprintf(`user id %v made customer successfully`, userID),
	}
	json.NewEncoder(w).Encode(res)
}
