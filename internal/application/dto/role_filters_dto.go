package dto

// RoleFiltersDto contains filtering criteria for listing roles.
// Additional filter fields can be added as needed.
type RoleFiltersDto struct {
	Type string `json:"type"`
}
