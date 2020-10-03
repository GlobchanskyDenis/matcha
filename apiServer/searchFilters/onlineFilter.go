package searchFilters

import (
	"MatchaServer/session"
	"strconv"
)

type onlineFilter struct {
	uidSlice []int
}

func newOnlineFilter(session *session.Session) (*onlineFilter, error) {
	var filter onlineFilter

	// Получаю слайс пользователей
	filter.uidSlice = session.GetLoggedUsersUidSlice()

	return &filter, nil
}

func (f onlineFilter) print() string {
	return "online"
}

func (onlineFilter) getFilterType() int {
	return int(onlineFilterType)
}

func (f onlineFilter) prepareQueryFilter() string {
	if len(f.uidSlice) == 0 {
		return "uid=0"
	}
	if len(f.uidSlice) == 1 {
		return "uid=" + strconv.Itoa(f.uidSlice[0])
	}

	var queryFilter string
	for _, uid := range f.uidSlice {
		if queryFilter == "" {
			queryFilter = "(uid=" + strconv.Itoa(uid)
		} else {
			queryFilter += " OR uid=" + strconv.Itoa(uid)
		}
	}
	queryFilter += ")"
	return queryFilter
}
