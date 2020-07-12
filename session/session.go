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

type SessionItem struct {
	Uid			int
	Expires     time.Time
	LastVisited time.Time
	TokenWS		string
	ws			[](*websocket.Conn)
}

type Session struct {
	session map[int]SessionItem
	mu      *sync.Mutex
}

func CreateSession() Session {
	var NewSession = Session{}
	NewSession.session = map[int]SessionItem{}
	NewSession.mu = &sync.Mutex{}
	return NewSession
}

func (T *Session) AddUserToSession(uid int) (string, error) {
	var newItem SessionItem
	var ch = make(chan tokenChanItem)
	var ret tokenChanItem

	go func(ch chan tokenChanItem, uid int) {
		token, err := handlers.TokenEncode(uid)
		ch <- tokenChanItem{token, err}
	}(ch, uid)

	newItem.Uid = uid
	newItem.LastVisited = time.Now()
	newItem.ws = []*websocket.Conn{}
	newItem.Expires = newItem.LastVisited.Add(1000000000 * 60 * 60 * 3) // 3 hour

	ret = <- ch
	if ret.err != nil {
		return ret.token, ret.err
	}

	T.mu.Lock()
	T.session[uid] = newItem
	T.mu.Unlock()

	return ret.token, ret.err
}

func (T *Session) IsUserLoggedByToken(token string) (bool, error) {
	uid, err := handlers.TokenDecode(token)
	if err != nil {
		return false, err
	}
	T.mu.Lock()
	item, isExists := T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return false, nil
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return false, nil
	}
	item.LastVisited = time.Now()
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return true, nil
}

func (T *Session) IsUserLoggedByUid(uid int) bool {
	T.mu.Lock()
	item, isExists := T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return false
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return false
	}
	item.LastVisited = time.Now()
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return true
}

func (T *Session) FindUserByToken(token string) (SessionItem, error) {
	var item SessionItem
	var isExists bool
	var uid int
	var err error

	uid, err = handlers.TokenDecode(token)
	if err != nil {
		return SessionItem{}, err
	}
	T.mu.Lock()
	item, isExists = T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return SessionItem{}, fmt.Errorf("hmm... looks like user #%d isnt logged", uid)
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return SessionItem{}, fmt.Errorf("this session is expired")
	}
	item.LastVisited = time.Now()
	T.mu.Lock()
	T.session[uid] = item
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

func (T *Session) CreateTokenWS(uid int) (string, error) {
	var ch = make(chan string)
	var item SessionItem
	var isExists bool

	go func(ch chan string, uid int) {
		ch <- handlers.TokenWebSocketAuth(uid)
	}(ch, uid)

	T.mu.Lock()
	item, isExists = T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return "", fmt.Errorf("hmm... looks like user #%d isnt logged", uid)
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return "", fmt.Errorf("this session is expired")
	}
	item.LastVisited = time.Now()
	item.TokenWS = <- ch
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return item.TokenWS, nil
}

func (T *Session) GetTokenWS(uid int) (string, error) {
	var item SessionItem
	var isExists bool

	T.mu.Lock()
	item, isExists = T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return "", fmt.Errorf("hmm... looks like user #%d isnt logged", uid)
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return "", fmt.Errorf("this session is expired")
	}
	return item.TokenWS, nil
}

func (T *Session) AddWSConnection(token string, newWebSocket *websocket.Conn, wsMeta string) error {
	var item SessionItem
	var err error
	var uid int

	item, err = T.FindUserByToken(token)
	if err != nil {
		return err
	}
	uid = item.Uid
	// if len(item.ws) != 0 {
		// Предупредить все остальные соединения о соединении с новым устройством
		// используя в качестве информации о новом устройстве wsMeta
	// }
	item.ws = append(item.ws, newWebSocket)
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return nil
}

func (T *Session) RemoveWSConnection(token string, webSocketToRemove *websocket.Conn) (isUserWasRemoved bool, err error) {
	var item SessionItem
	var uid int

	item, err = T.FindUserByToken(token)
	if err != nil {
		return false, err
	}
	uid = item.Uid
	if len(item.ws) < 1 {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return true, fmt.Errorf("hmm... looks like this user has no ws connections")
	}
	if len(item.ws) == 1 {
		T.mu.Lock()
		delete(T.session, uid)
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
			T.session[uid] = item
			T.mu.Unlock()
			return false, nil
		}
	}
	return false, fmt.Errorf("hmm... looks like this websocket isnt belong to this user")
}

func (T *Session) DeleteUserSessionByUid(uid int) {
	T.mu.Lock()
	delete(T.session, uid)
	T.mu.Unlock()
}

func (T Session) GetLoggedUsersUidSlice() []int {
	var (
		uid		int
		result	= []int{}
	)
	T.mu.Lock()
	for uid = range T.session {
		result = append(result, uid)
		// fmt.Println("\033[36m", "user", user.Login, "is logged", "\033[m") //////////////////
	}
	T.mu.Unlock()
	return result
}
