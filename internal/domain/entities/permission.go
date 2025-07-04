package entities

// Permission represents a permission in the system.
// It can have nested permissions to build a tree structure.
type Permission struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Required    []string     `json:"required,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

// HasNestedPermissions returns true if the permission has nested permissions.
func (p Permission) HasNestedPermissions() bool {
	return len(p.Permissions) > 0
}
