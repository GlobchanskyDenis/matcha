package searchFilters

import (
)

type interestsFilter struct {
	interests []string
}

func (interestsFilter) getFilterType() int {
	return int(interestsFilterType)
}

func (f interestsFilter) prepareQueryFilter() string {
	if len(f.interests) == 0 {
		return "1=1"
	}
	if len(f.interests) == 1 {
		return "'"+f.interests[0] + "'=ANY(interests)"
	}

	var queryFilter string
	for _, interest := range f.interests {
		if queryFilter == "" {
			queryFilter = "('"+interest + "'=ANY(interests)"
		} else {
			queryFilter += " AND '"+interest + "'=ANY(interests)"
		}
	}
	queryFilter += ")"
	return queryFilter
}