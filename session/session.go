package session

import (
	"fmt"
	"sync"
	"time"
	"MatchaServer/handlers"
	"github.com/gorilla/websocket"
)

type tokenChanItem struct {
	token string
	err	  error
}

type SessionUserInfo struct {
	Id          int    `json:"id,"`
	Login       string `json:"login,"`
	Passwd      string `json:"-"`
	Mail        string `json:"mail,,omitempty"`
	Phone       string `json:"phone,,omitempty"`
	Age         int    `json:"age,,omitempty"`
	Gender      string `json:"gender,,omitempty"`
	Orientation string `json:"orientation,,omitempty"`
	Biography   string `json:"orientation,,omitempty"`
	AvaPhotoID  int    `json:"avaPhotoID,,omitempty"`
	AccType		string `json:"-"`
	Rating      int    `json:"rating,"`
}

type SessionItem struct {
	UserInfo	SessionUserInfo
	Expires     time.Time
	LastVisited time.Time
	ws			[](*websocket.Conn)
}

type Session struct {
	session map[string]SessionItem
	mu      *sync.Mutex
}

func CreateSession() Session {
	var NewSession = Session{}
	NewSession.session = map[string]SessionItem{}
	NewSession.mu = &sync.Mutex{}
	return NewSession
}

func (T *Session) AddUserToSession(id int, login string, passwd string, mail string) (string, error) {
	var newItem SessionItem
	var ch = make(chan tokenChanItem)
	var ret tokenChanItem

	go func(ch chan tokenChanItem, login string) {
		token, err := handlers.TokenEncode(login)
		ch <- tokenChanItem{token, err}
	}(ch, login)

	newItem.UserInfo.Id = id
	newItem.UserInfo.Login = login
	newItem.UserInfo.Passwd = passwd
	newItem.UserInfo.Mail = mail
	newItem.LastVisited = time.Now()
	newItem.ws = []*websocket.Conn{}
	newItem.Expires = newItem.LastVisited.Add(1000000000 * 60 * 60 * 3) // 3 hour

	ret = <- ch
	if ret.err != nil {
		return ret.token, ret.err
	}

	T.mu.Lock()
	T.session[login] = newItem
	T.mu.Unlock()

	return ret.token, ret.err
}

func (T *Session) UpdateSessionUser(token string, newUserInfo SessionUserInfo) error {
	var oldLogin string
	var err error

	oldLogin, err = handlers.TokenDecode(token)
	if err != nil {
		return err
	}

	T.mu.Lock()
	item := T.session[oldLogin]
	delete(T.session, oldLogin)
	item.UserInfo = newUserInfo
	T.session[newUserInfo.Login] = item
	T.mu.Unlock()
	return nil
}

func (T Session) IsUserLogged(token string) (bool, error) {
	var isExists bool
	var login string
	var err error

	login, err = handlers.TokenDecode(token)
	if err != nil {
		return false, err
	}

	T.mu.Lock()
	_, isExists = T.session[login]
	T.mu.Unlock()
	return isExists, nil
}

func (T *Session) FindUserByToken(token string) (SessionItem, error) {
	var item SessionItem
	var isExists bool
	var login string
	var err error

	login, err = handlers.TokenDecode(token)
	if err != nil {
		return SessionItem{}, err
	}

	T.mu.Lock()
	item, isExists = T.session[login]
	T.mu.Unlock()
	if !isExists {
		///// DEBUG START //////
		T.mu.Lock()
		for tempLogin := range T.session {
			fmt.Println("\033[36m", "user", tempLogin, "logged", "\033[m")
		}
		T.mu.Unlock()
		///// DEBUG END //////
		return SessionItem{}, fmt.Errorf("hmm... looks like user %s isnt logged", login)
	}

	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, login)
		T.mu.Unlock()
		return SessionItem{}, fmt.Errorf("this session is expired")
	}

	item.LastVisited = time.Now()
	T.mu.Lock()
	T.session[login] = item
	T.mu.Unlock()

	return item, nil
}

func (T SessionItem) expiresDate() bool {
	var now = time.Now()
	var lastVisited time.Time

	if now.After(T.Expires) {
		return true
	}

	lastVisited = T.LastVisited.Add(1000000000 * 60 * 15) // 15 min after LastVisited
	if now.After(lastVisited) {
		return true
	}

	return false
}

func (T *Session) AddWSConnection(token string, newWebSocket *websocket.Conn, wsMeta string) error {
	var item SessionItem
	var err error
	var login string

	item, err = T.FindUserByToken(token)
	if err != nil {
		return err
	}
	login = item.UserInfo.Login
	// if len(item.ws) != 0 {
		// Предупредить все остальные соединения о соединении с новым устройством
		// используя в качестве информации о новом устройстве wsMeta
	// }
	item.ws = append(item.ws, newWebSocket)
	T.mu.Lock()
	T.session[login] = item
	T.mu.Unlock()
	return nil
}

func (T *Session) RemoveWSConnection(token string, webSocketToRemove *websocket.Conn) (isUserWasRemoved bool, err error) {
	var item SessionItem
	var login string

	item, err = T.FindUserByToken(token)
	if err != nil {
		return false, err
	}
	login = item.UserInfo.Login
	if len(item.ws) < 1 {
		T.mu.Lock()
		delete(T.session, login)
		T.mu.Unlock()
		return true, fmt.Errorf("hmm... looks like this user has no ws connections")
	}
	if len(item.ws) == 1 {
		T.mu.Lock()
		delete(T.session, login)
		T.mu.Unlock()
		return true, nil
	}
	for i:=0; i<len(item.ws); i++ {
		if item.ws[i] == webSocketToRemove {
			if i == 0 {
				item.ws = item.ws[1:]
			} else if i == len(item.ws) {
				item.ws = item.ws[:i]
			} else {
				item.ws = append(item.ws[:i], item.ws[(i + 1):]...)
			}
			T.mu.Lock()
			T.session[login] = item
			T.mu.Unlock()
			return false, nil
		}
	}
	return false, fmt.Errorf("hmm... looks like this websocket isnt belong to this user")
}

func (T *Session) DeleteUserSessionByLogin(login string) {
	T.mu.Lock()
	delete(T.session, login)
	T.mu.Unlock()
}

func (T Session) GetLoggedUsersInfo() []SessionUserInfo {
	var (
		user	SessionUserInfo
		users	= []SessionUserInfo{}
		login	string
	)
	T.mu.Lock()
	for login = range T.session {
		user = T.session[login].UserInfo
		users = append(users, user)
		fmt.Println("\033[36m", "user", user.Login, "is logged", "\033[m") //////////////////
	}
	T.mu.Unlock()
	return users
}
