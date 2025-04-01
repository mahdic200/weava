package Http

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PaginationMetadata struct {
	TotalPages   int  `json:"totalPages"`
	PreviousPage *int `json:"previousPage"`
	NextPage     *int `json:"nextPage"`
	Offset       int  `json:"offset"`
	LimitPerPage int  `json:"limitPerPage"`
	CurrentPage  int  `json:"currentPage"`
}

func PaginateMetadata(page int64, count int64, limit int64) PaginationMetadata {
	total_pages := int(math.Ceil(float64(count) / float64(limit)))
	current := int(page)
	if current > total_pages {
		current = total_pages
	}
	var prev *int
	if current-1 > 0 {
		temp := current - 1
		prev = &temp
	} else {
		prev = nil
	}
	var next *int
	if current+1 <= total_pages {
		temp := current + 1
		next = &temp
	} else {
		next = nil
	}
	var offset int
	if current-1 > 0 {
		offset = (current - 1) * int(limit)
	} else {
		offset = 0
	}
	return PaginationMetadata{
		TotalPages:   total_pages,
		PreviousPage: prev,
		NextPage:     next,
		Offset:       offset,
		LimitPerPage: int(limit),
		CurrentPage:  current,
	}
}

func getPageQueryArg(c *fiber.Ctx) int64 {
	query_arg := ""
	if page := GetQueryArg(c, "page"); page != nil {
		query_arg = *page
	}
	page, err := strconv.ParseInt(query_arg, 10, 64)
	if err != nil {
		return 1
	}
	if page == 0 {
		page = 1
	}
	return page
}

func Paginate(tx *gorm.DB, c *fiber.Ctx, limit uint8) (*gorm.DB, PaginationMetadata) {
	page := getPageQueryArg(c)
	var count int64
	tx.Count(&count)
	metadata := PaginateMetadata(page, count, int64(limit))
	tx = tx.Limit(metadata.LimitPerPage).Offset(metadata.Offset)
	return tx, metadata
}
