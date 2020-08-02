package postgres

import (
	"MatchaServer/config"
	"errors"
	"strconv"
)

func (conn ConnDB) AddInterests(unknownInterests []config.Interest) error {
	var query = "INSERT INTO interests (name) VALUES "
	var nameArr = []interface{}{}
	if len(unknownInterests) == 0 {
		return nil
	}
	for nbr, interest := range unknownInterests {
		query += "($" + strconv.Itoa(nbr + 1) + "), "
		nameArr = append(nameArr, interest.Name)
	}
	query = string(query[:len(query) - 2])
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(nameArr...)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) GetInterests() ([]config.Interest, error) {
	var interests = []config.Interest{}
	var interest config.Interest

	stmt, err := conn.db.Prepare("SELECT * FROM interests")
	if err != nil {
		return interests, errors.New(err.Error() + " in preparing")
	}
	rows, err := stmt.Query()
	if err != nil {
		return interests, errors.New(err.Error() + " in executing")
	}
	for rows.Next() {
		err = rows.Scan(&interest.Id, &interest.Name)
		if err != nil {
			return interests, errors.New(err.Error() + " in rows")
		}
		interests = append(interests, interest)
	}
	return interests, nil
}