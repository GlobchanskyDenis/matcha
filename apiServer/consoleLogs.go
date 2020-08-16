package apiServer

import (
	"log"
	"net/http"
)

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
