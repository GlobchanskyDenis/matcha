package postgres

import (
	"MatchaServer/config"
	"testing"
)

var connDev ConnDB

func TestInitTables(t * testing.T) {
	print(config.NO_COLOR)
	defer print(config.YELLOW)

	var connDev = New()
	err := connDev.Connect()
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + config.NO_COLOR + "\n")
		return
	}
	defer connDev.Close()

	TestCases := []struct {
		name        string
		function	func() error
	}{
		{
			name: "Drop all tables",
			function: connDev.DropAllTables,
		}, {
			name: "Create users table",
			function: connDev.CreateUsersTable,
		}, {
			name: "Create messages table",
			function: connDev.CreateMessagesTable,
		}, {
			name: "Create notifications table",
			function: connDev.CreateNotifsTable,
		}, {
			name: "Create photos table",
			function: connDev.CreatePhotosTable,
		}, {
			name: "Create devices table",
			function: connDev.CreateDevicesTable,
		},
	}

	for _, tc := range TestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := tc.function()
			if err != nil {
				t_.Errorf(config.RED_BG + "ERROR: " + err.Error() + config.NO_COLOR + "\n")
				return
			}
			t_.Log(config.GREEN_BG + "SUCCESS" + config.NO_COLOR + "\n")
		})
	}
}

func TestDevice(t *testing.T) {
	print(config.NO_COLOR)
	defer print(config.YELLOW)

	var connDev = New()
	err := connDev.Connect()
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + config.NO_COLOR + "\n")
		return
	}
	defer connDev.Close()

	setTestCases := []struct {
		name           string
		uid				int
		device			string
	}{
		{
			name: "set device #1",
			uid:		1,
			device:	"device_1",
		}, {
			name: "set device #2",
			uid:		1,
			device:	"device_2",
		}, {
			name: "set device #3",
			uid:		1,
			device:	"device_3",
		}, {
			name: "set device #4",
			uid:		2,
			device:	"device_4",
		},
	}

	for _, tc := range setTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			err := connDev.SetNewDevice(tc.uid, tc.device)
			if err != nil {
				t_.Errorf(config.RED_BG + "ERROR: " + err.Error() + config.NO_COLOR + "\n")
				return
			}
			t_.Log(config.GREEN_BG + "SUCCESS: device was added to database" + config.NO_COLOR + "\n")
		})
	}

	getTestCases := []struct {
		name        string
		uid			int
		devAmount	int
	}{
		{
			name: "get device with uid=1 and delete them",
			uid:		1,
			devAmount: 3,
		}, {
			name: "get device with uid=2 and delete them",
			uid:		2,
			devAmount: 1,
		}, {
			name: "get device with uid=3 - it should be no devices",
			uid:		3,
			devAmount: 0,
		}, {
			name: "get device with uid=1 - it should be no devices",
			uid:		1,
			devAmount: 0,
		},
	}

	for _, tc := range getTestCases {
		t.Run(tc.name, func(t_ *testing.T) {
			devices, err := connDev.GetDevicesByUid(tc.uid)
			if err != nil {
				t_.Errorf(config.RED_BG + "ERROR: " + err.Error() + config.NO_COLOR + "\n")
				return
			}
			if len(devices) != tc.devAmount {
				t_.Errorf(config.RED_BG + "ERROR: amount of devices is invalid. Expected %d, received %d" + config.NO_COLOR + "\n", tc.devAmount, len(devices))
				return
			}
			t_.Log(config.GREEN_BG + "SUCCESS: devices was received from database" + config.NO_COLOR + "\n")
			for _, device := range devices {
				err = connDev.DeleteDevice(device.Id)
				if err != nil {
					t_.Errorf(config.RED_BG + "ERROR: " + err.Error() + config.NO_COLOR + "\n")
					return
				}
				t_.Logf(config.GREEN_BG + "SUCCESS: device with id #%d was removed from database" + config.NO_COLOR + "\n", device.Id)
			}
		})
	}
}
