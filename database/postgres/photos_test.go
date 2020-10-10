package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errors"
	"testing"
)

var (
	connPhoto ConnDB
	photoUser1 User
	photoUser2 User
)

func TestConnect_PhotoTest(t *testing.T) {
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
		return
	}
	err = connPhoto.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)
}

func TestDropTables_PhotoTest(t *testing.T) {
	err := connPhoto.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR)
}

func TestCreateTables_PhotoTest(t *testing.T) {
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

func TestCreateUsers_PhotoTest(t *testing.T) {
	var err error
	photoUser1, err = connPhoto.SetNewUser("photoUser1@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	photoUser2, err = connPhoto.SetNewUser("photoUser2@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
}

func TestPhotos(t *testing.T) {
	print(NO_COLOR)
	var (
		body1 = "qwerty" //[]byte("qwerty")
		body2 = "ytrewq" //[]byte("ytrewq")
		body3 = "asd"    //[]byte("asd")
	)

	t.Run("Create Photo #1", func(t_ *testing.T) {
		_, err := connPhoto.SetNewPhoto(photoUser1.Uid, body1)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Create Photo #2", func(t_ *testing.T) {
		_, err := connPhoto.SetNewPhoto(photoUser1.Uid, body2)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Create Photo #3", func(t_ *testing.T) {
		_, err := connPhoto.SetNewPhoto(photoUser2.Uid, body3)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Get photos by photoUser1.Uid #1", func(t_ *testing.T) {
		photos, err := connPhoto.GetPhotosByUid(photoUser1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		if len(photos) != 2 {
			t_.Errorf(RED_BG+"Error: expected 2 photos, got %d"+NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Get photos by photoUser2.Uid #1", func(t_ *testing.T) {
		photos, err := connPhoto.GetPhotosByUid(photoUser2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		if len(photos) != 1 {
			t_.Errorf(RED_BG+"Error: expected 1 photos, got %d"+NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Delete photo", func(t_ *testing.T) {
		var pid int
		photos, err := connPhoto.GetPhotosByUid(photoUser1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = connPhoto.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Get photos by photoUser1.Uid #2", func(t_ *testing.T) {
		photos, err := connPhoto.GetPhotosByUid(photoUser1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		if len(photos) != 1 {
			t_.Errorf(RED_BG+"Error: expected 1 photos, got %d"+NO_COLOR+"\n", len(photos))
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	t.Run("Invalid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := connPhoto.GetPhotosByUid(photoUser1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = connPhoto.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		_, err = connPhoto.GetPhotoByPid(pid)
		if errors.RecordNotFound.IsOverlapWithError(err) {
			t_.Log(GREEN_BG + "Success: there is `RecordNotFoundError` as it expected" + NO_COLOR)
		} else if err != nil {
			t_.Errorf(RED_BG + "Error: it should be RecordNotFound but it dont - " + err.Error() + NO_COLOR)
			return
		} else {
			t_.Errorf(RED_BG + "Error: it should be RecordNotFound but there is no error at all" + NO_COLOR)
			return
		}
	})
	t.Run("Valid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := connPhoto.GetPhotosByUid(photoUser2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		_, err = connPhoto.GetPhotoByPid(pid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		err = connPhoto.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})
	print(YELLOW)
}
