package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// DSN hace referencia a Database string

// Remote DB
var dsn = "host=ep-wispy-queen-49748303.us-east-2.aws.neon.tech user=fl0user password=9YwNQ3XVcsbA dbname=metrobus-dev port=5432"

// Local DB
// var dsn = "host=localhost user=admin password=admin dbname=metrobus_dev port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Connection to database established")
	}
}
