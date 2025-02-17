package repositories

type PaginateInput struct {
	Page  uint64 `json:"page" form:"page"`
	Limit uint64 `json:"limit" form:"limit"`
}

type Pagination[T any] struct {
	Data []*T `json:"data"`
	Meta Meta `json:"meta"`
}

type Meta struct {
	ItemsPerPage uint64 `json:"items_per_page"`
	TotalItems   uint64 `json:"total_items"`
	CurrentPage  uint64 `json:"current_page"`
	TotalPages   uint64 `json:"total_pages"`
}
