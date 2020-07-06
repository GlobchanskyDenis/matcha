package handlers

import (
	"time"
	"testing"
)

func TestPasswdHash_1(t *testing.T) {
	var passwd = "admin"
	var wasError bool
	var hash = PasswdHash(passwd)

	for i:=0; i<len(hash); i++ {
		if (hash[i] < '0' || hash[i] > '9') && 
		(hash[i] < 'a' || hash[i] > 'z') && 
		(hash[i] < 'A' || hash[i] > 'Z') {
			t.Errorf("\033[31mFAILED\033[m - wrong char %s\n\n", hash)
			wasError = true
		}
	}
	if len(hash) < 4 || len(hash) > 30 {
		t.Errorf("\033[31mFAILED\033[m - length %s\n\n", hash)
		wasError = true
	}
	if !wasError {
		t.Logf("\033[32mDONE\033[m - %s\n\n", hash)
	}
}

func TestPasswdHash_2(t *testing.T) {
	var passwd = "admin"
	var hash1 = PasswdHash(passwd)
	var hash2 = PasswdHash(passwd)

	if hash1 != hash2 {
		t.Errorf("\033[31mFAILED\033[m %s %s\n\n", hash1, hash2)
	} else {
		t.Logf("\033[32mDONE\033[m - %s %s\n\n", hash1, hash2)
	}
}

func TestTokenHash_1(t *testing.T) {
	var login = "admin"
	var wasError bool
	var time1 = time.Now()
	var hash = TokenHash(login, time1)

	for i:=0; i<len(hash); i++ {
		if (hash[i] < '0' || hash[i] > '9') && 
		(hash[i] < 'a' || hash[i] > 'z') && 
		(hash[i] < 'A' || hash[i] > 'Z') {
			t.Errorf("\033[31mFAILED\033[m - wrong char %s\n\n", hash)
			wasError = true
		}
	}
	if len(hash) < 8 || len(hash) > 40 {
		t.Errorf("\033[31mFAILED\033[m - length %s\n\n", hash)
		wasError = true
	}
	if !wasError {
		t.Logf("\033[32mDONE\033[m - %s\n\n", hash)
	}
}

func TestTokenHash_2(t *testing.T) {
	var login = "admin"
	var time1 = time.Now()
	var time2 = time.Now()
	var hash1 = TokenHash(login, time1)
	var hash2 = TokenHash(login, time2)
	if hash1 == hash2 {
		t.Errorf("\033[31mFAILED\033[m %s %s\n\n", hash1, hash2)
	} else {
		t.Logf("\033[32mDONE\033[m - %s %s\n\n", hash1, hash2)
	}
}

func TestTokenEnkoder(t *testing.T) {
	var login = "admin"
	var encodedToken string
	var newEncodedToken string
	var err Errorf
	var wasError bool
	
	encodedToken, err = TokenEncode(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - error was returned at encoding - %s\n\n", fmt.Sprintf("%s", err))
		wasError = true
	}
	if len(encodedToken) < 8 || len(encodedToken) > 100 {
		t.Errorf("\033[31mFAILED\033[m - length %d %s\n\n", len(encodedToken), encodedToken)
		wasError = true
	}

	newEncodedToken, err = TokenEncode(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - error was returned at encoding - %s\n\n", fmt.Sprintf("%s", err))
		wasError = true
	}
	if len(newEncodedToken) < 8 || len(newEncodedToken) > 100 {
		t.Errorf("\033[31mFAILED\033[m - length %d %s\n\n", len(newEncodedToken), newEncodedToken)
		wasError = true
	}

	if encodedToken == newEncodedToken {
		t.Errorf("\033[31mFAILED\033[m - tokens should not be identical\n\n")
		wasError = true
	}
	if !wasError {
		t.Logf("\033[32mDONE\033[m - %s %s\n\n", encodedToken, newEncodedToken)
	}
}

func TestTokenDecoder(t *testing.T) {
	var login = "admin"
	var encodedToken string
	var expectedLogin string = login
	var err Errorf
	var wasError bool
	
	encodedToken, err = TokenEncode(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - error was returned at encoding - %s\n\n", fmt.Sprintf("%s", err))
		wasError = true
	}
	if len(encodedToken) < 8 || len(encodedToken) > 100 {
		t.Errorf("\033[31mFAILED\033[m - length %d %s\n\n", len(encodedToken), encodedToken)
		wasError = true
	}

	login, err = TokenDecode(encodedToken)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - error was returned at decoding - %s\n\n", fmt.Sprintf("%s", err))
		wasError = true
	}
	if login != expectedLogin {
		t.Errorf("\033[31mFAILED\033[m - login after encoding/decoding is spoiled. Expected %s Received %s\n\n", expectedLogin, login)
		wasError = true
	}
	if !wasError {
		t.Logf("\033[32mDONE\033[m - %s %s\n\n", encodedToken, newEncodedToken)
	}
}
