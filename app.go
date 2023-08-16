package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GoldenPanda1104/metrobus-api/db"
	"github.com/GoldenPanda1104/metrobus-api/models"
	"github.com/GoldenPanda1104/metrobus-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	db.DBConnection()

	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Card{})

	router := mux.NewRouter()

	//En este caso, el parametro w hace referencia al ResponseWriter y el parametro r hace referencia al Request(Reader)
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	log.Print("Application listening on port " + port)
	http.ListenAndServe(":"+port, router)

}
