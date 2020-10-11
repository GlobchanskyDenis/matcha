package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"testing"
)

func TestDevices(t *testing.T) {
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
	setTestCases := []struct {
		name   string
		uid    int
		device string
	}{
		{
			name:   "set device #1",
			uid:    user1.Uid,
			device: "device_1",
		}, {
			name:   "set device #2",
			uid:    user1.Uid,
			device: "device_2",
		}, {
			name:   "set device #3",
			uid:    user1.Uid,
			device: "device_3",
		}, {
			name:   "set device #4",
			uid:    user2.Uid,
			device: "device_4",
		},
	}

	for _, tc := range setTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := conn.SetNewDevice(tc.uid, tc.device)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
				return
			}
			t_.Log(GREEN_BG + "SUCCESS: device was added to database" + NO_COLOR)
		})
	}

	getTestCases := []struct {
		name      string
		uid       int
		devAmount int
	}{
		{
			name:      "get device with uid=1 and delete them",
			uid:       user1.Uid,
			devAmount: 3,
		}, {
			name:      "get device with uid=2 and delete them",
			uid:       user2.Uid,
			devAmount: 1,
		}, {
			name:      "get device with uid=3 - it should be no devices",
			uid:       user3.Uid,
			devAmount: 0,
		}, {
			name:      "get device with uid=1 - it should be no devices",
			uid:       user1.Uid,
			devAmount: 0,
		},
	}

	for _, tc := range getTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			devices, err := conn.GetDevicesByUid(tc.uid)
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
				return
			}
			if len(devices) != tc.devAmount {
				t_.Errorf(RED_BG+"ERROR: amount of devices is invalid. Expected %d, received %d"+NO_COLOR+"\n", tc.devAmount, len(devices))
				return
			}
			t_.Log(GREEN_BG + "SUCCESS: devices was received from database" + NO_COLOR)
			for _, device := range devices {
				err = conn.DeleteDevice(device.Id)
				if err != nil {
					t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
					return
				}
				t_.Logf(GREEN_BG+"SUCCESS: device with id #%d was removed from database"+NO_COLOR+"\n", device.Id)
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

		conn.Close()
	})
}
