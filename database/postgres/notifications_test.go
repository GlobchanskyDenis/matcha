package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	// "strconv"
	"testing"
)

func TestNotif(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		conn  ConnDB
		user1 User
		user2 User
		user3 User
		user4 User
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
		user4, err = conn.SetNewUser("user4@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		t_.Log(GREEN_BG + "SUCCESS: test users was created" + NO_COLOR)
	})

	/*
	**	Test cases - main part of testing
	**	Set notifications
	 */
	testCasesPut := []struct {
		name         string
		uid1         int
		uid2         int
		notification string
	}{
		{
			name:         "set notification #1",
			uid1:         user1.Uid,
			uid2:         user4.Uid,
			notification: "test notification",
		}, {
			name:         "set notification #2",
			uid1:         user1.Uid,
			uid2:         user4.Uid,
			notification: "test notification",
		}, {
			name:         "set notification #3",
			uid1:         user2.Uid,
			uid2:         user4.Uid,
			notification: "test notification",
		}, {
			name:         "set notification #4",
			uid1:         user3.Uid,
			uid2:         user4.Uid,
			notification: "test notification",
		},
	}

	for _, tc := range testCasesPut {
		t.Run(tc.name, func(t *testing.T) {
			nid, err := conn.SetNewNotif(tc.uid1, tc.uid2, tc.notification)
			if err != nil {
				t.Errorf(RED_BG + "ERROR: cannot set notification - " + err.Error() + NO_COLOR)
				return
			}
			t.Logf(GREEN_BG+"SUCCESS: notification #%d was added to database"+NO_COLOR, nid)
		})
	}

	/*
	**	Test cases - get notifications
	 */
	testCasesGet := []struct {
		name        string
		uid         int
		expectedLen int
	}{
		{
			name:        "get notification of user #1",
			uid:         user1.Uid,
			expectedLen: 2,
		}, {
			name:        "get notification of user #2",
			uid:         user2.Uid,
			expectedLen: 1,
		}, {
			name:        "get notification of user #3",
			uid:         user3.Uid,
			expectedLen: 1,
		}, {
			name:        "get notification of user #4",
			uid:         user4.Uid,
			expectedLen: 0,
		},
	}

	for _, tc := range testCasesGet {
		t.Run(tc.name, func(t_ *testing.T) {
			notifs, err := conn.GetNotifByUidReceiver(tc.uid)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: cannot get notification - " + err.Error() + NO_COLOR)
				return
			}
			if len(notifs) != tc.expectedLen {
				t_.Errorf(RED_BG+"ERROR: amount of notification is invalid. Expected %d, received %d"+NO_COLOR, tc.expectedLen, len(notifs))
				return
			}
			t_.Logf(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR)
			for _, notif := range notifs {
				err = conn.DeleteNotif(notif.Nid)
				if err != nil {
					t_.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
					return
				}
				t_.Logf(GREEN_BG+"SUCCESS: notification with nid #%d was removed from database"+NO_COLOR, notif.Nid)
			}
		})
	}

	/*
	**	Delete test users and close connection
	 */
	t.Run("Delete test users and close connection", func(t_ *testing.T) {
		err := conn.DeleteUser(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
		}
		err = conn.DeleteUser(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
		}
		err = conn.DeleteUser(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
		}
		t_.Log(GREEN_BG + "SUCCESS: all test users was deleted" + NO_COLOR)
		err = conn.DeleteUser(user4.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
		}
		t_.Log(GREEN_BG + "SUCCESS: all test users was deleted" + NO_COLOR)

		conn.Close()
	})
}
