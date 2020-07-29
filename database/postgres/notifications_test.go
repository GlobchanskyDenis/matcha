package postgres

import (
	. "MatchaServer/config"
	"strconv"
	"testing"
)

var connNotif ConnDB

func TestConnect_NotifTest(t *testing.T) {
	err := connNotif.Connect()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR + "\n")
}

func TestDropTables_NotifTest(t *testing.T) {
	err := connNotif.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestCreateTables_NotifTest(t *testing.T) {
	err := connNotif.CreateUsersTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connNotif.CreateMessageTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connNotif.CreateNotifTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connNotif.CreatePhotoTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestSetNotif_1(t *testing.T) {
	_, err := connNotif.SetNewNotif(1, 4, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR + "\n")
}

func TestSetNotif_2(t *testing.T) {
	_, err := connNotif.SetNewNotif(1, 4, "test notification. User #1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR + "\n")
}

func TestSetNotif_3(t *testing.T) {
	_, err := connNotif.SetNewNotif(2, 4, "test notification. User #2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR + "\n")
}

func TestSetNotif_4(t *testing.T) {
	_, err := connNotif.SetNewNotif(3, 4, "test notification. User #3")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notification was added to database" + NO_COLOR + "\n")
}

func TestGetNotif_1(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(1)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(notifs) != 2 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR + "\n")
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetNotif_2(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(2)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(notifs) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR + "\n")
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetNotif_3(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(3)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(notifs) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR + "\n")
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetNotif_4(t *testing.T) {
	notifs, err := connNotif.GetNotifByUidReceiver(4)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(notifs) != 0 {
		t.Errorf(RED_BG + "ERROR: amount of notifications is invalid. Expected 2, received " + strconv.Itoa(len(notifs)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: notifications was received from database" + NO_COLOR + "\n")
	for _, notif := range notifs {
		err = connNotif.DeleteNotif(notif.Nid)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: notification with nid #" + strconv.Itoa(notif.Nid) + " was removed from database" + NO_COLOR + "\n")
	}
}
