package httpHandlers

import (
	. "MatchaServer/config"
	"MatchaServer/database"
	"MatchaServer/session"
	"log"
	"net/http"
)

type Server struct {
	Db		database.Storage
	session session.Session
}

func (server *Server) New(storage database.Storage) error {
	server.Db = storage
	err := server.Db.Connect()
	if err != nil {
		return err
	}
	server.session = session.CreateSession()
	return nil
}

func CreateConnectionsStruct() (*Server, error) {
	var dst = &Server{}

	err := dst.Db.Connect()
	if err != nil {
		return dst, err
	}
	dst.session = session.CreateSession()
	return dst, nil
}

func consoleLog(r *http.Request, section string, message string) {
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, message)
}

func consoleLogSuccess(r *http.Request, section string, message string) {
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, GREEN_BG+"SUCCESS: "+NO_COLOR+message)
}

func consoleLogWarning(r *http.Request, section string, message string) {
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, YELLOW_BG+"WARNING: "+NO_COLOR+message)
}

func consoleLogError(r *http.Request, section string, message string) {
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, RED_BG+"ERROR: "+NO_COLOR+message)
}
