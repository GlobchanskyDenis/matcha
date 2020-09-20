package searchFilters

import (
	"MatchaServer/errDef"
	"strconv"
)

type locationFilter struct {
	isMinLatSet  bool
	isMaxLatSet  bool
	isMinLonSet  bool
	isMaxLonSet  bool
	minLatitude  float64
	maxLatitude  float64
	minLongitude float64
	maxLongitude float64
}

func newLocationFilter(in interface{}) (*locationFilter, error) {
	var (
		filter  locationFilter
		ok      bool
		payload map[string]interface{}
		item    interface{}
	)

	// Преобразую полезную нагрузку в нужный нам формат
	payload, ok = in.(map[string]interface{})
	if !ok {
		return nil, errDef.NewArg("неверный тип фильтра локации", "wrong type of location filter")
	}

	// Обрабатываю параметр
	item, ok = payload["minLatitude"]
	if ok {
		filter.isMinLatSet = true
		filter.minLatitude, ok = item.(float64)
		if !ok {
			return nil, errDef.NewArg("неверный тип параметра минимальной широты фильтра локации", 
				"wrong type of min latitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["maxLatitude"]
	if ok {
		filter.isMaxLatSet = true
		filter.maxLatitude, ok = item.(float64)
		if !ok {
			return nil, errDef.NewArg("неверный тип параметра максимальной широты фильтра локации", 
				"wrong type of max latitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["minLongitude"]
	if ok {
		filter.isMinLonSet = true
		filter.minLongitude, ok = item.(float64)
		if !ok {
			return nil, errDef.NewArg("неверный тип параметра минимальной долготы фильтра локации", 
				"wrong type of min longitude parameter")
		}
	}

	// Обрабатываю параметр
	item, ok = payload["maxLongitude"]
	if ok {
		filter.isMaxLonSet = true
		filter.maxLongitude, ok = item.(float64)
		if !ok {
			return nil, errDef.NewArg("неверный тип параметра максимальной долготы фильтра локации", 
				"wrong type of max longitude parameter")
		}
	}

	// Если ни одного параметра не найдено
	if !filter.isMinLatSet && !filter.isMaxLatSet &&
		!filter.isMinLonSet && !filter.isMaxLonSet {
		return nil, errDef.NewArg("не найдены параметры в фильтре локации", "no location parameters found")
	}

	// Валидация широты
	if (filter.isMinLatSet && filter.isMaxLatSet &&
		filter.minLatitude > filter.maxLatitude) ||
		filter.minLatitude < -90.0 || filter.maxLatitude > 90.0 {
		return nil, errDef.NewArg("ошибка параметра широты фильтра локации", "invalid latitude")
	}

	// Валидация долготы
	if (filter.isMinLonSet && filter.isMaxLonSet &&
		filter.minLongitude > filter.maxLongitude) ||
		filter.minLongitude < -180.0 || filter.maxLongitude > 180.0 {
		return nil, errDef.NewArg("ошибка параметра долготы фильтра локации", "invalid longitude")
	}

	return &filter, nil
}

func (f locationFilter) print() string {
	var minLat, maxLat, minLon, maxLon = "not set", "not set", "not set", "not set"

	if f.isMinLatSet {
		minLat = strconv.FormatFloat(f.minLatitude, 'G', -1, 64)
	}
	if f.isMaxLatSet {
		maxLat = strconv.FormatFloat(f.maxLatitude, 'G', -1, 64)
	}
	if f.isMinLonSet {
		minLon = strconv.FormatFloat(f.minLongitude, 'G', -1, 64)
	}
	if f.isMaxLonSet {
		maxLon = strconv.FormatFloat(f.maxLongitude, 'G', -1, 64)
	}
	return "latitude(" + minLat + ".." + maxLat + ") longitude(" +
		minLon + ".." + maxLon + ")"
}

func (locationFilter) getFilterType() int {
	return int(locationFilterType)
}

func (f locationFilter) prepareQueryFilter() string {
	var isfirstExistingParam bool = true
	var dst string

	if f.isMinLatSet {
		if !isfirstExistingParam {
			dst += " AND "
		}
		isfirstExistingParam = false
		dst += "latitude>=" + strconv.FormatFloat(f.minLatitude, 'G', -1, 64)
	}

	if f.isMaxLatSet {
		if !isfirstExistingParam {
			dst += " AND "
		}
		isfirstExistingParam = false
		dst += "latitude<=" + strconv.FormatFloat(f.maxLatitude, 'G', -1, 64)
	}

	if f.isMinLonSet {
		if !isfirstExistingParam {
			dst += " AND "
		}
		isfirstExistingParam = false
		dst += "longitude>=" + strconv.FormatFloat(f.minLongitude, 'G', -1, 64)
	}

	if f.isMaxLonSet {
		if !isfirstExistingParam {
			dst += " AND "
		}
		isfirstExistingParam = false
		dst += "longitude<=" + strconv.FormatFloat(f.maxLongitude, 'G', -1, 64)
	}
	return dst
}
