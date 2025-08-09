package config

import (
	"os"

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

func InitConfig() {
	// Load Env
	godotenv.Load()

	Db_config.DbHost = os.Getenv("MYSQL_HOST")
	Db_config.DbUser = os.Getenv("MYSQL_USER")
	Db_config.DbPassword = os.Getenv("MYSQL_PASSWORD")
	Db_config.Database = os.Getenv("MYSQL_DATABASE")
	Db_config.DbPort = os.Getenv("MYSQL_PORT")

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
}
