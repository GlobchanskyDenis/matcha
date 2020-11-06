package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"strconv"
	"testing"
)

func TestMessages(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		conn  ConnDB
		user1 User
		user2 User
		user3 User
	)

	/*
	**	Initialize connection and test users
	 */
	t.Run("Initialize", func(t_ *testing.T) {
		conf, err := config.Create("../../config/")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		err = conn.Connect(&conf.Sql)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		t_.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)

		user1, err = conn.SetNewUser("user1@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user2, err = conn.SetNewUser("user2@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user3, err = conn.SetNewUser("user3@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		t_.Log(GREEN_BG + "SUCCESS: test users was created" + NO_COLOR)
	})

	/*
	**	Test cases - main part of testing
	 */
	testCasesPut := []struct {
		name    string
		uid1    int
		uid2    int
		message string
	}{
		{
			name:    "set message #1",
			uid1:    user1.Uid,
			uid2:    user2.Uid,
			message: "transmit message from 1 to 2",
		}, {
			name:    "set message #2",
			uid1:    user2.Uid,
			uid2:    user1.Uid,
			message: "transmit message from 2 to 1",
		}, {
			name:    "set message #3",
			uid1:    user2.Uid,
			uid2:    user2.Uid,
			message: "transmit message from 2 to 2",
		}, {
			name:    "set message #4",
			uid1:    user3.Uid,
			uid2:    user1.Uid,
			message: "transmit message from 3 to 1",
		},
	}

	for nbr, tc := range testCasesPut {
		t.Run(tc.name, func(t_ *testing.T) {
			mid, err := conn.SetNewMessage(tc.uid1, tc.uid2, tc.message)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: cannot send message - " + err.Error() + NO_COLOR)
				return
			}
			t_.Log(GREEN_BG + "SUCCESS #" + strconv.Itoa(nbr) + ": message #" + strconv.Itoa(mid) + " was added to database" + NO_COLOR)
		})
	}

	testCasesGet := []struct {
		name        string
		uid1        int
		uid2        int
		expectedLen int
	}{
		{
			name:        "get messages from user#1 && user#2",
			uid1:        user1.Uid,
			uid2:        user2.Uid,
			expectedLen: 2,
		}, {
			name:        "get messages from user#2 && user#2",
			uid1:        user2.Uid,
			uid2:        user2.Uid,
			expectedLen: 1,
		}, {
			name:        "get messages from user#1 && user#3",
			uid1:        user1.Uid,
			uid2:        user3.Uid,
			expectedLen: 1,
		},
	}

	for nbr, tc := range testCasesGet {
		t.Run(tc.name, func(t_ *testing.T) {
			messages, err := conn.GetMessagesFromChat(tc.uid1, tc.uid2)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
				return
			}
			if len(messages) != tc.expectedLen {
				t_.Errorf(RED_BG+"ERROR: amount of messages is invalid. Expected %d, received %d"+NO_COLOR, tc.expectedLen, len(messages))
				return
			}
			t_.Logf(GREEN_BG+"SUCCESS #%d: message was received from database"+NO_COLOR, nbr)
			for _, message := range messages {
				err = conn.DeleteMessage(message.Mid)
				if err != nil {
					t_.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
					return
				}
				t_.Logf(GREEN_BG+"SUCCESS #%d: message with mid #%d was removed from database"+NO_COLOR, nbr, message.Mid)
			}
		})
	}

	/*
	**	Delete test users and close connection
	 */
	t.Run("Delete test users and close connection", func(t_ *testing.T) {
		var wasError bool
		err := conn.DropUserIgnores(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}

		err = conn.DropUserIgnores(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}

		err = conn.DropUserIgnores(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}
		if !wasError {
			t_.Log(GREEN_BG + "SUCCESS: all test users was deleted" + NO_COLOR)
		}
		conn.Close()
	})
}
