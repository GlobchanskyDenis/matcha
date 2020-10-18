package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/config"
)

type ConnFake struct {
	users     map[int]common.User
	devices   map[int]common.Device
	messages  map[int]common.Message
	notif     map[int]common.Notif
	photos    map[int]common.Photo
	interests map[int]common.Interest
	likes     map[int]struct {
		uidSender   int
		uidReceiver int
	}
}

func New() *ConnFake {
	return &(ConnFake{})
}

func (conn *ConnFake) Connect(conf *config.Sql) error {
	conn.users = map[int]common.User{}
	conn.devices = map[int]common.Device{}
	conn.messages = map[int]common.Message{}
	conn.notif = map[int]common.Notif{}
	conn.photos = map[int]common.Photo{}
	conn.interests = map[int]common.Interest{}
	conn.likes = map[int]struct {
		uidSender   int
		uidReceiver int
	}{}
	return nil
}

func (conn *ConnFake) Close() {
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
	for key, _ := range conn.likes {
		delete(conn.likes, key)
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
	conn.users = map[int]common.User{}
	return nil
}

func (conn *ConnFake) CreateNotifsTable() error {
	conn.notif = map[int]common.Notif{}
	return nil
}

func (conn *ConnFake) CreateMessagesTable() error {
	conn.messages = map[int]common.Message{}
	return nil
}

func (conn *ConnFake) CreatePhotosTable() error {
	conn.photos = map[int]common.Photo{}
	return nil
}

func (conn *ConnFake) CreateDevicesTable() error {
	conn.devices = map[int]common.Device{}
	return nil
}

func (conn *ConnFake) CreateInterestsTable() error {
	conn.interests = map[int]common.Interest{}
	return nil
}

func (conn *ConnFake) CreateLikesTable() error {
	conn.likes = map[int]struct {
		uidSender   int
		uidReceiver int
	}{}
	return nil
}
