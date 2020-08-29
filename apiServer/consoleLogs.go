package apiServer

import (
	. "MatchaServer/common"
	"log"
	"net/http"
)

func (server *Server) Log(r *http.Request, section string, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, message)
}

func (server *Server) LogSuccess(r *http.Request, section string, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, GREEN_BG+"SUCCESS: "+NO_COLOR+message)
}

func (server *Server) LogWarning(r *http.Request, section string, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, YELLOW_BG+"WARNING: "+NO_COLOR+message)
}

func (server *Server) LogError(r *http.Request, section string, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, section, RED_BG+"ERROR: "+NO_COLOR+message)
}
