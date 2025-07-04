package dto

// UpdateRoleDto represents the allowed fields for updating a role.
// All fields are optional and nil values will be ignored.
type UpdateRoleDto struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Editable    *bool    `json:"editable"`
	Type        *int     `json:"type"`
	Permissions []string `json:"permissions"`
}
