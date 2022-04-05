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
		{"flutter", ImageToBase64("C:/Users/Ojas Gupta/Downloads/business.jpg")},
		{"Hadoop", ImageToBase64("C:/Users/Ojas Gupta/Downloads/business.jpg")},
		{"Hadoop", ImageToBase64("C:/Users/Ojas Gupta/Downloads/business.jpg")},
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

	//var base64Encoding string
	//
	//// Determine the content type of the image file
	//mimeType := http.DetectContentType(bytes)
	//
	//// Prepend the appropriate URI scheme header depending
	//// on the MIME type
	//switch mimeType {
	//case "image/jpeg":
	//	base64Encoding += "data:image/jpeg;base64,"
	//case "image/png":
	//	base64Encoding += "data:image/png;base64,"
	//}
	//
	//// Append the base64 encoded output
	//base64Encoding += toBase64(bytes)
	//// Print the full base64 representation of the image
	return bytes

}
