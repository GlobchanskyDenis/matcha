package config

import (
)

const (
	DB_HOST	= "localhost"
	DB_NAME	= "matcha_db"
	DB_USER	= "bsabre"
	DB_PASSWD	= "23"
	DB_TYPE	= "postgres"

	LOGIN_MIN_LEN = 3
	LOGIN_MAX_LEN = 15
	PASSWD_MIN_LEN = 6
	MAIL_MAX_LEN = 30
	MAIL_MIN_LEN = 6
	// PHONE_MIN_LEN = 6
	// PHONE_MAX_LEN = 18

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