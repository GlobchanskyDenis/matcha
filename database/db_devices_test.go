package database

import (
	. "MatchaServer/config"
	"testing"
	"strconv"
)

var connDev ConnDB

func TestConnect_DeviceTest(t *testing.T) {
	err := connDev.Connect()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR + "\n")
}

func TestDropTables_DeviceTest(t *testing.T) {
	err := connDev.DropAllTables()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestCreateTables_DeviceTest(t *testing.T) {
	err := connDev.CreateUsersTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connDev.CreateMessageTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connDev.CreateNotifTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
	err = connDev.CreatePhotoTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	err = connDev.CreateDevicesTable()
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: all tables was droped" + NO_COLOR + "\n")
}

func TestSetDevice_1(t *testing.T) {
	err := connDev.SetNewDevice(1, "device_1")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: device was added to database" + NO_COLOR + "\n")
}

func TestSetDevice_2(t *testing.T) {
	err := connDev.SetNewDevice(1, "device_2")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: device was added to database" + NO_COLOR + "\n")
}

func TestSetDevice_3(t *testing.T) {
	err := connDev.SetNewDevice(1, "device_3")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: device was added to database" + NO_COLOR + "\n")
}

func TestSetDevice_4(t *testing.T) {
	err := connDev.SetNewDevice(2, "device_123")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: device was added to database" + NO_COLOR + "\n")
}

func TestGetDevice_1(t *testing.T) {
	devices, err := connDev.GetDevicesByUid(1)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(devices) != 3 {
		t.Errorf(RED_BG + "ERROR: amount of devices is invalid. Expected 3, received " + strconv.Itoa(len(devices)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: devices was received from database" + NO_COLOR + "\n")
	for _, device := range devices {
		err = connDev.DeleteDevice(device.Id)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: device with id #" + strconv.Itoa(device.Id) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetDevice_2(t *testing.T) {
	devices, err := connDev.GetDevicesByUid(2)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(devices) != 1 {
		t.Errorf(RED_BG + "ERROR: amount of devices is invalid. Expected 1, received " + strconv.Itoa(len(devices)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: devices was received from database" + NO_COLOR + "\n")
	for _, device := range devices {
		err = connDev.DeleteDevice(device.Id)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: device with id #" + strconv.Itoa(device.Id) + " was removed from database" + NO_COLOR + "\n")
	}
}

func TestGetDevice_3(t *testing.T) {
	devices, err := connDev.GetDevicesByUid(3)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
		return
	}
	if len(devices) != 0 {
		t.Errorf(RED_BG + "ERROR: amount of devices is invalid. Expected 0, received " + strconv.Itoa(len(devices)) + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS: devices was received from database" + NO_COLOR + "\n")
	for _, device := range devices {
		err = connDev.DeleteDevice(device.Id)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: database returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Log(GREEN_BG + "SUCCESS: device with id #" + strconv.Itoa(device.Id) + " was removed from database" + NO_COLOR + "\n")
	}
}