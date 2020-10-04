package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
)

func (conn ConnDB) SetNewMessage(uidSender int, uidReceiver int, body string) (int, error) {
	var mid int
	stmt, err := conn.db.Prepare("INSERT INTO messages (uidSender, uidReceiver, body) VALUES ($1, $2, $3) RETURNING mid")
	if err != nil {
		return mid, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(uidSender, uidReceiver, body).Scan(&mid)
	if err != nil {
		return mid, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	return mid, nil
}

func (conn ConnDB) DeleteMessage(nid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM messages WHERE mid=$1")
	if err != nil {
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nid)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error) {
	var messages = []common.Message{}
	var message common.Message

	stmt, err := conn.db.Prepare("SELECT * FROM messages WHERE " +
		"(uidSender=$1 AND uidReceiver=$2) OR (uidSender=$2 AND uidReceiver=$1)")
	if err != nil {
		return nil, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uidSender, uidReceiver)
	if err != nil {
		return nil, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(message.Mid), &(message.UidSender), &(message.UidReceiver), &(message.Body))
		if err != nil {
			return messages, errDef.DatabaseScanError.AddOriginalError(err)
		}
		messages = append(messages, message)
	}
	return messages, nil
}
