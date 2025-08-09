package config

import "os"

var DbHost = os.Getenv("MYSQL_HOST")
var DbUser = os.Getenv("MYSQL_USER")
var DbPassword = os.Getenv("MYSQL_PASSWORD")
var Database = os.Getenv("MYSQL_DATABASE")
var DbPort = os.Getenv("MYSQL_PORT")

var ValidCategories = []string{
	"Appetizers",
	"Main Course",
	"Desserts",
	"Beverages",
	"Salads",
	"Soups",
	"Snacks",
	"Combos",
}
var ValidStatus = []string{"placed", "cooking", "served", "billed", "paid"}

var JWTkey = os.Getenv("JWT_SECRET")
var SessionSecret = os.Getenv("SESSION_SECRET")
