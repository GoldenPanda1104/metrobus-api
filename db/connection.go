package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DSN hace referencia a Database string

// Remote DB
var dsn = "postgres://fl0user:BtYGCsu1aj6O@ep-misty-morning-21300543.us-east-2.aws.neon.tech:5432/metrobusdev?sslmode=require"

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
