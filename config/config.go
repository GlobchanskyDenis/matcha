package config

import (
	"github.com/BurntSushi/toml"
)

type Sql struct {
	Host          string `toml:"host"`
	User          string `toml:"user"`
	Pass          string `toml:"pass"`
	DBName        string `toml:"database_name"`
	DBType        string `toml:"database_type"`
	ConnectionMax int    `toml:"connection_max"`
	IsEnabled     bool   `toml:"enabled"`
}

type Mail struct {
	Host string `toml:"host"`
	Mail string `toml:"mail"`
	Pass string `toml:"pass"`
}

type Config struct {
	IsLogEnabled bool `toml:"log_enabled"`
	IsSqlDB      bool `toml:"connect_database"`
	Port         int  `toml:"port"`
	Mail         Mail `toml:"mail_server"`
	Sql          Sql  `toml:"database"`
}

func Create(path string) (*Config, error) {
	var conf = &Config{}
	_, err := toml.DecodeFile(path+"config.toml", conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
