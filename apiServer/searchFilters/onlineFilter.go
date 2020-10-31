package searchFilters

import (
	"MatchaServer/session"
	"strconv"
	"strings"
)

type onlineFilter struct {
	uidSlice []int
}

func newOnlineFilter(session *session.Session) *onlineFilter {
	var filter onlineFilter

	// Получаю слайс пользователей
	filter.uidSlice = session.GetLoggedUsersUidSlice()

	return &filter
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

	var uidArray []string
	for _, uid := range f.uidSlice {
		uidArray = append(uidArray, strconv.Itoa(uid))
	}
	return "(uid IN (" + strings.Join(uidArray, ",") + "))"
}
