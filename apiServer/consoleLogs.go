package apiServer

import (
	. "MatchaServer/common"
	"log"
	"time"
	"strconv"
	"net/http"
)

func (server *Server) Log(r *http.Request, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, message)
}

func (server *Server) LogSuccess(r *http.Request, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, GREEN_BG+"SUCCESS: "+NO_COLOR+message)
}

func (server *Server) LogWarning(r *http.Request, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, YELLOW_BG+"WARNING: "+NO_COLOR+message)
}

func (server *Server) LogError(r *http.Request, message string) {
	if server.isLogsEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, RED_BG+"ERROR: "+NO_COLOR+message)
}

func (server *Server) TimeLog(r *http.Request, dur time.Duration) {
	if server.isLogsEnabled == false {
		return
	}
	milliseconds := (int)(dur.Milliseconds())
	color := GREEN_BG
	if milliseconds > 2 {
		color = YELLOW_BG
	}
	if milliseconds > 10 {
		color = RED_BG
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, 
		"time : " + color + strconv.Itoa(milliseconds) + NO_COLOR)
}
