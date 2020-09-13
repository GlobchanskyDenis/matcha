package searchFilters

import (
	"strconv"
	"time"
)

type ageFilter struct {
	minAge int
	maxAge int
}

func stringifyDate(dateNum int) string {
	if dateNum > 9 {
		return strconv.Itoa(dateNum)
	}
	return "0" + strconv.Itoa(dateNum)
}

func (ageFilter) getFilterType() int {
	return int(ageFilterType)
}

func (f ageFilter) prepareQueryFilter() string {
	var (
		maxAge, minAge string
		year, day int
		month time.Month
	)

	year, month, day = time.Now().Date()

	minAge = strconv.Itoa(year - f.maxAge + 1) + "-" + 
		stringifyDate(int(month)) + "-" + stringifyDate(day)

	maxAge = strconv.Itoa(year - f.minAge - 1) + "-" + 
		stringifyDate(int(month)) + "-" + stringifyDate(day)

	return "birth<'" + minAge + "' AND birth>'" + maxAge + "'"
}
