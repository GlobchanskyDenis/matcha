package httpHandlers

import (
	. "MatchaServer/config"
	"MatchaServer/database"
	"MatchaServer/session"
	"log"
	"net/http"
)

type ConnAll struct {
	Db      database.ConnDB
	session session.Session
}

func (conn *ConnAll) ConnectAll() error {
	conn.Db = database.ConnDB{}
	err := conn.Db.Connect()
	if err != nil {
		return err
	}
	conn.session = session.CreateSession()
	return nil
}

func CreateConnectionsStruct() (*ConnAll, error) {
	var dst = &ConnAll{}

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
