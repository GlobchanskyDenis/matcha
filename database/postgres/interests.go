package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
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
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nameArr...)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetInterests() ([]common.Interest, error) {
	var interests = []common.Interest{}
	var interest common.Interest

	stmt, err := conn.db.Prepare("SELECT * FROM interests")
	if err != nil {
		return interests, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		return interests, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return interests, errDef.DatabaseScanError.AddOriginalError(err)
		}
		interests = append(interests, interest)
	}
	return interests, nil
}
