package postgres

import (
	"MatchaServer/config"
	"MatchaServer/errDef"
	"testing"
)

func TestPhotos(t *testing.T) {
	print(config.NO_COLOR)
	var (
		repo = New()
		uid1 = 1000
		uid2 = 1002
		body1 = []byte("qwerty")
		body2 = []byte("ytrewq")
		body3 = []byte("asd")
	)
	_ = repo.Connect()

	t.Run("Create Photo #1", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid1, body1)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Create Photo #2", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid1, body2)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Create Photo #3", func(t_ *testing.T) {
		_, err := repo.SetNewPhoto(uid2, body3)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Get photos by uid1 #1", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		if len(photos) != 2 {
			t_.Errorf(config.RED_BG + "Error: expected 2 photos, got %d" + config.NO_COLOR + "\n", len(photos))
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Get photos by uid2 #1", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid2)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		if len(photos) != 1 {
			t_.Errorf(config.RED_BG + "Error: expected 1 photos, got %d" + config.NO_COLOR + "\n", len(photos))
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Delete photo", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Get photos by uid1 #2", func(t_ *testing.T) {
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		if len(photos) != 1 {
			t_.Errorf(config.RED_BG + "Error: expected 1 photos, got %d" + config.NO_COLOR + "\n", len(photos))
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	t.Run("Invalid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid1)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		_, err = repo.GetPhotoByPid(pid)
		if errDef.IsRecordNotFoundError(err) {
			t_.Log(config.GREEN_BG + "Success: there is `RecordNotFoundError` as it expected" + config.NO_COLOR + "\n")
		} else if err != nil {
			t_.Errorf(config.RED_BG + "Error: it should be RecordNotFound but it dont - " + err.Error() + config.NO_COLOR + "\n")
			return
		} else {
			t_.Errorf(config.RED_BG + "Error: it should be RecordNotFound but there is no error at all" + config.NO_COLOR + "\n")
			return
		}
	})
	t.Run("Valid GetPhotoByPid", func(t_ *testing.T) {
		var pid int
		photos, err := repo.GetPhotosByUid(uid2)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		for _, photo := range photos {
			pid = photo.Pid
		}
		_, err = repo.GetPhotoByPid(pid)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		err = repo.DeletePhoto(pid)
		if err != nil {
			t_.Errorf(config.RED_BG + "Error: " + err.Error() + config.NO_COLOR + "\n")
			return
		}
		t_.Log(config.GREEN_BG + "Success" + config.NO_COLOR + "\n")
	})
	print(config.YELLOW)
}