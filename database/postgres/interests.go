package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
)

func (conn ConnDB) AddInterests(unknownInterests []common.Interest) error {
	var query = "INSERT INTO interests (name) VALUES "
	var nameArr = []interface{}{}
	if len(unknownInterests) == 0 {
		return nil
	}
	for nbr, interest := range unknownInterests { ////// ПОХОЖЕ НА ГОВНОКОД. УЗНАТЬ ПОДРОБНЕЕ
		query += "($" + strconv.Itoa(nbr+1) + "), "
		nameArr = append(nameArr, interest.Name) /// УБРАТЬ АЛЛОЦИРОВАНИЕ СЛАЙСА - ПРИНИМАТЬ СЛАЙС ИНТЕРФЕЙСОВ
	}
	query = string(query[:len(query)-2])
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(nameArr...)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	_, err = result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetInterests() ([]common.Interest, error) {
	var interests = []common.Interest{}
	var interest common.Interest

	stmt, err := conn.db.Prepare("SELECT * FROM interests")
	if err != nil {
		return interests, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return interests, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return interests, errors.DatabaseScanError.AddOriginalError(err)
		}
		interests = append(interests, interest)
	}
	return interests, nil
}
