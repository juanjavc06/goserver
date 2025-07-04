// Paquete dto define estructuras de datos utilizadas entre capas de la aplicación.
package dto

// CreateRoleDto representa el payload para crear un nuevo rol.
// La validación de los campos se debe realizar en la capa de transporte.
type CreateRoleDto struct {
	Name        string   `json:"name" binding:"required,min=5"`
	Description string   `json:"description"`
	Editable    *bool    `json:"editable"`
	Type        *int     `json:"type"`
	Permissions []string `json:"permissions"`
}
