package app

import (
	"Banking_System/services"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Courses struct {
	Name string
	Link []byte
}

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}
type CustomerHandler struct { // rest handler
	service services.CustomerService
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	courses := []Courses{
		{"flutter", ImageToBase64("images/business.jpg")},
		{"Hadoop", ImageToBase64("images/hcl.jpg")},
		{"Hadoop", ImageToBase64("images/business.jpg")},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
func ImageToBase64(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return bytes

}
