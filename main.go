package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)

	log.Println("App started, listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

//
// Handlers
//
// These are the basic handlers, feel free to make them more interesting if you wish
//

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "goodbye")
}

//
// Middleware
//
// Each endpoint should track/record/report:
// -- starting
// -- completing
// -- total duration
//
