package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Read the port from the PORT environment variable, falling back to 3000 if it's not set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Register a handler for every path ("/" matches anything not matched by a more specific route).
	// w is where we write the response, r holds the incoming request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from todo app")
	})

	fmt.Printf("Server started in port %s\n", port)

	// Start listening on all interfaces (":3000" means 0.0.0.0:3000, needed inside a container).
	// ListenAndServe blocks forever; it only returns if the server fails, so we log the error and exit.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
