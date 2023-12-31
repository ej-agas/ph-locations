package stores

import "github.com/ej-agas/ph-locations/models"

type Collection[T models.Model] struct {
	Data           []T            `json:"data"`
	PaginationInfo PaginationInfo `json:"pagination"`
}

type PaginationInfo struct {
	Total       int `json:"total"`
	TotalPages  int `json:"total_pages"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
}
