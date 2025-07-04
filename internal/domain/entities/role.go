package entities

// Role representa un rol de usuario dentro del sistema.
// Define una serie de permisos que describen las acciones permitidas.
type Role struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Type        int      `json:"type"`
	Editable    bool     `json:"editable"`
	Permissions []string `json:"permissions"`
}

// HasPermissionByName devuelve true si el rol posee el permiso indicado.
func (r *Role) HasPermissionByName(name string) bool {
	for _, p := range r.Permissions {
		if p == name {
			return true
		}
	}
	return false
}
