package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"goserver/internal/application/dto"
	"goserver/internal/application/services"
	repo "goserver/internal/infrastructure/repositories"
)

// TestRoleHandler_BasicFlow prueba el ciclo completo de manejo de roles por HTTP.
func TestRoleHandler_BasicFlow(t *testing.T) {
	repository := repo.NewMemoryRoleRepository()
	service := services.NewRoleService(repository)
	handler := NewRoleHandler(service)

	mux := http.NewServeMux()
	handler.Register(mux)

	// Crear un rol
	createBody, _ := json.Marshal(dto.CreateRoleDto{Name: "admin"})
	req := httptest.NewRequest(http.MethodPost, "/roles/create", bytes.NewReader(createBody))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rec.Code)
	}
	var created map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	id, ok := created["id"].(string)
	if !ok || id == "" {
		t.Fatalf("expected returned role to have id")
	}

	// Listar roles
	req = httptest.NewRequest(http.MethodGet, "/roles", nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	var list []*map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &list); err != nil {
		t.Fatalf("invalid list json: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 role in list, got %d", len(list))
	}

	// Obtener rol por ID
	req = httptest.NewRequest(http.MethodGet, "/roles/"+id, nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	// Eliminar rol
	req = httptest.NewRequest(http.MethodDelete, "/roles/"+id, nil)
	rec = httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 on delete, got %d", rec.Code)
	}
}
