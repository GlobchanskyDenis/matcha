package searchFilters

import (
	"MatchaServer/errDef"
)

type interestsFilter struct {
	interests []string
}

func newInterestsFilter(in interface{}) (*interestsFilter, error) {
	var (
		filter  interestsFilter
		ok      bool
		payload []interface{}
		item    interface{}
		str     string
	)

	// Заранее выделю память под некоторое количество интересов
	// Для меньшего количества дополнительных аллокаций
	filter.interests = make([]string, 0, 5)

	// Преобразую полезную нагрузку в нужный нам формат
	payload, ok = in.([]interface{})
	if !ok {
		return nil, errDef.NewArg("неверный тип фильтра интересов", "wrong type of interests filter")
	}

	// Обрабатываю все имеющиеся интересы
	for _, item = range payload {
		str, ok = item.(string)
		if !ok {
			return nil, errDef.NewArg("неверный тип интереса в фильтре интересов", "wrong type of interests item")
		} else {
			filter.interests = append(filter.interests, str)
		}
	}

	// В случае если не было никаких интересов
	if len(payload) == 0 {
		return nil, errDef.NewArg("в фильтре интересов не найдены интересы", "no interests found")
	}

	return &filter, nil
}

func (f interestsFilter) print() string {
	var dst = "interests("
	for i, interest := range f.interests {
		if i != 0 {
			dst += ", "
		}
		dst += interest
	}
	return dst + ")"
}

func (interestsFilter) getFilterType() int {
	return int(interestsFilterType)
}

func (f interestsFilter) prepareQueryFilter() string {
	if len(f.interests) == 0 {
		return "1=1"
	}
	if len(f.interests) == 1 {
		return "'" + f.interests[0] + "'=ANY(interests)"
	}

	var queryFilter string
	for _, interest := range f.interests {
		if queryFilter == "" {
			queryFilter = "('" + interest + "'=ANY(interests)"
		} else {
			queryFilter += " AND '" + interest + "'=ANY(interests)"
		}
	}
	queryFilter += ")"
	return queryFilter
}
