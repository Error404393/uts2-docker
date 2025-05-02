package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!\n", name)
	fmt.Fprintf(w, "This is Task For Docker - Cloud Computing!\n")
	fmt.Fprintf(w, "This is a simple Dockerized Go app\n")
}