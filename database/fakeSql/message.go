package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnFake) SetNewMessage(uidSender int, uidReceiver int, body string) (int, error) {
	var message config.Message

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

func (conn ConnFake) DeleteMessage(nid int) error {
	delete(conn.messages, nid)
	return nil
}

func (conn ConnFake) GetMessagesFromChat(uidSender int, uidReceiver int) ([]config.Message, error) {
	var messages = []config.Message{}
	var message config.Message

	for _, message = range conn.messages {
		if (message.UidSender == uidSender && message.UidReceiver == uidReceiver) &&
			(message.UidSender == uidReceiver && message.UidReceiver == uidSender) {
			messages = append(messages, message)
		}
	}
	return messages, nil
}
