package logger

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Logger struct {
	isLoggingEnabled bool
}

func (logger *Logger) Init(cfg *config.Config) {
	logger.isLoggingEnabled = cfg.IsLogEnabled
}

func (logger Logger) Log(r *http.Request, message string) {
	if logger.isLoggingEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, message)
}

func (logger Logger) LogSuccess(r *http.Request, message string) {
	if logger.isLoggingEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, GREEN_BG+"SUCCESS: "+NO_COLOR+message)
}

func (logger Logger) LogWarning(r *http.Request, message string) {
	if logger.isLoggingEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, YELLOW_BG+"WARNING: "+NO_COLOR+message)
}

func (logger Logger) LogError(r *http.Request, message string) {
	if logger.isLoggingEnabled == false {
		return
	}
	log.Printf("%s %7s %20s %s\n", r.RemoteAddr, r.Method, r.URL.Path, RED_BG+"ERROR: "+NO_COLOR+message)
}

func (logger Logger) TimeLog(r *http.Request, dur time.Duration) {
	if logger.isLoggingEnabled == false {
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
		"time : "+color+strconv.Itoa(milliseconds)+NO_COLOR)
}
