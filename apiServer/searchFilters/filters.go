package searchFilters

import (
	"MatchaServer/apiServer/logger"
	"MatchaServer/database"
	"MatchaServer/errors"
	"MatchaServer/session"
	"strconv"
)

const (
	onlineFilterType = 1 + iota
	ageFilterType
	ratingFilterType
	interestsFilterType
	locationFilterType
	radiusFilterType
)

type Filter interface {
	getFilterType() int
	prepareQueryFilter() string
	print() string
}

type Filters struct {
	uid     int
	filters []Filter
}

func New() *Filters {
	return &Filters{}
}

func (f *Filters) Parse(in map[string]interface{}, uid int,
	connDB database.Storage, session *session.Session) error {
	var (
		filter  Filter
		err     error
		isExist bool
		item    interface{}
	)

	if session == nil {
		return errors.NewArg("найден пустой указатель сессии", "empty session pointer found")
	}
	item, isExist = in["age"]
	if isExist {
		filter, err = newAgeFilter(item)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	item, isExist = in["interests"]
	if isExist {
		filter, err = newInterestsFilter(item)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	item, isExist = in["location"]
	if isExist {
		filter, err = newLocationFilter(item)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	_, isExist = in["online"]
	if isExist {
		filter, err = newOnlineFilter(session)
		if err != nil {
			return nil
		}
		f.filters = append(f.filters, filter)
	}
	item, isExist = in["rating"]
	if isExist {
		filter, err = newRatingFilter(item)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	item, isExist = in["radius"]
	if isExist {

		filter, err = newRadiusFilter(item, uid, connDB)
		if err != nil {
			return err
		}
		f.filters = append(f.filters, filter)
	}
	f.uid = uid
	return nil
}

func (f *Filters) Print() string {
	var dst string
	for i, item := range f.filters {
		if i == 0 {
			dst += item.print()
		} else {
			dst += " " + item.print()
		}
	}
	return dst
}

func (f *Filters) PrepareQuery(sexRestrictions string, logger *logger.Logger) string {
	var query = "SELECT * FROM users LEFT JOIN (SELECT uidReceiver FROM likes WHERE uidSender=1) AS tmp ON users.uid = tmp.uidReceiver WHERE uid!=" + strconv.Itoa(f.uid)

	if sexRestrictions != "" {
		query += " AND " + sexRestrictions
	}
	for _, item := range f.filters {
		query += " AND " + item.prepareQueryFilter()
	}
	logger.LogQuery(query)
	return query
}
