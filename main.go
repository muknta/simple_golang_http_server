package main

import (
	"log"
	"net/http"

	"simple_http_server/handlers"
)

func main() {
	http.HandleFunc("/process", handlers.ProcessHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
