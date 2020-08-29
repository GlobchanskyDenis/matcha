package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/database"
	"MatchaServer/database/fakeSql"
	"MatchaServer/database/postgres"
	"MatchaServer/errDef"
	"MatchaServer/session"
	"net/http"
)

type Server struct {
	Port          int
	Db            database.Storage
	session       session.Session
	isLogsEnabled bool
	mailConf      config.Mail
}

func (server Server) error(w http.ResponseWriter, err errDef.ApiError) {
	w.WriteHeader(err.HttpResponseStatus)
	w.Write(err.ToJson())
}

func New(confPath string) (*Server, error) {
	var conf *config.Config
	var server = &Server{}
	var newStorage database.Storage

	conf, err := config.Create(confPath)
	if err != nil {
		return nil, err
	}
	println(GREEN + "Configuration file was received" + NO_COLOR)
	server.isLogsEnabled = conf.IsLogEnabled
	server.mailConf = conf.Mail
	server.Port = conf.Port

	if !conf.IsSqlDB {
		println(YELLOW + "Using MOC object as database connection" + NO_COLOR)
		newStorage = fakeSql.New()
	} else {
		println(GREEN + "Using postgres as database connection" + NO_COLOR)
		newStorage = postgres.New()
	}
	(*server).Db = newStorage
	(*server).session = session.CreateSession()
	err = server.Db.Connect(&conf.Sql)
	return server, err
}
