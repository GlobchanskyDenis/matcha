package searchFilters

import (
	// "strconv"
	// "time"
)

const (
	onlineFilterType = 1 + iota
	ageFilterType
	ratingFilterType
	interestsFilterType
)

type Filter interface {
	getFilterType() int
	prepareQueryFilter() string
}

type Filters struct {
	filters []Filter
}

func (f Filters) PrepareQuery(sexRestrictions string) string {
	var query = "SELECT * FROM users"
	var queryRestrictions []string

	if sexRestrictions != "" {
		queryRestrictions = append(queryRestrictions, sexRestrictions)
	}
	for _, item := range f.filters {
		if queryRestrictions == nil {
			queryRestrictions = append(queryRestrictions, " WHERE " + item.prepareQueryFilter())
		} else {
			queryRestrictions = append(queryRestrictions, " AND " + item.prepareQueryFilter())
		}
	}
	for _, restrict := range queryRestrictions {
		query += restrict
	}
	return query
}