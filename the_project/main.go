package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Embed the templates folder into the binary, so the final "scratch" image doesn't need any extra files.
//
//go:embed templates/index.html
var templates embed.FS

// pageData holds the values that the HTML template fills in.
type pageData struct {
	Title string
	Port  string
}

func main() {
	// Read the port from the PORT environment variable, falling back to 3000 if it's not set.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Parse the template once at startup; template.Must panics if the HTML is invalid.
	tmpl := template.Must(template.ParseFS(templates, "templates/index.html"))

	// Register a handler for every path ("/" matches anything not matched by a more specific route).
	// w is where we write the response, r holds the incoming request.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := pageData{
			Title: "Todo App",
			Port:  port,
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("Server started in port %s\n", port)

	// Start listening on all interfaces (":3000" means 0.0.0.0:3000, needed inside a container).
	// ListenAndServe blocks forever; it only returns if the server fails, so we log the error and exit.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
