package model

type PaginatedResult struct {
	Page   *Page
	Result interface{}
}

type Page struct {
	Current *int `json:"current,omitempty"`
	Next    *int `json:"next,omitempty"`
	Total   *int `json:"total,omitempty"`
}

type Response struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Page       `json:"pagination,omitempty"`
}
