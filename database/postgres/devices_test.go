package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"testing"
)

var (
	connDev ConnDB
	deviceUser1 User
	deviceUser2 User
	deviceUser3 User
)

func TestConnect_DeviceTest(t *testing.T) {
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
		return
	}
	err = connDev.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)
}

func TestInitTables(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	TestCases := []struct {
		name     string
		function func() error
	}{
		{
			name:     "Drop all tables",
			function: connDev.DropAllTables,
		}, {
			name:     "Create users table",
			function: connDev.CreateUsersTable,
		}, {
			name:     "Create messages table",
			function: connDev.CreateMessagesTable,
		}, {
			name:     "Create notifications table",
			function: connDev.CreateNotifsTable,
		}, {
			name:     "Create photos table",
			function: connDev.CreatePhotosTable,
		}, {
			name:     "Create devices table",
			function: connDev.CreateDevicesTable,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := tc.function()
			if err != nil {
				t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
				return
			}
			t_.Log(GREEN_BG + "SUCCESS" + NO_COLOR)
		})
	}
}

func TestCreateUsers_DeviceTest(t *testing.T) {
	var err error
	deviceUser1, err = connDev.SetNewUser("deviceUser1@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	deviceUser2, err = connDev.SetNewUser("deviceUser2@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
	deviceUser3, err = connDev.SetNewUser("deviceUser3@gmail.com", "qwerty")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
		return
	}
}

func TestDevice(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	setTestCases := []struct {
		name   string
		uid    int
		device string
	}{
		{
			name:   "set device #1",
			uid:    deviceUser1.Uid,
			device: "device_1",
		}, {
			name:   "set device #2",
			uid:    deviceUser1.Uid,
			device: "device_2",
		}, {
			name:   "set device #3",
			uid:    deviceUser1.Uid,
			device: "device_3",
		}, {
			name:   "set device #4",
			uid:    deviceUser2.Uid,
			device: "device_4",
		},
	}

	for _, tc := range setTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := connDev.SetNewDevice(tc.uid, tc.device)
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
			uid:       deviceUser1.Uid,
			devAmount: 3,
		}, {
			name:      "get device with uid=2 and delete them",
			uid:       deviceUser2.Uid,
			devAmount: 1,
		}, {
			name:      "get device with uid=3 - it should be no devices",
			uid:       deviceUser3.Uid,
			devAmount: 0,
		}, {
			name:      "get device with uid=1 - it should be no devices",
			uid:       deviceUser1.Uid,
			devAmount: 0,
		},
	}

	for _, tc := range getTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			devices, err := connDev.GetDevicesByUid(tc.uid)
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
				err = connDev.DeleteDevice(device.Id)
				if err != nil {
					t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
					return
				}
				t_.Logf(GREEN_BG+"SUCCESS: device with id #%d was removed from database"+NO_COLOR+"\n", device.Id)
			}
		})
	}
}
