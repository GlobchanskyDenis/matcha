package postgres

import (
	"MatchaServer/config"
	"errors"
)

func (conn ConnDB) SetNewDevice(uid int, device string) error {
	stmt, err := conn.db.Prepare("INSERT INTO devices (uid, device) VALUES ($1, $2)")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, device)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) DeleteDevice(id int) error {
	stmt, err := conn.db.Prepare("DELETE FROM devices WHERE id=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) GetDevicesByUid(uid int) ([]config.Device, error) {
	var devices = []config.Device{}
	var device config.Device

	stmt, err := conn.db.Prepare("SELECT * FROM devices WHERE uid=$1")
	if err != nil {
		return devices, errors.New(err.Error() + " in preparing")
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return devices, errors.New(err.Error() + " in executing")
	}
	for rows.Next() {
		err = rows.Scan(&(device.Id), &(device.Uid), &(device.Device))
		if err != nil {
			return devices, errors.New(err.Error() + " in rows")
		}
		devices = append(devices, device)
	}
	return devices, nil
}