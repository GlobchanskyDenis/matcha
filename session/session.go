package session

import (
	"fmt"
	"sync"
	"time"
	"MatchaServer/handlers"
)

type SessionItem struct {
	Id          int
	Login       string
	Expires     time.Time
	LastVisited time.Time
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

func (T *Session) AddUserToSession(login string, id int) string {
	var newItem SessionItem
	var token string

	newItem.Id = id
	newItem.Login = login
	newItem.LastVisited = time.Now()
	newItem.Expires = newItem.LastVisited.Add(1000000000 * 60 * 60) // 1 hour

	token = handlers.TokenHash(login, newItem.LastVisited)

	_, isExists := (*T).session[token]
	for isExists {
		token = handlers.TokenHash(login, newItem.LastVisited)
		_, isExists = (*T).session[token]
	}

	(*T).mu.Lock()
	(*T).session[token] = newItem
	(*T).mu.Unlock()
	return token
}

func (T *Session) FindUserByToken(token string) (SessionItem, error) {
	var item SessionItem
	var isExists bool

	(*T).mu.Lock()
	item, isExists = (*T).session[token]
	(*T).mu.Unlock()
	if !isExists {
		return SessionItem{}, fmt.Errorf("this token isnt exists")
	}

	if item.expiresDate() {
		(*T).mu.Lock()
		delete((*T).session, token)
		(*T).mu.Unlock()
		return SessionItem{}, fmt.Errorf("this session is expired")
	}

	item.LastVisited = time.Now()
	(*T).mu.Lock()
	(*T).session[token] = item
	(*T).mu.Unlock()

	return item, nil
}

func (T SessionItem) expiresDate() bool {
	var now = time.Now()
	var lastVisited time.Time

	if now.After(T.Expires) {
		return true
	}

	lastVisited = T.LastVisited.Add(1000000000 * 60 * 5) // 5 min after LastVisited
	if now.After(lastVisited) {
		return true
	}

	return false
}
