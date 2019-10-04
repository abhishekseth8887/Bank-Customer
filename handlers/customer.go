package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"myapps/bank/entity"
	"myapps/bank/storage"
	"net/http"
)

var singleCustomer entity.Customer

func Customer(w http.ResponseWriter, r *http.Request){

	if r.Method == "POST" {

		// get the body of our POST request
		// unmarshal this into a new Article struct
		// append this to our Articles array.
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &singleCustomer)
		// update our global Articles array to include
		// our new Article
		storage.Insert(singleCustomer)
		json.NewEncoder(w).Encode(singleCustomer)
	}
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(singleCustomer)
	}
	if r.Method == "PUT" {
		fmt.Fprintf(w, "customer updated")
	}
	if r.Method == "DELETE" {
		fmt.Fprintf(w, "customer deleted")
	}
	fmt.Println("Endpoint Hit: customer")
}