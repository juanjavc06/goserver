package services

import (
	"errors"
	"goserver/internal/application/dto"
	"goserver/internal/application/mappers"
	"goserver/internal/domain/entities"
	derrors "goserver/internal/domain/errors"
	"goserver/internal/domain/repositories"
)

// RoleService provides business logic around roles

type RoleService struct {
	Repo repositories.RoleRepository
}

// NewRoleService creates a new RoleService instance.
func NewRoleService(r repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: r}
}

// CreateRole persists a new role using the repository.
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

// FindAllRoles retrieves all roles from the repository.
func (s *RoleService) FindAllRoles() ([]*entities.Role, error) {
	return s.Repo.FindAll()
}

// FindRoleByID fetches a role by its identifier.
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

// UpdateRole updates an existing role with the provided data.
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

// DeleteRole removes a role by ID.
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

// GetRoles returns a paginated list of roles matching the query.
func (s *RoleService) GetRoles(q dto.PaginationQuery[dto.RoleFiltersDto]) (*dto.PaginatedResult[*entities.Role], error) {
	return s.Repo.FindPaginated(&q)
}
