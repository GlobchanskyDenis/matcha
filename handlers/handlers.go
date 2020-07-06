package handlers

import (
	"hash/crc32"
	"strconv"
	"time"
	"fmt"
	"unicode/utf8"
	. "MatchaServer/config"
)

const (
	passwdSalt = "+++"
)

func isLetter(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= 'а' && c <= 'я' {
		return true
	}
	if c >= 'А' && c <= 'Я' {
		return true
	}
	return false
}

func isLoginRunePermitted(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	if c >= 'а' && c <= 'я' {
		return true
	}
	if c >= 'А' && c <= 'Я' {
		return true
	}
	if c == '_' || c == '-' || c == ' ' {
		return true
	}
	return false
}

func isMailRunePermitted(c rune) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	if c >= '0' && c <= '9' {
		return true
	}
	if c == '_' || c == '-' || c == '.' || c == '@' {
		return true
	}
	return false
}

func isPhoneRunePermitted(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	if c == ' ' || c == '-' || c == '(' || c == ')' {
		return true
	}
	return false
}

func CheckLogin(login string) error {
	var (
		runeSlice = []rune(login)
		length	= len(runeSlice)
		wasLetter bool
	)

	if utf8.RuneCountInString(login) < LOGIN_MIN_LEN {
		return fmt.Errorf("too short login")
	}
	if utf8.RuneCountInString(login) > LOGIN_MAX_LEN {
		return fmt.Errorf("too long login")
	}

	if runeSlice[0] == ' ' {
		return fmt.Errorf("first symbol should not be space")
	}
	if runeSlice[length - 1] == ' ' {
		return fmt.Errorf("last symbol should not be space")
	}

	for i:=0; i<length; i++ {
		if !isLoginRunePermitted(runeSlice[i]) {
			return fmt.Errorf("forbidden symbol in login")
		}
		if isLetter(runeSlice[i]) {
			wasLetter = true
		}
	}
	if !wasLetter {
		return fmt.Errorf("no letters in login")
	}
	return nil
}

func CheckPasswd(passwd string) error {
	var (
		wasLetter bool
		wasDigit bool
		wasSpacialChar bool
		buf = []rune(passwd)
	)
	if utf8.RuneCountInString(passwd) < PASSWD_MIN_LEN {
		return fmt.Errorf("too short password")
	}

	for i:=0; i<len(buf); i++ {
		if isLetter(buf[i]) {
			wasLetter = true
		}
		if buf[i] >= '0' && buf[i] <= '9' {
			wasDigit = true
		}
		if buf[i] == '!' || buf[i] == '@' || buf[i] == '#' || buf[i] == '$' ||
				buf[i] == '%' || buf[i] == '^' || buf[i] == '&' || buf[i] == '*' {
			wasSpacialChar = true
		}
	}
	if !wasLetter {
		return fmt.Errorf("Password should contain letters")
	}
	if !wasDigit {
		return fmt.Errorf("Password should contain digits")
	}
	if !wasSpacialChar {
		return fmt.Errorf("Password should contain special chars")
	}
	return nil
}

func CheckMail(mail string) error {
	var (
		buf = []rune(mail)
		length = len(buf)
		doggyCount int
		dots int
	)

	if utf8.RuneCountInString(mail) < MAIL_MIN_LEN {
		return fmt.Errorf("too short mail address")
	}
	if utf8.RuneCountInString(mail) > MAIL_MAX_LEN {
		return fmt.Errorf("too long mail address")
	}

	if buf[0] == '_' || buf[0] == '-' || buf[0] == '@' ||
			buf[0] == '.' || (buf[0] >= '0' && buf[0] <= '9') {
				return fmt.Errorf("invalid first mail address symbol")
	}

	if buf[length - 1] == '_' || buf[length - 1] == '-' || buf[length - 1] == '@' ||
			buf[length - 1] == '.' || (buf[length - 1] >= '0' && buf[length - 1] <= '9') {
				return fmt.Errorf("invalid last mail address symbol")
	}

	for i:=0; i<length; i++ {
		if !isMailRunePermitted(buf[i]) {
			return fmt.Errorf("forbidden symbol in mail")
		}
		if (buf[i] == '@') {
			doggyCount++
			if i>0 && buf[i - 1] == '.' {
				return fmt.Errorf("invalid mail address")
			}
		}
		if (buf[i] == '.' && doggyCount > 0) {
			dots++
			if buf[i - 1] == '.' || buf[i - 1] == '@' {
				return fmt.Errorf("invalid mail address")
			}
		}
	}
	if doggyCount != 1 {
		return fmt.Errorf("invalid amount of '@' symbols in mail address")
	}
	if dots != 1 && dots != 2 {
		return fmt.Errorf("invalid amount of '.' symbols in mail address")
	}
	return nil
}

func CheckPhone(phone string) error {
	var (
		buf = []rune(phone)
		length = len(buf)
		startScope bool
		endScope bool
	)

	if length < PHONE_MIN_LEN {
		return fmt.Errorf("too short phone number")
	}
	if length > PHONE_MAX_LEN {
		return fmt.Errorf("too long phone number")
	}

	if !(buf[0] >= '0' && buf[0] <= '9') && buf[0] != '+' {
		return fmt.Errorf("first char of phone number must be digit or +")
	}

	if buf[length - 1] < '0' || buf[length - 1] > '9' {
		return fmt.Errorf("last char of phone number must be digit")
	}

	for i:=1; i<length; i++ {
		if !isPhoneRunePermitted(buf[i]) {
			return fmt.Errorf("invalid phone number")
		}
		if (buf[i] < '0' || buf[i] > '9') && (buf[i - 1] < '0' || buf[i - 1] > '9') {
			return fmt.Errorf("invalid phone number")
		}
		if buf[i] == '(' && startScope {
			return fmt.Errorf("invalid phone number")
		}
		if buf[i] == ')' && endScope {
			return fmt.Errorf("invalid phone number")
		}
		if buf[i] == '(' {
			startScope = true
		}
		if buf[i] == ')' {
			endScope = true
		}
	}
	if startScope != endScope {
		return fmt.Errorf("invalid phone number")
	}
	return nil
}

func PasswdHash(passwd string) string {
	passwd += passwdSalt
	crcH := crc32.ChecksumIEEE([]byte(passwd))
	passwdHash := strconv.FormatUint(uint64(crcH), 20)
	return passwdHash
}

func TokenHash(login string, lastVisited time.Time) string {

	dataToHash := fmt.Sprintf("%s%s", login, lastVisited)
	tmpHash := crc32.ChecksumIEEE([]byte(dataToHash))
	hash := strconv.FormatUint(uint64(tmpHash), 35)
	token := string(hash[:])

	dataToHash = fmt.Sprintf("%s", lastVisited)
	tmpHash = crc32.ChecksumIEEE([]byte(dataToHash))
	hash = strconv.FormatUint(uint64(tmpHash), 35)
	token += string(hash[:])

	return token
}

func TokenEncode(login string) (string, error) {

	var currentTime = time.Now()
	var token string
	var err error

	token = login + fmt.Sprintf("%s", currentTime)
	return token, nil
}

func TokenDecode(token string) (string, error) {
	var length = len(token)
	var lenLogin = length - 6 // заменить потом
	var slice = []byte(token)
	var login = [:lenlogin]slice
	return string(login), nil
}