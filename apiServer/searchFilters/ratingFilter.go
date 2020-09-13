package searchFilters

import (
	"strconv"
)

type ratingFilter struct {
	minRating int
	maxRating int
}

func (ratingFilter) getFilterType() int {
	return int(ratingFilterType)
}

func (f ratingFilter) prepareQueryFilter() string {
	return "rating<=" + strconv.Itoa(f.maxRating) + " AND rating>=" + strconv.Itoa(f.minRating)
}