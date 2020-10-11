package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errors"
	"testing"
)

func TestPhotos(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		conn  ConnDB
		user1 User
		user2 User
		pid   int
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
		t_.Log(GREEN_BG + "SUCCESS: test users was created" + NO_COLOR)
	})

	/*
	**	Test cases - main part of testing
	**	Set photos
	 */
	testCasesPut := []struct {
		name string
		uid  int
		body string
	}{
		{
			name: "create photo #1",
			uid:  user1.Uid,
			body: "photo body",
		}, {
			name: "create photo #2",
			uid:  user1.Uid,
			body: "photo body",
		}, {
			name: "create photo #3",
			uid:  user2.Uid,
			body: "photo body",
		},
	}

	for _, tc := range testCasesPut {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			pid, err = conn.SetNewPhoto(tc.uid, tc.body)
			if err != nil {
				t.Errorf(RED_BG + "ERROR: cannot set photo - " + err.Error() + NO_COLOR)
				return
			}
			t.Logf(GREEN_BG+"SUCCESS: photo #%d was added to database"+NO_COLOR, pid)
		})
	}

	/*
	**	Valid test of GetPhotoByPid function
	 */
	t.Run("valid get photo by pid", func(t_ *testing.T) {
		_, err := conn.GetPhotoByPid(pid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
			return
		}
		t_.Log(GREEN_BG + "Success" + NO_COLOR)
	})

	/*
	**	Test cases - get notifications
	 */
	testCasesGet := []struct {
		name        string
		uid         int
		expectedLen int
	}{
		{
			name:        "get photos of user1",
			uid:         user1.Uid,
			expectedLen: 2,
		}, {
			name:        "get photos of user2",
			uid:         user2.Uid,
			expectedLen: 1,
		}, {
			name:        "get photos of user1",
			uid:         user1.Uid,
			expectedLen: 0,
		},
	}

	for _, tc := range testCasesGet {
		t.Run(tc.name, func(t_ *testing.T) {
			photos, err := conn.GetPhotosByUid(tc.uid)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: cannot get photos - " + err.Error() + NO_COLOR)
				return
			}
			if len(photos) != tc.expectedLen {
				t_.Errorf(RED_BG+"ERROR: amount of photos is invalid. Expected %d, received %d"+NO_COLOR, tc.expectedLen, len(photos))
				return
			}
			t_.Logf(GREEN_BG + "SUCCESS: photos was received from database" + NO_COLOR)
			for _, photo := range photos {
				err = conn.DeletePhoto(photo.Pid)
				if err != nil {
					t_.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR)
					return
				}
				t_.Logf(GREEN_BG+"SUCCESS: photo with id #%d was removed from database"+NO_COLOR, photo.Pid)
			}
		})
	}

	/*
	**	Invalid test of GetPhotoByPid function (all photos are already deleted)
	 */
	t.Run("invalid get photo by pid", func(t_ *testing.T) {
		_, err := conn.GetPhotoByPid(pid)
		if err != nil && errors.RecordNotFound.IsOverlapWithError(err) {
			t_.Log(GREEN_BG + "Success: record not found as it expected" + NO_COLOR)
		} else if err != nil {
			t_.Errorf(RED_BG + "Error: " + err.Error() + NO_COLOR)
		} else {
			t_.Errorf(RED_BG + "Error: no errors returned - but it should be..." + NO_COLOR)
		}
	})

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
		t_.Log(GREEN_BG + "SUCCESS: all test users was deleted" + NO_COLOR)

		conn.Close()
	})
}
