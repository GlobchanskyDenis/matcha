package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"strconv"
	"testing"
)

var (
	connNotif ConnDB
	notifUser1 User
	notifUser2 User
	notifUser3 User
	notifUser4 User
)

func TestConnect_NotifTest(t *testing.T) {
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
		return
	}
	err = connNotif.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)
}

func TestDropTables_NotifTest(t *testing.T) {
	err := connNotif.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
}

func TestCreateTables_NotifTest(t *testing.T) {
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

func TestCreateUsers_NotifTest(t *testing.T) {
	var err error
	notifUser1, err = connNotif.SetNewUser("notifUser1@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	notifUser2, err = connNotif.SetNewUser("notifUser2@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	notifUser3, err = connNotif.SetNewUser("notifUser3@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	notifUser4, err = connNotif.SetNewUser("notifUser4@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
}

func TestSetNotif_1(t *testing.T) {
	_, err := connNotif.SetNewNotif(notifUser1.Uid, notifUser4.Uid, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_2(t *testing.T) {
	_, err := connNotif.SetNewNotif(notifUser1.Uid, notifUser4.Uid, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_3(t *testing.T) {
	_, err := connNotif.SetNewNotif(notifUser2.Uid, notifUser4.Uid, "test notification. User #2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_4(t *testing.T) {
	_, err := connNotif.SetNewNotif(notifUser3.Uid, notifUser4.Uid, "test notification. User #3")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestGetNotif_1(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(notifUser1.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(notifs) != 2 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR)
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR)
	}
}

func TestGetNotif_2(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(notifUser2.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(notifs) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR)
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR)
	}
}

func TestGetNotif_3(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(notifUser3.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(notifs) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR)
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR)
	}
}

func TestGetNotif_4(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(notifUser4.Uid)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	if len(notifs) != 0 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR)
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR)
	}
}
