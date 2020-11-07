package session

import (
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"github.com/gorilla/websocket"
	"strconv"
	"sync"
	"time"
)

type tokenChanItem struct {
	token string
	err   error
}

type wsItem struct {
	conn *websocket.Conn
}

type sessionItem struct {
	Uid         int
	Expires     time.Time
	LastVisited time.Time
	ws          map[string]wsItem
}

type Session struct {
	session map[int]sessionItem
	mu      *sync.Mutex
}

func CreateSession() Session {
	var NewSession = Session{}
	NewSession.session = map[int]sessionItem{}
	NewSession.mu = &sync.Mutex{}
	return NewSession
}

func (T *Session) AddUserToSession(uid int) (string, error) {
	var newItem sessionItem
	var ch = make(chan tokenChanItem)
	var ret tokenChanItem

	go func(ch chan tokenChanItem, uid int) {
		token, err := handlers.TokenUidEncode(uid)
		ch <- tokenChanItem{token, err}
	}(ch, uid)

	if T.IsUserLoggedByUid(uid) {
		T.mu.Lock()
		newItem = T.session[uid]
		T.mu.Unlock()
		newItem.LastVisited = time.Now()
	} else {
		newItem.Uid = uid
		newItem.LastVisited = time.Now()
		newItem.ws = map[string]wsItem{}                                    //[]*websocket.Conn{}
		newItem.Expires = newItem.LastVisited.Add(1000000000 * 60 * 60 * 3) // 3 hour
	}

	ret = <-ch
	if ret.err != nil {
		return ret.token, errors.NewArg("Ошибка добавления пользователя в сессию",
			"user add to session error").AddOriginalError(ret.err)
	}

	T.mu.Lock()
	T.session[uid] = newItem
	T.mu.Unlock()

	return ret.token, nil
}

func (T *Session) IsUserLoggedByToken(token string) (bool, error) {
	uid, err := handlers.TokenUidDecode(token)
	if err != nil {
		return false, errors.NewArg("Ошибка декодирования токена", "Token decode error").AddOriginalError(err)
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

func (T *Session) findUserByToken(token string) (sessionItem, error) {
	var item sessionItem
	var isExists bool
	var uid int
	var err error

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		return sessionItem{}, errors.NewArg("Ошибка декодирования токена", "Token decode error").AddOriginalError(err)
	}
	T.mu.Lock()
	item, isExists = T.session[uid]
	T.mu.Unlock()
	if !isExists {
		return sessionItem{}, errors.NewArg("Пользователь #"+strconv.Itoa(uid)+"не залогинен",
			"user #"+strconv.Itoa(uid)+" isnt logged")
	}
	if item.expiresDate() {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return sessionItem{}, errors.NewArg("Сессия просрочена", "this session is expired")
	}
	item.LastVisited = time.Now()
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return item, nil
}

func (T sessionItem) expiresDate() bool {
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

func (T *Session) AddWSConnection(uid int, newWebSocket *websocket.Conn, userAgent string) {
	var item sessionItem

	T.mu.Lock()
	item = T.session[uid]
	item.ws[userAgent] = wsItem{conn: newWebSocket}
	T.session[uid] = item
	T.mu.Unlock()
}

func (T *Session) RemoveWSConnection(uid int, userAgent string, isLogout bool) (bool, error) {
	var item sessionItem

	T.mu.Lock()
	item = T.session[uid]
	T.mu.Unlock()

	// invalid case
	if len(item.ws) < 1 {
		T.mu.Lock()
		delete(T.session, uid)
		T.mu.Unlock()
		return true, errors.NewArg("у вашего пользователя нет открытых ws соединений",
			"this user has no ws connections")
	}

	// Проверяю чтобы такой юзер агент вообще существовал
	item_, isExist := item.ws[userAgent]
	if !isExist {
		return false, errors.NewArg("не найден браузер", "user device is not found")
	}

	if isLogout {
		if len(item.ws) == 1 {
			// полное удаление сессии юзера (когда соединение последнее)
			T.mu.Lock()
			delete(T.session, uid)
			T.mu.Unlock()
			return true, nil
		} else {
			// удаление только конкретного юзер агента из сессии
			delete(item.ws, userAgent)
			T.mu.Lock()
			T.session[uid] = item
			T.mu.Unlock()
			return false, nil
		}
	}

	// Удаляю ws из сессии. При этом элемент в мапе с именем юзер агента остается в сессии
	// Никто не сможет нам написать, или отправить уведомление, но сессия будет ожидать
	// Переподключения пользователя
	item_.conn = nil
	item.ws[userAgent] = item_
	T.mu.Lock()
	T.session[uid] = item
	T.mu.Unlock()
	return false, nil
}

func (T *Session) SendNotifToLoggedUser(uidReceiver int, uidSender int, notifBody string) error {
	var item sessionItem
	var message string
	var err error
	var ws wsItem

	T.mu.Lock()
	item = T.session[uidReceiver]
	T.mu.Unlock()
	for _, ws = range item.ws {
		if ws.conn == nil {
			continue
		}
		message = `{"type":"notif","uidSender":"` + strconv.Itoa(uidSender) + `","body":"` + notifBody + `"}`
		err = ws.conn.WriteMessage(1, []byte(message))
		if err != nil {
			return err
		}
	}
	return nil
}

func (T *Session) SendMessageToLoggedUser(uidReceiver int, uidSender int, messageBody string) error {
	var item sessionItem
	var message string
	var err error
	var ws wsItem

	T.mu.Lock()
	item = T.session[uidReceiver]
	T.mu.Unlock()
	for _, ws = range item.ws {
		if ws.conn == nil {
			continue
		}
		message = `{"type":"message","uidSender":"` + strconv.Itoa(uidSender) + `","body":"` + messageBody + `"}`
		err = ws.conn.WriteMessage(1, []byte(message))
		if err != nil {
			return err
		}
	}
	return nil
}

func (T *Session) DeleteUserSessionByUid(uid int) {
	T.mu.Lock()
	delete(T.session, uid)
	T.mu.Unlock()
}

func (T Session) GetLoggedUsersUidSlice() []int {
	var (
		uid    int
		result = []int{}
	)
	T.mu.Lock()
	for uid = range T.session {
		result = append(result, uid)
	}
	T.mu.Unlock()
	return result
}
