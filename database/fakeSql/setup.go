package fakeSql

import (
	"MatchaServer/config"
)

type ConnFake struct {
	users    map[int]config.User
	devices  map[int]config.Device
	messages map[int]config.Message
	notif    map[int]config.Notif
	photos   map[int]config.Photo
	interests map[int]config.Interest
}

func New() *ConnFake {
	return &(ConnFake{})
}

func (conn *ConnFake) Connect() error {
	conn.users = map[int]config.User{}
	conn.devices = map[int]config.Device{}
	conn.messages = map[int]config.Message{}
	conn.notif = map[int]config.Notif{}
	conn.photos = map[int]config.Photo{}
	conn.interests = map[int]config.Interest{}
	return nil
}

func (conn ConnFake) TruncateAllTables() error {
	for key, _ := range conn.users {
		delete(conn.users, key)
	}
	for key, _ := range conn.devices {
		delete(conn.devices, key)
	}
	for key, _ := range conn.messages {
		delete(conn.messages, key)
	}
	for key, _ := range conn.notif {
		delete(conn.notif, key)
	}
	for key, _ := range conn.interests {
		delete(conn.interests, key)
	}
	return nil
}

func (conn ConnFake) DropAllTables() error {
	for key, _ := range conn.users {
		delete(conn.users, key)
	}
	for key, _ := range conn.devices {
		delete(conn.devices, key)
	}
	for key, _ := range conn.messages {
		delete(conn.messages, key)
	}
	for key, _ := range conn.notif {
		delete(conn.notif, key)
	}
	for key, _ := range conn.interests {
		delete(conn.interests, key)
	}
	return nil
}

func (conn ConnFake) DropEnumTypes() error {
	return nil
}

func (conn ConnFake) CreateEnumTypes() error {
	return nil
}

func (conn *ConnFake) CreateUsersTable() error {
	conn.users = map[int]config.User{}
	return nil
}

func (conn *ConnFake) CreateNotifTable() error {
	conn.notif = map[int]config.Notif{}
	return nil
}

func (conn *ConnFake) CreateMessageTable() error {
	conn.messages = map[int]config.Message{}
	return nil
}

func (conn *ConnFake) CreatePhotoTable() error {
	conn.photos = map[int]config.Photo{}
	return nil
}

func (conn *ConnFake) CreateDevicesTable() error {
	conn.devices = map[int]config.Device{}
	return nil
}

func (conn *ConnFake) CreateInterestsTable() error {
	conn.interests = map[int]config.Interest{}
	return nil
}
