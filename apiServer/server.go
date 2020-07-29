package apiServer

import (
	"MatchaServer/database"
	"MatchaServer/session"
)

type Server struct {
	Db		database.Storage
	session session.Session
}

func New(storage database.Storage) (*Server, error) {
	var server = &Server{}
	server.Db = storage
	err := server.Db.Connect()
	if err != nil {
		return server, err
	}
	server.session = session.CreateSession()
	return server, nil
}