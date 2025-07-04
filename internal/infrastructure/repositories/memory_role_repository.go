// Paquete repositories contiene implementaciones concretas de los
// repositorios de dominio. En este caso se define una versión en memoria
// que se utiliza para pruebas o aplicaciones pequeñas.
package repositories

import (
	"fmt"     // formateo de cadenas para generar identificadores
	"strings" // utilidades de búsqueda y comparación
	"sync"    // primitiva de exclusión para concurrencia segura

	"goserver/internal/application/dto"
	"goserver/internal/domain/entities"
	derrors "goserver/internal/domain/errors"
)

// MemoryRoleRepository implementa RoleRepository utilizando un mapa en memoria.
// Es seguro para concurrencia y resulta útil en pruebas o aplicaciones pequeñas.
type MemoryRoleRepository struct {
	mu     sync.Mutex
	roles  map[string]*entities.Role
	nextID int
}

// NewMemoryRoleRepository crea un nuevo repositorio de roles en memoria.
func NewMemoryRoleRepository() *MemoryRoleRepository {
	return &MemoryRoleRepository{
		roles:  make(map[string]*entities.Role),
		nextID: 1,
	}
}

// generateID genera un identificador incremental.
func (r *MemoryRoleRepository) generateID() string {
	id := fmt.Sprintf("%d", r.nextID)
	r.nextID++
	return id
}

// Create guarda un nuevo rol en memoria asegurando que el nombre sea único.
func (r *MemoryRoleRepository) Create(role *entities.Role) (*entities.Role, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Check if role name exists
	for _, existing := range r.roles {
		if existing.Name == role.Name {
			return nil, derrors.ErrRoleAlreadyExists
		}
	}
	role.ID = r.generateID()
	r.roles[role.ID] = role
	return role, nil
}

// FindAll devuelve todos los roles almacenados.
func (r *MemoryRoleRepository) FindAll() ([]*entities.Role, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	res := make([]*entities.Role, 0, len(r.roles))
	for _, v := range r.roles {
		res = append(res, v)
	}
	return res, nil
}

// FindByID obtiene un rol por su identificador.
func (r *MemoryRoleRepository) FindByID(id string) (*entities.Role, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	role, ok := r.roles[id]
	if !ok {
		return nil, nil
	}
	return role, nil
}

func (r *MemoryRoleRepository) matchesFilters(role *entities.Role, filters dto.RoleFiltersDto, search string) bool {
	if filters.Type != "" {
		if fmt.Sprintf("%d", role.Type) != filters.Type {
			return false
		}
	}
	if search != "" {
		if !containsIgnoreCase(role.Name, search) {
			return false
		}
	}
	return true
}

func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// FindPaginated devuelve un subconjunto de roles aplicando paginación y filtros.
func (r *MemoryRoleRepository) FindPaginated(q *dto.PaginationQuery[dto.RoleFiltersDto]) (*dto.PaginatedResult[*entities.Role], error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.Limit <= 0 {
		q.Limit = 10
	}
	filtered := make([]*entities.Role, 0)
	for _, role := range r.roles {
		if r.matchesFilters(role, q.Filters, q.Search) {
			filtered = append(filtered, role)
		}
	}
	// sort: not implemented
	start := (q.Page - 1) * q.Limit
	end := start + q.Limit
	if start > len(filtered) {
		start = len(filtered)
	}
	if end > len(filtered) {
		end = len(filtered)
	}
	pageData := filtered[start:end]
	result := &dto.PaginatedResult[*entities.Role]{
		Data: pageData,
		Meta: dto.PaginationMeta{
			Total: len(filtered),
			Limit: q.Limit,
			Pages: int((len(filtered) + q.Limit - 1) / q.Limit),
			Page:  q.Page,
		},
	}
	return result, nil
}

// Update modifica un rol identificado por su id.
func (r *MemoryRoleRepository) Update(id string, data *entities.Role) (*entities.Role, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	existing, ok := r.roles[id]
	if !ok {
		return nil, nil
	}
	existing.Name = data.Name
	existing.Description = data.Description
	existing.Editable = data.Editable
	existing.Type = data.Type
	existing.Permissions = data.Permissions
	return existing, nil
}

// Delete elimina un rol del repositorio.
func (r *MemoryRoleRepository) Delete(id string) (*entities.Role, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	role, ok := r.roles[id]
	if !ok {
		return nil, nil
	}
	delete(r.roles, id)
	return role, nil
}
