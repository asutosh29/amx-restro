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
	userID, _ := strconv.Atoi(vars["id"])
	models.MakeAdminById(userID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`user id %v made admin successfully`, userID),
	}
	json.NewEncoder(w).Encode(res)
}

func HandleMakeCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	models.MakeCustomerById(userID)
	res := map[string]interface{}{
		"message": fmt.Sprintf(`user id %v made customer successfully`, userID),
	}
	json.NewEncoder(w).Encode(res)
}
