package app

import (
	"Banking_System/services"
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
type CustomerHandler struct {
	service services.CustomerService
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
func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
