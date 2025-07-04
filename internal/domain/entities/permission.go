package entities

// Permission describe un permiso en el sistema.
// Puede tener permisos anidados para construir una estructura en árbol.
type Permission struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	Description string       `json:"description,omitempty"`
	Required    []string     `json:"required,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

// HasNestedPermissions devuelve true si existen permisos anidados.
func (p Permission) HasNestedPermissions() bool {
	return len(p.Permissions) > 0
}
