package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"goserver/internal/application/dto"
	"goserver/internal/application/services"
	"goserver/internal/domain/entities"
)

// RoleHandler handles HTTP requests for roles.
type RoleHandler struct {
	Service *services.RoleService
}

// NewRoleHandler creates a new RoleHandler instance.
func NewRoleHandler(s *services.RoleService) *RoleHandler {
	return &RoleHandler{Service: s}
}

// Register attaches handler routes to the provided ServeMux.
func (h *RoleHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/roles/permissions", h.GetPermissions)
	mux.HandleFunc("/roles/create", h.Create)
	mux.HandleFunc("/roles", h.Roles)
	mux.HandleFunc("/roles/", h.RoleByID)
}

// writeJSON is a small helper to respond with JSON.
func (h *RoleHandler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// GetPermissions returns the list of available permissions.
func (h *RoleHandler) GetPermissions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	h.writeJSON(w, http.StatusOK, entities.ModulosPermisos)
}

// Roles handles listing and querying roles.
func (h *RoleHandler) Roles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		roles, err := h.Service.FindAllRoles()
		if err != nil {
			h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		h.writeJSON(w, http.StatusOK, roles)
	case http.MethodPost:
		body, _ := ioutil.ReadAll(r.Body)
		var query dto.PaginationQuery[dto.RoleFiltersDto]
		if err := json.Unmarshal(body, &query); err != nil {
			h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		result, err := h.Service.GetRoles(query)
		if err != nil {
			h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		h.writeJSON(w, http.StatusOK, result)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Create registers a new role.
func (h *RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var input dto.CreateRoleDto
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	role, err := h.Service.CreateRole(input)
	if err != nil {
		h.writeJSON(w, http.StatusConflict, map[string]string{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusCreated, role)
}

// RoleByID handles operations on a single role.
func (h *RoleHandler) RoleByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/roles/")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		role, err := h.Service.FindRoleByID(id)
		if err != nil {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		h.writeJSON(w, http.StatusOK, role)
	case http.MethodPut:
		var input dto.UpdateRoleDto
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		role, err := h.Service.UpdateRole(id, input)
		if err != nil {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		h.writeJSON(w, http.StatusOK, role)
	case http.MethodDelete:
		role, err := h.Service.DeleteRole(id)
		if err != nil {
			h.writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		h.writeJSON(w, http.StatusOK, role)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
