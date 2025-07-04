package dto

// PaginationQuery defines parameters used for paginated queries.
// Generic type T allows passing filter structures specific to each use case.
type PaginationQuery[T any] struct {
	Page    int      `json:"page"`
	Limit   int      `json:"limit"`
	Search  string   `json:"search"`
	Sorts   []string `json:"sorts"`
	Filters T        `json:"filters"`
}

// PaginationMeta describes pagination information about a result set.
type PaginationMeta struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Page  int `json:"page"`
}

// PaginatedResult wraps a slice of items with pagination metadata.
type PaginatedResult[T any] struct {
	Data []T            `json:"data"`
	Meta PaginationMeta `json:"meta"`
}
