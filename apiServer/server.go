package apiServer

import (
	"MatchaServer/database"
	"MatchaServer/session"
)

type Server struct {
	Db      database.Storage
	session session.Session
}

func New(newStorage database.Storage) (*Server, error) {
	var server = &Server{}
	(*server).Db = newStorage
	(*server).session = session.CreateSession()
	err := server.Db.Connect()
	return server, err
}
