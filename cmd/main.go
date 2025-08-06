package main

import (
	"fmt"
	"log"

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
