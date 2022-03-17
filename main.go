package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func main() {
	// Setting up of routes.
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	http.HandleFunc("/customers", getAllCustomers)
	// starting the server.
	// listen and serve function returns an error so should always be wrapped with log.fatal
	http.ListenAndServe("localhost:8000", nil)
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ojas Gupta", "Jammu", "180012"},
		{"Ritu Gupta", "Jammu", "180013"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
