package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
)

func (conn ConnDB) SetNewMessage(uidSender int, uidReceiver int, body string) (int, error) {
	var mid int
	stmt, err := conn.db.Prepare("INSERT INTO messages (uidSender, uidReceiver, body) VALUES ($1, $2, $3) RETURNING mid")
	if err != nil {
		return mid, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(uidSender, uidReceiver, body).Scan(&mid)
	if err != nil {
		return mid, errors.DatabaseQueryError.AddOriginalError(err)
	}
	return mid, nil
}

func (conn ConnDB) DeleteMessage(nid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM messages WHERE mid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(nid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Неожиданное количество измененных строк - "+strconv.Itoa(int(nbr64)),
			"Unexpectable amount of changed lines - "+strconv.Itoa(int(nbr64)))
	}
	return nil
}

func (conn ConnDB) GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error) {
	var messages = []common.Message{}
	var message common.Message

	stmt, err := conn.db.Prepare("SELECT * FROM messages WHERE " +
		"(uidSender=$1 AND uidReceiver=$2) OR (uidSender=$2 AND uidReceiver=$1)")
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uidSender, uidReceiver)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(message.Mid), &(message.UidSender), &(message.UidReceiver), &(message.Body))
		if err != nil {
			return messages, errors.DatabaseScanError.AddOriginalError(err)
		}
		messages = append(messages, message)
	}
	return messages, nil
}
