package app

import (
	"encoding/json"
	"net/http"
)

type Courses struct {
	Name string
	Link string
}

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := []Courses{
		{"flutter", "https://google.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
		{"Hadoop", "https:udemy.com"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ojas Gupta", "Jammu", "180012"},
		{"Ritu Gupta", "Jammu", "180013"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
