package main

import (
	"github.com/gorilla/mux"
	"log"
	appHandlers "myapps/bank/handlers"
	"net/http"
)

func handleRequests() {

	//r := http.NewServeMux()
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc


	myRouter.HandleFunc("/customer", appHandlers.Customer)

	myRouter.HandleFunc("/", appHandlers.HomePage)
	//r.HandleFunc("/all", returnAllArticles)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10001", myRouter))
}

func main() {
	handleRequests()
}
