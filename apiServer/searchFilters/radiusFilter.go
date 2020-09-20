package searchFilters

import (
	"MatchaServer/database"
	"errors"
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
		return nil, errors.New("wrong type of new age filter")
	}

	// Обрабатываю параметр
	item, ok = payload["latitude"]
	if ok {
		isLatSet = true
		filter.latitude, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of latitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["longitude"]
	if ok {
		isLonSet = true
		filter.longitude, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of longitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["radius"]
	if ok {
		filter.radius, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of longitude parameter")
		}
	} else {
		return nil, errors.New("radius parameter expected")
	}

	// Если ни одного параметра не найдено
	if (!isLatSet && isLonSet) || (isLatSet && !isLonSet) {
		return nil, errors.New("incomplete parameter found")
	}

	// Валидация радиуса
	if filter.radius <= 0 {
		return nil, errors.New("invalid radius")
	}

	// Валидация широты
	if isLatSet && (filter.latitude < -90 || filter.latitude > 90) {
		return nil, errors.New("invalid latitude")
	}

	// Валидация долготы
	if isLonSet && (filter.longitude < -180 || filter.longitude > 180) {
		return nil, errors.New("invalid longitude")
	}

	// Если координаты пользователя не пришли в запросе
	if !isLatSet && !isLonSet {
		println("\033[33mCalling function 'get user by uid'\033[m")
		user, err := connDB.GetUserByUid(uid)
		if err != nil {
			return nil, errors.New("Error while connecting database: " + err.Error())
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
