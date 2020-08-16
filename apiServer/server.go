package apiServer

import (
	"MatchaServer/database"
	"MatchaServer/errDef"
	"MatchaServer/session"
	"net/http"
)

const (
	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	RED_BG    = "\033[41;30m"
	GREEN_BG  = "\033[42;30m"
	YELLOW_BG = "\033[43;30m"
	BLUE_BG   = "\033[44;30m"
	NO_COLOR  = "\033[m"
)

type Server struct {
	Db      database.Storage
	session session.Session
}

func (server Server) error(w http.ResponseWriter, err errDef.ApiError) {
	w.WriteHeader(err.HttpResponseStatus)
	w.Write(err.ToJson())
}

func New(newStorage database.Storage) (*Server, error) {
	var server = &Server{}
	(*server).Db = newStorage
	(*server).session = session.CreateSession()
	err := server.Db.Connect()
	return server, err
}
