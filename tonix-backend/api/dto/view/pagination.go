package view

import "tonix/backend/model"

type PaginationData struct {
	Pages          int `json:"pages"`
	Page           int `json:"page"`
	Count          int `json:"count"`
	ElementsOnPage int `json:"elements_on_page"`
}

type Paginated[T any] struct {
	Pagination PaginationData `json:"pagination"`
	Data       []T            `json:"data"`
}

func ToPaginationData(paginationData *model.PaginationData) *PaginationData {
	return &PaginationData{
		Pages:          paginationData.Pages,
		Page:           paginationData.Page,
		Count:          paginationData.Count,
		ElementsOnPage: paginationData.ElementsOnPage,
	}
}

func ToPaginated[T any](paginationData *model.PaginationData, data []T) *Paginated[T] {
	return &Paginated[T]{
		Pagination: *ToPaginationData(paginationData),
		Data:       data,
	}
}
