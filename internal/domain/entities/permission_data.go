package entities

var (
	access = Permission{Name: "access", Title: "Acceso al M\u00f3dulo"}
	read   = Permission{Name: "read", Title: "Visualizar"}
	create = Permission{Name: "create", Title: "Crear"}
	edit   = Permission{Name: "edit", Title: "Editar"}
	del    = Permission{Name: "delete", Title: "Eliminar"}
)

// ModulosPermisos define los módulos disponibles y sus permisos asociados.
var ModulosPermisos = []Permission{
	{
		Name:        "roles",
		Title:       "Roles de Usuario",
		Description: "Permisos de Roles",
		Required:    []string{"roles", "access"},
		Permissions: []Permission{
			{Name: access.Name, Title: access.Title, Description: "Permite acceder al m\u00f3dulo de roles", Required: []string{"roles", "read"}},
			{Name: read.Name, Title: read.Title, Description: "Permite visualizar los roles de usuario"},
			{Name: create.Name, Title: create.Title, Description: "Permite crear roles de usuario", Required: []string{"roles", "access"}},
			{Name: edit.Name, Title: edit.Title, Description: "Permite editar roles de usuario", Required: []string{"roles", "access"}},
			{Name: del.Name, Title: del.Title, Description: "Permite eliminar roles de usuario", Required: []string{"roles", "access"}},
		},
	},
	{
		Name:        "users",
		Title:       "Usuarios",
		Description: "Permisos de Usuarios",
		Required:    []string{"users", "access"},
		Permissions: []Permission{
			{Name: access.Name, Title: access.Title, Description: "Permite acceder al m\u00f3dulo de usuarios", Required: []string{"users", "read"}},
			{Name: read.Name, Title: read.Title, Description: "Permite visualizar los usuarios"},
			{Name: create.Name, Title: create.Title, Description: "Permite crear usuarios", Required: []string{"users", "access"}},
			{Name: edit.Name, Title: edit.Title, Description: "Permite editar usuarios", Required: []string{"users", "access"}},
			{
				Name:        "TITLE",
				Title:       "Title",
				Description: "Permite eliminar usuarios",
				Permissions: []Permission{
					{Name: access.Name, Title: access.Title, Description: "Permite acceder al m\u00f3dulo de roles"},
					{Name: read.Name, Title: read.Title, Description: "Permite visualizar los roles de usuario"},
					{Name: create.Name, Title: create.Title, Description: "Permite crear roles de usuario"},
					{Name: edit.Name, Title: edit.Title, Description: "Permite editar roles de usuario"},
					{Name: del.Name, Title: del.Title, Description: "Permite eliminar roles de usuario"},
					{Name: del.Name, Title: del.Title, Description: "Solcitude"},
				},
			},
		},
	},
}
