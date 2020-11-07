package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/config"
)

type relationMock struct {
	uidSender   int
	uidReceiver int
}

type ConnFake struct {
	users     map[int]common.User
	devices   map[int]common.Device
	messages  map[int]common.Message
	notif     map[int]common.Notif
	photos    map[int]common.Photo
	interests map[int]common.Interest
	likes     []relationMock
	ignores   []relationMock
	claims    []relationMock
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
	conn.likes = []relationMock{}
	conn.ignores = []relationMock{}
	conn.claims = []relationMock{}
	conn.SetNewUser("admin", "admin")
	return nil
}

func (conn *ConnFake) Close() {
}

func (conn ConnFake) TruncateAllTables() error {
	conn.users = map[int]common.User{}
	conn.devices = map[int]common.Device{}
	conn.messages = map[int]common.Message{}
	conn.notif = map[int]common.Notif{}
	conn.photos = map[int]common.Photo{}
	conn.interests = map[int]common.Interest{}
	conn.likes = []relationMock{}
	conn.ignores = []relationMock{}
	conn.claims = []relationMock{}
	return nil
}

func (conn *ConnFake) DropAllTables() error {
	conn.users = map[int]common.User{}
	conn.devices = map[int]common.Device{}
	conn.messages = map[int]common.Message{}
	conn.notif = map[int]common.Notif{}
	conn.photos = map[int]common.Photo{}
	conn.interests = map[int]common.Interest{}
	conn.likes = []relationMock{}
	conn.ignores = []relationMock{}
	conn.claims = []relationMock{}
	return nil
}

func (conn *ConnFake) DropEnumTypes() error {
	return nil
}

func (conn *ConnFake) CreateEnumTypes() error {
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
	conn.likes = []relationMock{}
	return nil
}

func (conn *ConnFake) CreateIgnoresTable() error {
	conn.ignores = []relationMock{}
	return nil
}

func (conn *ConnFake) CreateClaimsTable() error {
	conn.claims = []relationMock{}
	return nil
}

func isSliceContainsElement(slice []relationMock, element relationMock) bool {
	for _, item := range slice {
		if item.uidSender == element.uidSender &&
			item.uidReceiver == element.uidReceiver {
			return true
		}
	}
	return false
}

func (conn ConnFake) isUserExists(uid int) bool {
	for _, user := range conn.users {
		if user.Uid == uid {
			return true
		}
	}
	return false
}

func (conn *ConnFake) changeUserRating(uid int, deltaRating int) {
	for i, user := range conn.users {
		if user.Uid == uid {
			user.Rating += deltaRating
			conn.users[i] = user
			return
		}
	}
}
