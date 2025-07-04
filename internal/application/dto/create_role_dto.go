package dto

// CreateRoleDto represents the payload to create a new role.
// Validation of fields should be handled by the transport layer.
type CreateRoleDto struct {
	Name        string   `json:"name" binding:"required,min=5"`
	Description string   `json:"description"`
	Editable    *bool    `json:"editable"`
	Type        *int     `json:"type"`
	Permissions []string `json:"permissions"`
}
