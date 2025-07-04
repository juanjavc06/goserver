package main

import (
	"log"
	"net/http"

	"goserver/internal/application/services"
	repo "goserver/internal/infrastructure/repositories"
	handlers "goserver/internal/interfaces/http"
)

// main wires the application and starts the HTTP server.
func main() {
	repository := repo.NewMemoryRoleRepository()
	service := services.NewRoleService(repository)
	handler := handlers.NewRoleHandler(service)

	mux := http.NewServeMux()
	handler.Register(mux)

	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
