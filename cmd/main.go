package main

import (
	"fmt"
	"log"

	"github.com/asutosh29/amx-restro/pkg/api"
	"github.com/asutosh29/amx-restro/pkg/models"
	"github.com/asutosh29/amx-restro/pkg/utils/config"
	"github.com/asutosh29/amx-restro/pkg/utils/session_utils"
	"github.com/asutosh29/amx-restro/pkg/views"
)

func main() {
	fmt.Println("Starting server...")

	// Configuring
	fmt.Println("Configuring server...")
	config.InitConfig()
	views.InitViews()
	session_utils.InitiateStructSession()

	// Initializing database
	fmt.Println("Configuring Database...")
	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// API Routes
	fmt.Println("Configuring API Routers...")
	api.Start()
}
