package searchFilters

import (
	"strconv"
	"time"
	"errors"
)

type ageFilter struct {
	isMinSet bool
	isMaxSet bool
	minAge int
	maxAge int
}

func newAgeFilter(in interface{}) (*ageFilter, error) {
	var (
		filter ageFilter
		ok bool
		payload map[string]interface{}
		item interface{}
		fl64 float64
	)
	
	// Преобразую полезную нагрузку в нужный нам формат
	payload, ok = in.(map[string]interface{})
	if !ok {
		return nil, errors.New("wrong type of age filter")
	}

	// Обрабатываю параметр
	item, ok = payload["min"]
	if ok {
		fl64, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of min age parameter")
		}
		filter.isMinSet = true
		filter.minAge = int(fl64)
	}

	// Обрабатываю параметр
	item, ok = payload["max"]
	if ok {
		fl64, ok = item.(float64)
		if !ok {
			return nil, errors.New("wrong type of max age parameter")
		}
		filter.isMaxSet = true
		filter.maxAge = int(fl64)
	}

	// Если ни одного параметра не найдено
	if filter.isMinSet == false && filter.isMaxSet == false {
		return nil, errors.New("no age parameters found")
	}

	// Валидация данных
	if (filter.isMaxSet && filter.isMinSet && filter.minAge > filter.maxAge) ||
		filter.minAge < 0 || filter.maxAge < 0 {
		return nil, errors.New("invalid age parameter")
	}

	return &filter, nil
}

func (f ageFilter) print() string {
	var min, max string

	if !f.isMaxSet {
		max = "not set"
	} else {
		max = strconv.Itoa(f.maxAge)
	}

	if !f.isMinSet {
		min = "not set"
	} else {
		min = strconv.Itoa(f.minAge)
	}
	return "age(" + min + ".." + max + ")"
}

func stringifyDate(dateNum int) string {
	if dateNum > 9 {
		return strconv.Itoa(dateNum)
	}
	return "0" + strconv.Itoa(dateNum)
}

func (ageFilter) getFilterType() int {
	return int(ageFilterType)
}

func (f ageFilter) prepareQueryFilter() string {
	var (
		maxAge, minAge string
		year, day int
		month time.Month
	)

	year, month, day = time.Now().Date()

	if f.isMinSet {
		minAge = strconv.Itoa(year - f.maxAge + 1) + "-" + 
		stringifyDate(int(month)) + "-" + stringifyDate(day)
	}

	if f.isMaxSet {
		maxAge = strconv.Itoa(year - f.minAge - 1) + "-" + 
		stringifyDate(int(month)) + "-" + stringifyDate(day)
	}

	if f.isMinSet && !f.isMaxSet {
		return "birth<'" + minAge + "'"
	}

	if !f.isMinSet && f.isMaxSet {
		return "birth>'" + maxAge + "'"
	}

	return "birth<'" + minAge + "' AND birth>'" + maxAge + "'"
}
