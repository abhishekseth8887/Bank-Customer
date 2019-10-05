package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"myapps/Bank-Customer/entity"
	"myapps/Bank-Customer/storage"
	"net/http"
)


func Customer(w http.ResponseWriter, r *http.Request) {

	log.Println(logtag + " [Customer] Started")

	if r.Method == "POST" {
		var customer entity.Customer
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		err = json.Unmarshal(reqBody, &customer)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		err = storage.Insert(customer)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(customer)
	}

	if r.Method == "GET" {
		customers, err := storage.Get()
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(customers)
	}

	if r.Method == "PUT" {
		var customer entity.Customer
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(reqBody, &customer)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = storage.Update(customer)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(customer)
	}

	if r.Method == "DELETE" {
		var customer entity.Customer
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = json.Unmarshal(reqBody, &customer)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = storage.Delete(customer)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		json.NewEncoder(w).Encode(customer)
}

	log.Println(logtag + " [Customer] Finished")
}