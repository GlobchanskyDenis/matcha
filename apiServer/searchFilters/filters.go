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
	likeFilterType
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
		filter = newOnlineFilter(session)
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
	item, isExist = in["wasntLiked"]
	if isExist {
		filter = newLikeFilter(uid)
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
	/*
	**	Это предварительный шаблон в который остается только вставить параметры поисковых фильтров.
	**	Что тут происходит?
	**	Сначала в таблицу юзеров мы присоединяем аватарки, потом мы добавляем поле uidReceiver, которое
	**	либо содержит значение своего же uid если с этим юзером уже произошел обмен лайками, либо значение
	**	null - если кто-то (или оба) еще не поставили лайк. Фактически в следующем коде это поле используется
	**	для заполнения поля типа bool - разрешено ли общение пользователей или нет
	*/

	var query = `SELECT uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
		bio, avaid, latitude, longitude, interests, status, rating, src, uidSender FROM
	(SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
		bio, avaid, latitude, longitude, interests, status, rating, src FROM
	users LEFT JOIN photos ON avaId = pid)
	AS full_users LEFT JOIN
	(SELECT uidSender FROM
	(SELECT uidSender FROM likes WHERE uidReceiver = ` + strconv.Itoa(f.uid) + `) AS T1 INNER JOIN
	(SELECT uidReceiver FROM likes WHERE uidSender = ` + strconv.Itoa(f.uid) + `) AS T2
		ON T1.uidSender = T2.uidReceiver)
	AS can_talk ON full_users.uid = can_talk.uidSender WHERE uid != ` + strconv.Itoa(f.uid)

	if sexRestrictions != "" {
		query += " AND " + sexRestrictions
	}
	for _, item := range f.filters {
		query += " AND " + item.prepareQueryFilter()
	}
	logger.LogQuery(query)
	return query
}
