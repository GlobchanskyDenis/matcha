package fakeSql

import (
	"MatchaServer/config"
	"MatchaServer/errDef"
	// "errors"
)

func (conn ConnFake) SetNewPhoto(uid int, src string) (int, error) {
	var photo config.Photo

	photo.Uid = uid
	photo.Src = src

	for key := 1; ; key++ {
		if _, isExists := conn.photos[key]; !isExists {
			photo.Pid = key
			break
		}
	}

	conn.photos[photo.Pid] = photo
	return photo.Pid, nil
}

func (conn ConnFake) DeletePhoto(pid int) error {
	delete(conn.photos, pid)
	return nil
}

func (conn ConnFake) GetPhotosByUid(uid int) ([]config.Photo, error) {
	var photos = []config.Photo{}
	var photo config.Photo

	for _, photo = range conn.photos {
		if photo.Uid == uid {
			photos = append(photos, photo)
		}
	}
	return photos, nil
}

func (conn ConnFake) GetPhotoByPid(pid int) (config.Photo, error) {
	var photo config.Photo

	for _, photo = range conn.photos {
		if photo.Pid == pid {
			return photo, nil
		}
	}
	return photo, errDef.RecordNotFound
}