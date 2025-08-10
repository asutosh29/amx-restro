package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	Database   string
	DbPort     string
}

var Db_config DbConfig

var ValidCategories = []string{}
var ValidStatus = []string{}

var JWTkey string
var SessionSecret string
var PORT int

func InitConfig() {
	var err error
	// Load Env
	godotenv.Load()

	if host := os.Getenv("MYSQL_HOST"); host != "" {
		Db_config.DbHost = host
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		Db_config.DbUser = user
	}
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		Db_config.DbPassword = password
	}
	if database := os.Getenv("MYSQL_DATABASE"); database != "" {
		Db_config.Database = database
	}
	if port := os.Getenv("MYSQL_PORT"); port != "" {
		Db_config.DbPort = port
	}
	// Db_config.DbHost = os.Getenv("MYSQL_HOST")
	// Db_config.DbUser = os.Getenv("MYSQL_USER")
	// Db_config.DbPassword = os.Getenv("MYSQL_PASSWORD")
	// Db_config.Database = os.Getenv("MYSQL_DATABASE")
	// Db_config.DbPort = os.Getenv("MYSQL_PORT")

	ValidCategories = []string{
		"Appetizers",
		"Main Course",
		"Desserts",
		"Beverages",
		"Salads",
		"Soups",
		"Snacks",
		"Combos",
	}
	ValidStatus = []string{"placed", "cooking", "served", "billed", "paid"}

	JWTkey = os.Getenv("JWT_SECRET")
	SessionSecret = os.Getenv("SESSION_SECRET")
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error Invalid Port Format")
	}
}
