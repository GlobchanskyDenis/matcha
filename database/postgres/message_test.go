package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"strconv"
	"testing"
)

var (
	connMes ConnDB
	messageUser1 User
	messageUser2 User
	messageUser3 User
)

func TestConnect_MessageTest(t *testing.T) {
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
		return
	}
	err = connMes.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)
}

func TestDropTables_MessageTest(t *testing.T) {
	err := connMes.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
}

func TestCreateTables_MessageTest(t *testing.T) {
	err := connMes.CreateUsersTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: users table was created" + NO_COLOR)
	err = connMes.CreateMessagesTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: messages table was created" + NO_COLOR)
	err = connMes.CreateNotifsTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifs table was created" + NO_COLOR)
	err = connMes.CreatePhotosTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: photos table was created" + NO_COLOR)
}

func TestCreateUsers_MessageTest(t *testing.T) {
	var err error
	messageUser1, err = connMes.SetNewUser("messageUser1@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	messageUser2, err = connMes.SetNewUser("messageUser2@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	messageUser3, err = connMes.SetNewUser("messageUser3@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
}

func TestSetMessage_1(t *testing.T) {
	_, err := connMes.SetNewMessage(messageUser1.Uid, messageUser2.Uid, "transmit message from 1 to 2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR)
}

func TestSetMessage_2(t *testing.T) {
	_, err := connMes.SetNewMessage(messageUser2.Uid, messageUser1.Uid, "transmit message from 2 to 1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR)
}

func TestSetMessage_3(t *testing.T) {
	_, err := connMes.SetNewMessage(messageUser2.Uid, messageUser2.Uid, "transmit message from 2 to 2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR)
}

func TestSetMessage_4(t *testing.T) {
	_, err := connMes.SetNewMessage(messageUser3.Uid, messageUser1.Uid, "transmit message from 3 to 1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR)
}

func TestGetMessage_1(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(messageUser1.Uid, messageUser2.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(messages) != 2 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 2, received " + strconv.Itoa(len(messages)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR)
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR)
	}
}

func TestGetMessage_2(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(messageUser2.Uid, messageUser2.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(messages) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 1, received " + strconv.Itoa(len(messages)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR)
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR)
	}
}

func TestGetMessage_3(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(messageUser3.Uid, messageUser1.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(messages) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 1, received " + strconv.Itoa(len(messages)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR)
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR)
	}
}
