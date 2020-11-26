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
	CreateHistoryTable() error

	// user
	SetNewUser(mail string, passwd string) (common.User, error)
	DeleteUser(uid int) error
	UpdateUser(user common.User) error
	GetUserByUid(uid int) (common.User, error)
	GetTargetUserByUid(myUid int, targetUid int) (common.TargetUser, error)
	GetUserByMail(mail string) (common.User, error)
	GetUsersByQuery(query string, sourceUser common.User) ([]common.SearchUser, error)
	GetUserForAuth(mail string, passwd string) (common.User, error)
	IsUserExistsByMail(mail string) (bool, error)
	IsUserExistsByUid(uid int) (bool, error)
	GetUserWithLikeInfo(targetUid int, myUid int) (common.SearchUser, error)

	// devices
	SetNewDevice(uid int, device string) error
	DeleteDevice(id int) error
	GetDevicesByUid(uid int) ([]common.Device, error)

	// messages
	SetNewMessage(uidSender int, uidReceiver int, body string) (int, error)
	DeleteMessage(nid int) error
	GetMessageByMid(mid int) (common.Message, error)
	GetMessagesFromChat(uidSender int, uidReceiver int) ([]common.Message, error)
	GetActiveMessages(uidReceiver int) ([]common.Message, error)
	SetMessageInactive(mid int) error

	// notifications
	SetNewNotif(uidReceiver int, uidSender int, body string) (int, error)
	DeleteNotif(nid int) error
	DropUserNotifs(uid int) error
	DropReceiverNotifs(uid int) error
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
	GetUsersLikedMe(myUid int) ([]common.User, error)
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

	// history
	SetNewHistoryReference(uid int, targetUid int) error
	GetHistoryReferencesByUid(uid int) ([]common.HistoryReference, error)
	GetHistoryReferencesByTargetUid(uid int) ([]common.HistoryReference, error)
	DropUserHistory(uid int) error
}
