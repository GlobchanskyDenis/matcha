package handlers

import (
	"MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errors"
	"hash/crc32"
	"net/smtp"
	"strconv"
	"time"
	"unicode/utf8"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

const (
	passSalt  = "+++"
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

func CheckPass(pass string) error {
	var (
		wasLetter      bool
		wasDigit       bool
		wasSpacialChar bool
		buf            = []rune(pass)
	)
	if utf8.RuneCountInString(pass) < config.PASS_MIN_LEN {
		return errors.NewArg("слишком короткий пароль", "too short password")
	}

	for i := 0; i < len(buf); i++ {
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
		return errors.NewArg("пароль должен содержать буквы", "Password should contain letters")
	}
	if !wasDigit {
		return errors.NewArg("пароль должен содержать цифры", "Password should contain digits")
	}
	if !wasSpacialChar {
		return errors.NewArg("пароль должен содержать специальные символы", "Password should contain special chars")
	}
	return nil
}

func CheckMail(mail string) error {
	var (
		buf        = []rune(mail)
		length     = len(buf)
		doggyCount int
		dots       int
	)

	if utf8.RuneCountInString(mail) < config.MAIL_MIN_LEN {
		return errors.NewArg("слишком короткий почтовый адрес", "too short mail address")
	}
	if utf8.RuneCountInString(mail) > config.MAIL_MAX_LEN {
		return errors.NewArg("слишком длинный почтовый адрес", "too long mail address")
	}

	if buf[0] == '_' || buf[0] == '-' || buf[0] == '@' ||
		buf[0] == '.' || (buf[0] >= '0' && buf[0] <= '9') {
		return errors.NewArg("первый символ почтового адреса невалиден",
			"invalid first mail address symbol")
	}

	if buf[length-1] == '_' || buf[length-1] == '-' || buf[length-1] == '@' ||
		buf[length-1] == '.' || (buf[length-1] >= '0' && buf[length-1] <= '9') {
		return errors.NewArg("последний символ почтового адреса невалиден",
			"invalid last mail address symbol")
	}

	for i := 0; i < length; i++ {
		if !isMailRunePermitted(buf[i]) {
			return errors.NewArg("найден запрещенный символ почтового адреса",
				"forbidden symbol in mail")
		}
		if buf[i] == '@' {
			doggyCount++
			if i > 0 && buf[i-1] == '.' {
				return errors.NewArg("почтовый адрес невалиден", "invalid mail address")
			}
		}
		if buf[i] == '.' && doggyCount > 0 {
			dots++
			if buf[i-1] == '.' || buf[i-1] == '@' {
				return errors.NewArg("почтовый адрес невалиден", "invalid mail address")
			}
		}
	}
	if doggyCount != 1 {
		return errors.NewArg("невалидное количество символов '@' в почтовом адресе",
			"invalid amount of '@' symbols in mail address")
	}
	if dots != 1 && dots != 2 {
		return errors.NewArg("невалидное количество символов '.' в почтовом адресе",
			"invalid amount of '.' symbols in mail address")
	}
	return nil
}

func CheckName(name string) error {
	var runeSlice = []rune(name)

	if len(name) > config.NAME_MAX_LEN {
		return errors.NewArg("слишком длинное поле имени", "too long name length")
	}
	if utf8.RuneCountInString(name) < 1 {
		return errors.NewArg("поле имени пустое", "name is empty")
	}
	if !isLetter(runeSlice[0]) {
		return errors.NewArg("первый символ поля имени должен быть буквой", "first name symbol should be letter")
	}
	if !isLetter(runeSlice[(len(runeSlice)-1)]) && !isDigit(runeSlice[(len(runeSlice)-1)]) {
		return errors.NewArg("последний символ поля имени должен быть буквой или цифрой",
			"last name symbol should be letter or digit")
	}
	for i := 0; i < len(runeSlice); i++ {
		if !isNameRunePermitted(runeSlice[i]) {
			return errors.NewArg("символ поля имени "+string(runeSlice[i])+" запрещен",
				"name letter '"+string(runeSlice[i])+"' is not permitted")
		}
	}
	return nil
}

func CheckGender(gender string) error {
	if gender != "male" && gender != "female" {
		return errors.NewArg("гендер "+gender+" неизвестен", "gender '"+gender+"' is not known")
	}
	return nil
}

func CheckOrientation(orientation string) error {
	if orientation != "hetero" && orientation != "bi" && orientation != "homo" {
		return errors.NewArg("ориентация "+orientation+" неизвестна", "orientation '"+orientation+"' not exist in database")
	}
	return nil
}

func CheckBio(bio string) error {
	if len(bio) > config.BIO_MAX_LEN {
		return errors.NewArg("слишком длинная биография", "too long biography length")
	}
	return nil
}

func FindUnknownInterests(knownInterests []common.Interest, interestsNameArr []string) []common.Interest {
	var unknownInterests []common.Interest
	var newInterest common.Interest
	var isKnown bool
	for _, interestsName := range interestsNameArr {
		isKnown = false
		for _, interest := range knownInterests {
			if interest.Name == interestsName {
				isKnown = true
			}
		}
		if !isKnown {
			newInterest.Name = interestsName
			unknownInterests = append(unknownInterests, newInterest)
		}
	}
	return unknownInterests
}

func CheckInterest(interest string) error {
	if interest == "" {
		return errors.NewArg("поле интересов пустое", "empty interest")
	}
	if len(interest) > config.INTEREST_MAX_LEN {
		return errors.NewArg("поле интересов слишком длинное", "too big interest length")
	}
	return nil
}

func PassHash(pass string) string {
	pass += passSalt
	crcH := crc32.ChecksumIEEE([]byte(pass))
	return strconv.FormatUint(uint64(crcH), 20)
}

func TokenWebSocketAuth(uid int) string {

	str := strconv.Itoa(uid)
	curTime := time.Now()
	dataToHash := str + curTime.Format(time.RFC3339Nano)
	tmpHash := crc32.ChecksumIEEE([]byte(dataToHash))
	hash := strconv.FormatUint(uint64(tmpHash), 35)
	token := string(hash)

	dataToHash = curTime.Format(time.RFC3339Nano)
	tmpHash = crc32.ChecksumIEEE([]byte(dataToHash))
	hash = strconv.FormatUint(uint64(tmpHash), 35)
	token += string(hash)

	return token
}

func TokenUidEncode(uid int) (string, error) {

	// Thanks to https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/
	// for good explanation of Encoding with masterKey
	// AES - Advanced Encryption Standard

	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(masterKey))
	// if there are any errors, handle them
	if err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	token := gcm.Seal(nonce, nonce, []byte(strconv.Itoa(uid)), nil)
	return base64.URLEncoding.EncodeToString(token), nil
}

func TokenMailEncode(mail string) (string, error) {
	c, err := aes.NewCipher([]byte(masterKey))
	if err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.NewArg("ошибка кодирования токена", "token encode error").AddOriginalError(err)
	}
	token := gcm.Seal(nonce, nonce, []byte(mail), nil)
	return base64.URLEncoding.EncodeToString(token), nil
}

