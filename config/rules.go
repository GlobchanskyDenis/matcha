package config

import (
)

const (
	DB_HOST	= "localhost"
	DB_NAME	= "matcha_db"
	DB_USER	= "bsabre"
	DB_PASSWD = "23"
	DB_TYPE	= "postgres"

	PASSWD_MIN_LEN = 6
	MAIL_MAX_LEN = 30
	MAIL_MIN_LEN = 6
	NAME_MAX_LEN = 30
	BIOGRAPHY_MAX_LEN = 300

	RED		= "\033[31m"
	GREEN	= "\033[32m"
	YELLOW	= "\033[33m"
	BLUE	= "\033[34m"
	RED_BG		= "\033[41;30m"
	GREEN_BG	= "\033[42;30m"
	YELLOW_BG	= "\033[43;30m"
	BLUE_BG		= "\033[44;30m"
	NO_COLOR	= "\033[m"
)

type User struct {
	Uid         int    `json:"uid"`
	Mail        string `json:"mail,,omitempty"`
	Passwd      string `json:"-"`
	Fname		string `json:"fname"`
	Lname		string `json:"lname"`
	Age         int    `json:"age,,omitempty"`
	Gender      string `json:"gender,,omitempty"`
	Orientation string `json:"orientation,,omitempty"`
	Biography   string `json:"biography,,omitempty"`
	AvaPhotoID  int    `json:"avaPhotoID,,omitempty"`
	AccType		string `json:"-"`
	Rating      int    `json:"rating"`
}
