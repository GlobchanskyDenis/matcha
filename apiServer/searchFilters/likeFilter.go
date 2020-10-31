package searchFilters

import (
	"strconv"
)

type likeFilter struct {
	uid int
}

func newLikeFilter(uid int) *likeFilter {
	var filter likeFilter

	filter.uid = uid
	return &filter
}

func (f likeFilter) print() string {
	return "wasntLiked"
}

func (likeFilter) getFilterType() int {
	return int(likeFilterType)
}

func (f likeFilter) prepareQueryFilter() string {
	return "uid NOT IN (SELECT uidReceiver FROM likes WHERE uidSender = " + strconv.Itoa(f.uid) + ")"
}