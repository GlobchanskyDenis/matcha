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

func (conn ConnDB) DeleteMessage(mid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM messages WHERE mid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(mid)
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

func (conn *ConnDB) GetMessageByMid(mid int) (common.Message, error) {
	var message common.Message
	query := `SELECT
	mid, uidSender, uidReceiver, body, sender_users.fname, sender_users.lname,
	sender_users.src, receiver_users.fname, receiver_users.lname, receiver_users.src
	FROM messages 
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS sender_users ON uidSender=sender_users.uid
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS receiver_users ON uidReceiver=receiver_users.uid
	WHERE mid=$1`
	// "SELECT mid, uidSender, uidReceiver, body FROM messages WHERE mid=$1"
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return message, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(mid)
	if err != nil {
		return message, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if !rows.Next() {
		return message, errors.RecordNotFound
	}
	err = rows.Scan(&message.Mid, &message.UidSender, &message.UidReceiver, &message.Body,
		&message.SenderFname, &message.SenderLname, &message.SenderAvatar,
		&message.ReceiverFname, &message.ReceiverLname, &message.ReceiverAvatar)
	if err != nil {
		return message, errors.DatabaseScanError.AddOriginalError(err)
	}
	return message, nil
}

func (conn ConnDB) GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error) {
	var messages = []common.Message{}
	var message common.Message
	query := `SELECT
	mid, uidSender, uidReceiver, body, sender_users.fname, sender_users.lname,
	sender_users.src, receiver_users.fname, receiver_users.lname, receiver_users.src
	FROM messages 
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS sender_users ON uidSender=sender_users.uid
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS receiver_users ON uidReceiver=receiver_users.uid
	WHERE (uidSender=$1 AND uidReceiver=$2) OR (uidSender=$2 AND uidReceiver=$1)`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uidSender, uidReceiver)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&message.Mid, &message.UidSender, &message.UidReceiver, &message.Body,
			&message.SenderFname, &message.SenderLname, &message.SenderAvatar,
			&message.ReceiverFname, &message.ReceiverLname, &message.ReceiverAvatar)
		if err != nil {
			return messages, errors.DatabaseScanError.AddOriginalError(err)
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (conn ConnDB) GetActiveMessages(uidReceiver int) ([]common.Message, error) {
	var messages = []common.Message{}
	var message common.Message

	query := `SELECT
	mid, uidSender, uidReceiver, body, sender_users.fname, sender_users.lname,
	sender_users.src, receiver_users.fname, receiver_users.lname, receiver_users.src
	FROM messages 
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS sender_users ON uidSender=sender_users.uid
	LEFT JOIN (SELECT users.uid, fname, lname, src FROM users LEFT JOIN photos ON avaID=pid) AS receiver_users ON uidReceiver=receiver_users.uid
	WHERE uidReceiver = $1 AND active = true`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uidReceiver)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&message.Mid, &message.UidSender, &message.UidReceiver, &message.Body,
			&message.SenderFname, &message.SenderLname, &message.SenderAvatar,
			&message.ReceiverFname, &message.ReceiverLname, &message.ReceiverAvatar)
		if err != nil {
			return messages, errors.DatabaseScanError.AddOriginalError(err)
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (conn ConnDB) SetMessageInactive(mid int) error {
	stmt, err := conn.db.Prepare("UPDATE messages SET active=FALSE WHERE mid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(mid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) > 1 {
		return errors.NewArg("Неожиданное количество измененных строк - "+strconv.Itoa(int(nbr64)),
			"Unexpectable amount of changed lines - "+strconv.Itoa(int(nbr64)))
	}
	return nil
}
