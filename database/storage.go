package database

import (
	"MatchaServer/config"
)

type Storage interface {
	// setup
	Connect() error
	Close()
	TruncateAllTables() error
	DropAllTables() error
	DropEnumTypes() error
	CreateEnumTypes() error
	CreateUsersTable() error
	CreateNotifsTable() error
	CreateMessagesTable() error
	CreatePhotosTable() error
	CreateDevicesTable() error
	CreateInterestsTable() error

	// user
	SetNewUser(mail string, passwd string) (config.User, error)
	DeleteUser(uid int) error
	UpdateUser(user config.User) error
	SearchUsersByOneFilter(filter string) ([]config.User, error)
	GetUserByUid(uid int) (config.User, error)
	GetUserByMail(mail string) (config.User, error)
	GetUserForAuth(mail string, passwd string) (config.User, error)
	GetLoggedUsers(uid []int) ([]config.User, error)
	IsUserExistsByMail(mail string) (bool, error)
	IsUserExistsByUid(uid int) (bool, error)

	// devices
	SetNewDevice(uid int, device string) error
	DeleteDevice(id int) error
	GetDevicesByUid(uid int) ([]config.Device, error)

	// messages
	SetNewMessage(uidSender int, uidReceiver int, body string) (int, error)
	DeleteMessage(nid int) error
	GetMessagesFromChat(uidSender int, uidReceiver int) ([]config.Message, error)

	// notifications
	SetNewNotif(uidReceiver int, uidSender int, body string) (int, error)
	DeleteNotif(nid int) error
	GetNotifByUidReceiver(uid int) ([]config.Notif, error)

	// interests
	AddInterests(unknownInterests []config.Interest) error
	GetInterests() ([]config.Interest, error)

	// photos
	SetNewPhoto(uid int, body []byte) (int, error)
	DeletePhoto(pid int) error
	GetPhotosByUid(uid int) ([]config.Photo, error)
	GetPhotoByPid(pid int) (config.Photo, error)
}
