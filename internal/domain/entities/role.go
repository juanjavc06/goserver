package entities

// Role represents a user role in the system.
// It defines a list of permissions that describe allowed actions.
type Role struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Type        int      `json:"type"`
	Editable    bool     `json:"editable"`
	Permissions []string `json:"permissions"`
}

// HasPermissionByName returns true if the role has the given permission.
func (r *Role) HasPermissionByName(name string) bool {
	for _, p := range r.Permissions {
		if p == name {
			return true
		}
	}
	return false
}
