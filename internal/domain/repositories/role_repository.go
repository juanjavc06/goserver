package repositories

import (
	"goserver/internal/application/dto"
	"goserver/internal/domain/entities"
)

// RoleRepository defines interface for role persistence

type RoleRepository interface {
	Create(role *entities.Role) (*entities.Role, error)
	FindAll() ([]*entities.Role, error)
	FindByID(id string) (*entities.Role, error)
	FindPaginated(query *dto.PaginationQuery[dto.RoleFiltersDto]) (*dto.PaginatedResult[*entities.Role], error)
	Update(id string, role *entities.Role) (*entities.Role, error)
	Delete(id string) (*entities.Role, error)
}
