package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Simon")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "simon", City: "Paris", Zipcode: "75015"},
	}
	json.NewEncoder(w).Encode(customers)
}
