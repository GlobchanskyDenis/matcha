package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"strconv"
	"testing"
)

var connNotif ConnDB

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
	err := connNotif.CreateUsersTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
	err = connNotif.CreateMessagesTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
	err = connNotif.CreateNotifsTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
	err = connNotif.CreatePhotosTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
}

func TestSetNotif_1(t *testing.T) {
	_, err := connNotif.SetNewNotif(1, 4, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_2(t *testing.T) {
	_, err := connNotif.SetNewNotif(1, 4, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_3(t *testing.T) {
	_, err := connNotif.SetNewNotif(2, 4, "test notification. User #2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestSetNotif_4(t *testing.T) {
	_, err := connNotif.SetNewNotif(3, 4, "test notification. User #3")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR)
}

func TestGetNotif_1(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(1)
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
	notifs, err := connNotif.GetNotifByUidReceiver(2)
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
	notifs, err := connNotif.GetNotifByUidReceiver(3)
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
	notifs, err := connNotif.GetNotifByUidReceiver(4)
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
