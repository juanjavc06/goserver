// Paquete mappers contiene funciones que transforman DTOs en entidades y
// viceversa.
package mappers

import (
	"goserver/internal/application/dto"
	"goserver/internal/domain/entities"
)

// MapCreateRoleDtoToRole convierte un CreateRoleDto en una entidad Role.
// Aplica valores por defecto para los campos opcionales.
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

// MapUpdateRoleDtoToRole copia valores desde UpdateRoleDto sobre un Role existente.
// Solo se actualizan los campos que vengan definidos en el DTO.
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
