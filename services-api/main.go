package main

import (
	"fmt"
	"log"
	"net/http"
	"services-api/handlers"
)

func main() {
	handlers.InitializeData()

	// handing endpoints
	http.HandleFunc("/api/services", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetServices(w, r)
		case http.MethodPost:
			handlers.CreateService(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/services/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetService(w, r)
		case http.MethodPut:
			handlers.UpdateService(w, r)
		case http.MethodDelete:
			handlers.DeleteService(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/services/name/", handlers.GetServiceByName)

	// starting server
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
