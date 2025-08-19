package main

import (
	"fmt"
	"log"
	"net/http"
)

// helloHandler responds with "Hello World" to any request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "Hello World")
}

func main() {
	// Handle all routes with the hello handler
	http.HandleFunc("/", helloHandler)
	
	port := ":3000"
	log.Printf("Starting Go server on port %s", port)
	log.Printf("Server will respond with 'Hello World' to any request")
	
	// Start the server
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}