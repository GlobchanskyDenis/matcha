package searchFilters

import (
	"errors"
	"strconv"
)

type ratingFilter struct {
	isMinSet  bool
	isMaxSet  bool
	minRating int
	maxRating int
}

func newRatingFilter(in interface{}) (*ratingFilter, error) {
	var (
		filter  ratingFilter
		ok      bool
		payload map[string]interface{}
		item    interface{}
		fl64    float64
	)

	// Преобразую полезную нагрузку в нужный нам формат
	payload, ok = in.(map[string]interface{})
	if !ok {
		return nil, errors.New("wrong type of rating filter")
	}

	// Обрабатываю параметр
	item, ok = payload["min"]
	if ok {
		fl64, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of min rating parameter")
		}
		filter.isMinSet = true
		filter.minRating = int(fl64)
	}

	// Обрабатываю параметр
	item, ok = payload["max"]
	if ok {
		fl64, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of max rating parameter")
		}
		filter.isMaxSet = true
		filter.maxRating = int(fl64)
	}

	// Если ни одного параметра не найдено
	if !filter.isMinSet && !filter.isMaxSet {
		return nil, errors.New("no rating parameters found")
	}

	// Валидация данных
	if filter.minRating < 0 || filter.maxRating < 0 ||
		(filter.isMaxSet && filter.isMinSet && filter.maxRating < filter.minRating) {
		return nil, errors.New("invalid rating parameter")
	}

	return &filter, nil
}

func (f ratingFilter) print() string {
	if f.isMinSet && !f.isMaxSet {
		return "rating(" + strconv.Itoa(f.minRating) + "..not set)"
	}
	if !f.isMinSet && f.isMaxSet {
		return "rating(not set.." + strconv.Itoa(f.maxRating) + ")"
	}
	return "rating(" + strconv.Itoa(f.minRating) + ".." + strconv.Itoa(f.maxRating) + ")"
}

func (ratingFilter) getFilterType() int {
	return int(ratingFilterType)
}

func (f ratingFilter) prepareQueryFilter() string {
	if f.isMinSet && !f.isMaxSet {
		return "rating>=" + strconv.Itoa(f.minRating)
	}
	if !f.isMinSet && f.isMaxSet {
		return "rating<=" + strconv.Itoa(f.maxRating)
	}
	return "rating>=" + strconv.Itoa(f.minRating) + " AND rating<=" + strconv.Itoa(f.maxRating)
}
