package postgres

import (
	"MatchaServer/config"
	"errors"
)

func (conn ConnDB) SetNewMessage(uidSender int, uidReceiver int, body string) (int, error) {
	var mid int
	stmt, err := conn.db.Prepare("INSERT INTO message (uidSender, uidReceiver, body) VALUES ($1, $2, $3) RETURNING mid")
	if err != nil {
		return mid, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	err = stmt.QueryRow(uidSender, uidReceiver, body).Scan(&mid)
	if err != nil {
		return mid, errors.New(err.Error() + " in executing")
	}
	return mid, nil
}

func (conn ConnDB) DeleteMessage(nid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM message WHERE mid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(nid)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) GetMessagesFromChat(uidSender int, uidReceiver int) ([]config.Message, error) {
	var messages = []config.Message{}
	var message config.Message

	stmt, err := conn.db.Prepare("SELECT * FROM message WHERE " +
		"(uidSender=$1 AND uidReceiver=$2) OR (uidSender=$2 AND uidReceiver=$1)")
	if err != nil {
		return messages, errors.New(err.Error() + " in preparing")
	}
	rows, err := stmt.Query(uidSender, uidReceiver)
	if err != nil {
		return messages, errors.New(err.Error() + " in executing")
	}
	for rows.Next() {
		err = rows.Scan(&(message.Mid), &(message.UidSender), &(message.UidReceiver), &(message.Body))
		if err != nil {
			return messages, errors.New(err.Error() + " in rows")
		}
		messages = append(messages, message)
	}
	return messages, nil
}