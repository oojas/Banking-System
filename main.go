package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Setting up of routes.
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	// starting the server.
	// listen and serve function returns an error so should always be wrapped with log.fatal
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
