package dto

// UpdateRoleDto representa los campos permitidos para actualizar un rol.
// Todos los campos son opcionales y los valores nil se ignoran.
type UpdateRoleDto struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Editable    *bool    `json:"editable"`
	Type        *int     `json:"type"`
	Permissions []string `json:"permissions"`
}
