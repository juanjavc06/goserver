package dto

// PaginationQuery define los parámetros utilizados en consultas paginadas.
// El tipo genérico T permite pasar estructuras de filtros específicas.
type PaginationQuery[T any] struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Search  string   `json:"search"`
	Sorts   []string `json:"sorts"`
	Filters T        `json:"filters"`
}

// PaginationMeta describe información de paginación del conjunto de resultados.
type PaginationMeta struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Page  int `json:"page"`
}

// PaginatedResult envuelve un listado de elementos junto con la metainformación de paginación.
type PaginatedResult[T any] struct {
	Data []T            `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
