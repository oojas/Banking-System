package main

import (
	"fmt"
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	http.HandleFunc("/customers", getAllCustomers)
	// starting the server.
	// listen and serve function returns an error so should always be wrapped with log.fatal
	http.HandleFunc("/courses", getAllCourses)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
