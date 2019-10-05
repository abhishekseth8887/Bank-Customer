package main

import (
	"github.com/gorilla/mux"
	"log"
	appHandlers "myapps/Bank-Customer/handlers"
	"net/http"
)

func handleRequests() {

	log.Println("[handleRequests] Started")

	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/customer", appHandlers.Customer)

	log.Fatal(http.ListenAndServe(":8088", myRouter))
}

func main() {
	log.Println("[main] server started")
	handleRequests()
	log.Println("[main] server stopped")
}