func TokenUidDecode(token string) (int, error) {
	encodedToken, _ := base64.URLEncoding.DecodeString(token)

	c, err := aes.NewCipher([]byte(masterKey))
	if err != nil {
		return 0, errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return 0, errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}

	nonceSize := gcm.NonceSize()
	if len(encodedToken) < nonceSize {
		return 0, errors.NewArg("ошибка размера при декодировании токена", "size error in decoding")
	}

	nonce, encodedToken := encodedToken[:nonceSize], encodedToken[nonceSize:]
	desired, err := gcm.Open(nil, nonce, encodedToken, nil)
	if err != nil {
		return 0, err
	}
	uid, err := strconv.Atoi(string(desired))
	if err != nil {
		return 0, errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}
	return uid, nil
}

func TokenMailDecode(token string) (string, error) {
	encodedToken, _ := base64.URLEncoding.DecodeString(token)
	c, err := aes.NewCipher([]byte(masterKey))
	if err != nil {
		return "", errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}
	nonceSize := gcm.NonceSize()
	if len(encodedToken) < nonceSize {
		return "", errors.NewArg("ошибка размера при декодировании токена", "size error in decoding")
	}
	nonce, encodedToken := encodedToken[:nonceSize], encodedToken[nonceSize:]
	desired, err := gcm.Open(nil, nonce, encodedToken, nil)
	if err != nil {
		return "", errors.NewArg("ошибка декодирования токена", "token decode error").AddOriginalError(err)
	}
	mail := string(desired)
	return mail, nil
}

func SendMail(to string, xRegToken string, mailConf *config.Mail) error {
	auth := smtp.PlainAuth("", mailConf.Mail, mailConf.Pass, mailConf.Host)
	message := `To: <` + to + `>
From: "Matcha administration" <` + mailConf.Mail + `>
Subject: Confirm registration in Matcha

Hello, ` + to + `, I have registration code for you!
` + xRegToken + `
`

	if err := smtp.SendMail(mailConf.Host+":587",
		auth, mailConf.Mail, []string{to}, []byte(message)); err != nil {
		return errors.NewArg("не смог отправить письмо", "could not sent letter").AddOriginalError(err)
	}
	return nil
}
