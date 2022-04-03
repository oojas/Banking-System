package app

import (
	"Banking_System/domain"
	"Banking_System/services"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func Start() {
	p := getPort()
	//mux := http.NewServeMux() // New Multiplexer created which can be used to call handler functions.
	router := mux.NewRouter()
	router.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)
	//[0-9]+ means this api will only respond to numeric id's

	// Wiring
	ch := CustomerHandler{
		service: services.NewCustomerService(domain.NewCustomerRepositoryStub()),
	}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting the server.
	// listen and serve function returns an error so should always be wrapped with log.fatal
	router.HandleFunc("/courses", getAllCourses).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost) // this is a post request where we are creating a new customer.
	log.Fatal(http.ListenAndServe(p, router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	newCustomer := Customer{

		"Rajesh", "Delhi", "603203",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCustomer)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
