package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewMessage(uidSender int, uidReceiver int, body string) (int, error) {
	var message common.Message

	message.UidSender = uidSender
	message.UidReceiver = uidReceiver
	message.Body = body

	for key := 1; ; key++ {
		if _, isExists := conn.messages[key]; !isExists {
			message.Mid = key
			break
		}
	}

	conn.messages[message.Mid] = message
	return message.Mid, nil
}

func (conn *ConnFake) DeleteMessage(mid int) error {
	_, isExists := conn.messages[mid]
	if !isExists {
		return errors.RecordNotFound
	}
	delete(conn.messages, mid)
	return nil
}

func (conn ConnFake) GetMessageByMid(mid int) (common.Message, error) {
	var message common.Message

	for _, message = range conn.messages {
		if message.Mid == mid {
			return message, nil
		}
	}
	return message, errors.RecordNotFound
}

func (conn ConnFake) GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error) {
	var messages = []common.Message{}
	var message common.Message

	for _, message = range conn.messages {
		if (message.UidSender == uidSender && message.UidReceiver == uidReceiver) ||
			(message.UidSender == uidReceiver && message.UidReceiver == uidSender) {
			messages = append(messages, message)
		}
	}
	return messages, nil
}

func (conn ConnFake) GetActiveMessages(uidReceiver int) ([]common.Message, error) {
	return nil, nil
}

func (conn ConnFake) SetMessageInactive(mid int) error {
	return nil
}
