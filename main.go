// Programa principal que inicia el servidor HTTP y conecta las piezas de la
// aplicación. Este archivo demuestra un ejemplo sencillo de cómo inicializar
// dependencias y exponer endpoints REST en Go.
package main

import (
	"log"
	"net/http"

	"goserver/internal/application/services"
	repo "goserver/internal/infrastructure/repositories"
	handlers "goserver/internal/interfaces/http"
)

// main configura las dependencias y arranca el servidor HTTP.
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
