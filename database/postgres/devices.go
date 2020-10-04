package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
)

func (conn ConnDB) SetNewDevice(uid int, device string) error {
	stmt, err := conn.db.Prepare("INSERT INTO devices (uid, device) VALUES ($1, $2)")
	if err != nil {
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, device)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) DeleteDevice(id int) error {
	stmt, err := conn.db.Prepare("DELETE FROM devices WHERE id=$1")
	if err != nil {
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetDevicesByUid(uid int) ([]common.Device, error) {
	var devices = []common.Device{}
	var device common.Device

	stmt, err := conn.db.Prepare("SELECT * FROM devices WHERE uid=$1")
	if err != nil {
		return devices, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return devices, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(device.Id), &(device.Uid), &(device.Device))
		if err != nil {
			return devices, errDef.DatabaseScanError.AddOriginalError(err)
		}
		devices = append(devices, device)
	}
	return devices, nil
}
