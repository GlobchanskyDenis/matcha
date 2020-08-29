package postgres

import (
	"MatchaServer/common"
	"MatchaServer/config"
	"testing"
)

var connDev ConnDB

func TestInitTables(t *testing.T) {
	print(common.NO_COLOR)
	defer print(common.YELLOW)

	var connDev = New()
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot get config file - " + err.Error() + common.NO_COLOR)
		return
	}
	err = connDev.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + common.NO_COLOR)
		return
	}
	defer connDev.Close()

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
				t_.Errorf(common.RED_BG + "ERROR: " + err.Error() + common.NO_COLOR)
				return
			}
			t_.Log(common.GREEN_BG + "SUCCESS" + common.NO_COLOR)
		})
	}
}

func TestDevice(t *testing.T) {
	print(common.NO_COLOR)
	defer print(common.YELLOW)

	var connDev = New()
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot get config file - " + err.Error() + common.NO_COLOR)
		return
	}
	err = connDev.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + common.NO_COLOR)
		return
	}
	defer connDev.Close()

	setTestCases := []struct {
		name   string
		uid    int
		device string
	}{
		{
			name:   "set device #1",
			uid:    1,
			device: "device_1",
		}, {
			name:   "set device #2",
			uid:    1,
			device: "device_2",
		}, {
			name:   "set device #3",
			uid:    1,
			device: "device_3",
		}, {
			name:   "set device #4",
			uid:    2,
			device: "device_4",
		},
	}

	for _, tc := range setTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := connDev.SetNewDevice(tc.uid, tc.device)
			if err != nil {
				t_.Errorf(common.RED_BG + "ERROR: " + err.Error() + common.NO_COLOR)
				return
			}
			t_.Log(common.GREEN_BG + "SUCCESS: device was added to database" + common.NO_COLOR)
		})
	}

	getTestCases := []struct {
		name      string
		uid       int
		devAmount int
	}{
		{
			name:      "get device with uid=1 and delete them",
			uid:       1,
			devAmount: 3,
		}, {
			name:      "get device with uid=2 and delete them",
			uid:       2,
			devAmount: 1,
		}, {
			name:      "get device with uid=3 - it should be no devices",
			uid:       3,
			devAmount: 0,
		}, {
			name:      "get device with uid=1 - it should be no devices",
			uid:       1,
			devAmount: 0,
		},
	}

	for _, tc := range getTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			devices, err := connDev.GetDevicesByUid(tc.uid)
			if err != nil {
				t_.Errorf(common.RED_BG + "ERROR: " + err.Error() + common.NO_COLOR)
				return
			}
			if len(devices) != tc.devAmount {
				t_.Errorf(common.RED_BG+"ERROR: amount of devices is invalid. Expected %d, received %d"+common.NO_COLOR+"\n", tc.devAmount, len(devices))
				return
			}
			t_.Log(common.GREEN_BG + "SUCCESS: devices was received from database" + common.NO_COLOR)
			for _, device := range devices {
				err = connDev.DeleteDevice(device.Id)
				if err != nil {
					t_.Errorf(common.RED_BG + "ERROR: " + err.Error() + common.NO_COLOR)
					return
				}
				t_.Logf(common.GREEN_BG+"SUCCESS: device with id #%d was removed from database"+common.NO_COLOR+"\n", device.Id)
			}
		})
	}
}
