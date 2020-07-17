package handlers

import (
	"hash/crc32"
	"strconv"
	"time"
	"fmt"
	"errors"
	"unicode/utf8"
	"MatchaServer/config"

	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

const (
	passwdSalt = "+++"
	masterKey = "passphrasewhichneedstobe32bytes!"
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

func isDigit(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isNameRunePermitted(c rune) bool {
	if isLetter(c) || isDigit(c) {
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
	if isDigit(c) {
		return true
	}
	if c == '_' || c == '-' || c == '.' || c == '@' {
		return true
	}
	return false
}

func CheckPasswd(passwd string) error {
	var (
		wasLetter bool
		wasDigit bool
		wasSpacialChar bool
		buf = []rune(passwd)
	)
	if utf8.RuneCountInString(passwd) < config.PASSWD_MIN_LEN {
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

	if utf8.RuneCountInString(mail) < config.MAIL_MIN_LEN {
		return fmt.Errorf("too short mail address")
	}
	if utf8.RuneCountInString(mail) > config.MAIL_MAX_LEN {
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

func CheckName(name string) error {
	var runeSlice = []rune(name)

	if len(name) > config.NAME_MAX_LEN {
		return errors.New("too long name length")
	}
	if utf8.RuneCountInString(name) < 1 {
		return errors.New("name is empty")
	}
	if !isLetter(runeSlice[0]) {
		return errors.New("first name symbol should be letter")
	}
	if !isLetter(runeSlice[(len(runeSlice) - 1)]) && !isDigit(runeSlice[(len(runeSlice) - 1)]) {
		return errors.New("last name symbol should be letter or digit")
	}
	for i := 0; i < len(runeSlice); i++ {
		if !isNameRunePermitted(runeSlice[i]) {
			return errors.New("name letter '" + string(runeSlice[i]) + "' is not permitted")
		}
	}
	return nil
}

func CheckGender(gender string) error {
	if gender != "male" && gender != "female" {
		return errors.New("gender '" + gender + "' not exist in database")
	}
	return nil
}

func CheckOrientation(orientation string) error {
	if orientation != "getero" && orientation != "bi" && orientation != "gay" {
		return errors.New("orientation '" + orientation + "' not exist in database")
	}
	return nil
}

func CheckBiography(biography string) error {
	if len(biography) > config.BIOGRAPHY_MAX_LEN {
		return errors.New("too long biography length")
	}
	return nil
}

func PasswdHash(passwd string) string {
	passwd += passwdSalt
	crcH := crc32.ChecksumIEEE([]byte(passwd))
	passwdHash := strconv.FormatUint(uint64(crcH), 20)
	return passwdHash
}

func TokenWebSocketAuth(uid int) string {

	str := strconv.Itoa(uid)
	curTime := time.Now()

	dataToHash := fmt.Sprintf("%s%s", str, curTime)
	tmpHash := crc32.ChecksumIEEE([]byte(dataToHash))
	hash := strconv.FormatUint(uint64(tmpHash), 35)
	token := string(hash[:])

	dataToHash = fmt.Sprintf("%s", curTime)
	tmpHash = crc32.ChecksumIEEE([]byte(dataToHash))
	hash = strconv.FormatUint(uint64(tmpHash), 35)
	token += string(hash[:])

	return token
}

func TokenEncode(uid int) (string, error) {

	// Thanks to https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
	// for good explanation of Encoding with masterKey
	// AES - Advanced Encryption Standard

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(masterKey))
	// if there are any errors, handle them
	if err != nil {
		return "", err
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return "", err
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	token := gcm.Seal(nonce, nonce, []byte(strconv.Itoa(uid)), nil)
	return base64.URLEncoding.EncodeToString(token), nil
}

func TokenDecode(token string) (int, error) {
	encodedToken, _ := base64.URLEncoding.DecodeString(token)

	c, err := aes.NewCipher([]byte(masterKey))
	if err != nil {
		return 0, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return 0, err
	}

	nonceSize := gcm.NonceSize()
	if len(encodedToken) < nonceSize {
		return 0, fmt.Errorf("size error in decoding")
	}

	nonce, encodedToken := encodedToken[:nonceSize], encodedToken[nonceSize:]
	desired, err := gcm.Open(nil, nonce, encodedToken, nil)
	if err != nil {
		return 0, err
	}
	uid, err := strconv.Atoi(string(desired))
	if err != nil {
		return 0, err
	}
	return uid, nil
}