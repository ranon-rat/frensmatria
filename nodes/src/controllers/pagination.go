package controllers

import (
	"math"

	"github.com/ranon-rat/frensmatria/nodes/src/core"
)

var MaxVisiblePages = 7

func calculatePagination(currentPage, totalItems int) Pagination {
	totalPages := int(math.Ceil(float64(totalItems) / float64(core.LIMIT)))

	if currentPage < 1 {
		currentPage = 1
	} else if currentPage > totalPages {
		currentPage = totalPages
	}

	startPage := currentPage - (MaxVisiblePages / 2)
	endPage := currentPage + (MaxVisiblePages / 2)

	if startPage < 1 {
		startPage = 1
		endPage = min(totalPages, MaxVisiblePages)
	}

	if endPage > totalPages {
		endPage = totalPages
		startPage = max(1, totalPages-MaxVisiblePages+1)
	}

	pages := make([]PagesElement, 0)
	for i := startPage; i <= endPage; i++ {
		pages = append(pages, PagesElement{
			Current: i == currentPage,
			Page:    i,
		})
	}

	return Pagination{
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		Pages:       pages,
		Render:      len(pages) != 0,
		Back:        max(currentPage-1, 1),
		Next:        min(currentPage+1, endPage),
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
