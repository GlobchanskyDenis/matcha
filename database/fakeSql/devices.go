package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnFake) SetNewDevice(uid int, deviceName string) error {
	var device config.Device

	device.Uid = uid
	device.Device = deviceName

	for key := 1; ; key++ {
		if _, isExists := conn.devices[key]; !isExists {
			device.Id = key
			break
		}
	}

	conn.devices[device.Id] = device
	return nil
}

func (conn ConnFake) DeleteDevice(id int) error {
	delete(conn.devices, id)
	return nil
}

func (conn ConnFake) GetDevicesByUid(uid int) ([]config.Device, error) {
	var devices = []config.Device{}
	var device config.Device

	for _, device = range conn.devices {
		if device.Uid == uid {
			devices = append(devices, device)
		}
	}
	return devices, nil
}
