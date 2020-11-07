package database

import (
	"MatchaServer/common"
	"MatchaServer/config"
)

type Storage interface {
	// setup
	Connect(conf *config.Sql) error
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
	CreateLikesTable() error
	CreateIgnoresTable() error
	CreateClaimsTable() error

	// user
	SetNewUser(mail string, passwd string) (common.User, error)
	DeleteUser(uid int) error
	UpdateUser(user common.User) error
	GetUserByUid(uid int) (common.User, error)
	GetUserByMail(mail string) (common.User, error)
	GetUsersByQuery(query string) ([]common.SearchUser, error)
	GetUserForAuth(mail string, passwd string) (common.User, error)
	IsUserExistsByMail(mail string) (bool, error)
	IsUserExistsByUid(uid int) (bool, error)

	// devices
	SetNewDevice(uid int, device string) error
	DeleteDevice(id int) error
	GetDevicesByUid(uid int) ([]common.Device, error)

	// messages
	SetNewMessage(uidSender int, uidReceiver int, body string) (int, error)
	DeleteMessage(nid int) error
	GetMessageByMid(mid int) (common.Message, error)
	GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error)

	// notifications
	SetNewNotif(uidReceiver int, uidSender int, body string) (int, error)
	DeleteNotif(nid int) error
	GetNotifByNid(nid int) (common.Notif, error)
	GetNotifByUidReceiver(uid int) ([]common.Notif, error)

	// interests
	AddInterests(unknownInterests []common.Interest) error
	GetInterests() ([]common.Interest, error)

	// photos
	SetNewPhoto(uid int, src string) (int, error)
	DeletePhoto(pid int) error
	GetPhotosByUid(uid int) ([]common.Photo, error)
	GetPhotoByPid(pid int) (common.Photo, error)

	// likes
	SetNewLike(uidSender int, uidReceiver int) error
	UnsetLike(uidSender int, uidReceiver int) error
	DropUserLikes(uid int) error
	GetFriendUsers(myUid int) ([]common.FriendUser, error)
	IsICanSpeakWithUser(myUid, otherUid int) (bool, error)

	// ignores
	SetNewIgnore(uidSender int, uidReceiver int) error
	UnsetIgnore(uidSender int, uidReceiver int) error
	DropUserIgnores(uid int) error
	GetIgnoredUsers(uidSender int) ([]common.User, error)

	// claims
	SetNewClaim(uidSender int, uidReceiver int) error
	UnsetClaim(uidSender int, uidReceiver int) error
	DropUserClaims(uid int) error
	GetClaimedUsers(uidSender int) ([]common.User, error)
}
