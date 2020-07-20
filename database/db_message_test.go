package database

import (
	. "MatchaServer/config"
	"testing"
	"strconv"
)

var connMes ConnDB

func TestConnect_MessageTest(t *testing.T) {
	err := connMes.Connect()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR + "\n")
}

func TestDropTables_MessageTest(t *testing.T) {
	err := connMes.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestCreateTables_MessageTest(t *testing.T) {
	err := connMes.CreateUsersTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connMes.CreateMessageTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connMes.CreateNotifTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connMes.CreatePhotoTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestSetMessage_1(t *testing.T) {
	err := connMes.SetNewMessage(1, 2, "transmit message from 1 to 2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR + "\n")
}

func TestSetMessage_2(t *testing.T) {
	err := connMes.SetNewMessage(2, 1, "transmit message from 2 to 1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR + "\n")
}

func TestSetMessage_3(t *testing.T) {
	err := connMes.SetNewMessage(2, 2, "transmit message from 2 to 2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR + "\n")
}

func TestSetMessage_4(t *testing.T) {
	err := connMes.SetNewMessage(3, 1, "transmit message from 3 to 1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was added to database" + NO_COLOR + "\n")
}

func TestGetMessage_1(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(1, 2)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(messages) != 2 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 2, received " + strconv.Itoa(len(messages)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR + "\n")
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetMessage_2(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(2, 2)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(messages) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 1, received " + strconv.Itoa(len(messages)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR + "\n")
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetMessage_3(t *testing.T) {
	messages, err := connMes.GetMessagesFromChat(3, 1)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(messages) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of messages is invalid. Expected 1, received " + strconv.Itoa(len(messages)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: message was received from database" + NO_COLOR + "\n")
	for _, message := range messages {
		err = connMes.DeleteMessage(message.Mid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: message with mid #" + strconv.Itoa(message.Mid) + " was removed from database" + NO_COLOR + "\n")
	}
}