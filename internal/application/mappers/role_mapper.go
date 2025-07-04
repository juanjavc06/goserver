package mappers

import (
	"goserver/internal/application/dto"
	"goserver/internal/domain/entities"
)

// MapCreateRoleDtoToRole converts a CreateRoleDto into a Role entity.
// It applies default values for optional fields.
func MapCreateRoleDtoToRole(d dto.CreateRoleDto) *entities.Role {
	editable := true
	if d.Editable != nil {
		editable = *d.Editable
	}
	tp := 0
	if d.Type != nil {
		tp = *d.Type
	}
	return &entities.Role{
		Name:        d.Name,
		Description: d.Description,
		Editable:    editable,
		Type:        tp,
		Permissions: d.Permissions,
	}
}

// MapUpdateRoleDtoToRole copies values from UpdateRoleDto into an existing Role.
// Only fields provided in the DTO are updated.
func MapUpdateRoleDtoToRole(d dto.UpdateRoleDto, existing *entities.Role) *entities.Role {
	if d.Name != nil {
		existing.Name = *d.Name
	}
	if d.Description != nil {
		existing.Description = *d.Description
	}
	if d.Editable != nil {
		existing.Editable = *d.Editable
	}
	if d.Type != nil {
		existing.Type = *d.Type
	}
	if d.Permissions != nil {
		existing.Permissions = d.Permissions
	}
	return existing
}
