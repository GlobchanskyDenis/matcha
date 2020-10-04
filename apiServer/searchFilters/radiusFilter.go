package searchFilters

import (
	"MatchaServer/database"
	"MatchaServer/errors"
	"strconv"
)

type radiusFilter struct {
	latitude  float64
	longitude float64
	radius    float64
}

func newRadiusFilter(in interface{}, uid int, connDB database.Storage) (*radiusFilter, error) {
	var (
		isLatSet bool
		isLonSet bool
		filter   radiusFilter
		ok       bool
		payload  map[string]interface{}
		item     interface{}
	)

	// Преобразую полезную нагрузку в нужный нам формат
	payload, ok = in.(map[string]interface{})
	if !ok {
		return nil, errors.NewArg("неверный тип фильтра радиуса", "wrong type of radius filter")
	}

	// Обрабатываю параметр
	item, ok = payload["latitude"]
	if ok {
		isLatSet = true
		filter.latitude, ok = item.(float64)
		if !ok {
			return nil, errors.NewArg("неверный тип параметра широты фильтра радиуса", "wrong type of latitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["longitude"]
	if ok {
		isLonSet = true
		filter.longitude, ok = item.(float64)
		if !ok {
			return nil, errors.NewArg("неверный тип параметра долготы фильтра радиуса", "wrong type of longitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["radius"]
	if ok {
		filter.radius, ok = item.(float64)
		if !ok {
			return nil, errors.NewArg("неверный тип параметра радиуса фильтра радиуса", "wrong type of radius parameter")
		}
	} else {
		return nil, errors.NewArg("не найден параметр радиуса фильтра радиуса", "radius parameter expected")
	}

	// Если ни одного параметра не найдено
	if (!isLatSet && isLonSet) || (isLatSet && !isLonSet) {
		return nil, errors.NewArg("не верно заполнены параметры фильтра радиуса", "incomplete parameter found")
	}

	// Валидация радиуса
	if filter.radius <= 0 {
		return nil, errors.NewArg("ошибка параметра радиус фильтра радиуса", "invalid radius")
	}

	// Валидация широты
	if isLatSet && (filter.latitude < -90 || filter.latitude > 90) {
		return nil, errors.NewArg("ошибка параметра широты фильтра радиуса", "invalid latitude")
	}

	// Валидация долготы
	if isLonSet && (filter.longitude < -180 || filter.longitude > 180) {
		return nil, errors.NewArg("ошибка параметра долготы фильтра радиуса", "invalid longitude")
	}

	// Если координаты пользователя не пришли в запросе
	if !isLatSet && !isLonSet {
		user, err := connDB.GetUserByUid(uid)
		if err != nil {
			return nil, errors.NewArg("ошибка соединения с базой данных", "database connecting error").AddOriginalError(err)
		}
		filter.latitude = float64(user.Latitude)
		filter.longitude = float64(user.Longitude)
	}

	return &filter, nil
}

func (f radiusFilter) print() string {
	return "radius " + strconv.FormatFloat(f.radius, 'G', -1, 64) + "km coordinates(" +
		strconv.FormatFloat(f.latitude, 'G', -1, 64) + ".." +
		strconv.FormatFloat(f.longitude, 'G', -1, 64) + ")"
}

func (radiusFilter) getFilterType() int {
	return int(radiusFilterType)
}

func (f radiusFilter) prepareQueryFilter() string {
	var minLat, minLon, maxLat, maxLon string

	minLat = strconv.FormatFloat(f.latitude-f.radius/111.0, 'G', -1, 64)
	maxLat = strconv.FormatFloat(f.latitude+f.radius/111.0, 'G', -1, 64)
	minLon = strconv.FormatFloat(f.longitude-f.radius/111.0, 'G', -1, 64)
	maxLon = strconv.FormatFloat(f.longitude+f.radius/111.0, 'G', -1, 64)
	return "latitude>=" + minLat + " AND latitude<=" + maxLat +
		" AND longitude>=" + minLon + " AND longitude<=" + maxLon
}
