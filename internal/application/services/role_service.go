// Paquete services contiene la lógica de negocio de la aplicación.
// RoleService orquesta las operaciones relacionadas con los roles.
package services

import (
	"errors"                                // manejo de errores con la librería estándar
	"goserver/internal/application/dto"     // estructuras de entrada/salida
	"goserver/internal/application/mappers" // conversión entre DTOs y entidades
	"goserver/internal/domain/entities"
	derrors "goserver/internal/domain/errors"
	"goserver/internal/domain/repositories"
)

// RoleService provee la lógica de negocio relacionada con los roles.

type RoleService struct {
	Repo repositories.RoleRepository
}

// NewRoleService devuelve una nueva instancia de RoleService.
func NewRoleService(r repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: r}
}

// CreateRole persiste un nuevo rol utilizando el repositorio.
func (s *RoleService) CreateRole(input dto.CreateRoleDto) (*entities.Role, error) {
	role := mappers.MapCreateRoleDtoToRole(input)
	created, err := s.Repo.Create(role)
	if err != nil {
		if errors.Is(err, derrors.ErrRoleAlreadyExists) {
			return nil, err
		}
		return nil, err
	}
	return created, nil
}

// FindAllRoles obtiene todos los roles del repositorio.
func (s *RoleService) FindAllRoles() ([]*entities.Role, error) {
	return s.Repo.FindAll()
}

// FindRoleByID busca un rol por su identificador.
func (s *RoleService) FindRoleByID(id string) (*entities.Role, error) {
	role, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role not found")
	}
	return role, nil
}

// UpdateRole actualiza un rol existente con los datos proporcionados.
func (s *RoleService) UpdateRole(id string, input dto.UpdateRoleDto) (*entities.Role, error) {
	role, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role not found")
	}
	updatedRole := mappers.MapUpdateRoleDtoToRole(input, role)
	updated, err := s.Repo.Update(id, updatedRole)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return nil, errors.New("role not found")
	}
	return updated, nil
}

// DeleteRole elimina un rol por su ID.
func (s *RoleService) DeleteRole(id string) (*entities.Role, error) {
	role, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, errors.New("role not found")
	}
	deleted, err := s.Repo.Delete(id)
	if err != nil {
		return nil, err
	}
	if deleted == nil {
		return nil, errors.New("role not found")
	}
	return deleted, nil
}

// GetRoles retorna una lista paginada de roles según la consulta.
func (s *RoleService) GetRoles(q dto.PaginationQuery[dto.RoleFiltersDto]) (*dto.PaginatedResult[*entities.Role], error) {
	return s.Repo.FindPaginated(&q)
}
