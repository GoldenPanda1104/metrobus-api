package main

import (
	"github.com/GoldenPanda1104/metrobus-api/routes"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// port := "3000"
	// if os.Getenv("PORT") != "" {
	// 	port = os.Getenv("PORT")
	// }

	router := mux.NewRouter()

	//En este caso, el parametro w hace referencia al ResponseWriter y el parametro r hace referencia al Request(Reader)
	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)

}
