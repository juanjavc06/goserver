package dto

// RoleFiltersDto contiene los criterios de filtrado para listar roles.
// Se pueden agregar más campos de ser necesario.
type RoleFiltersDto struct {
	Type string `json:"type"`
}
