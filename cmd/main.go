package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asutosh29/amx-restro/pkg/api"
	"github.com/asutosh29/amx-restro/pkg/models"
)

func main() {
	fmt.Println("Starting server...")
	// Initialize database
	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	api.Start()
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode("Welcome to the server!")
}
