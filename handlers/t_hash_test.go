package handlers

import (
	. "MatchaServer/config"
	"testing"
)

func TestPasswdHash_1(t *testing.T) {
	var passwd = "admin"
	var wasError bool
	var hash = PasswdHash(passwd)

	for i := 0; i < len(hash); i++ {
		if (hash[i] < '0' || hash[i] > '9') &&
			(hash[i] < 'a' || hash[i] > 'z') &&
			(hash[i] < 'A' || hash[i] > 'Z') {
			t.Errorf(RED_BG+"FAILED - wrong char '%c' %s"+NO_COLOR+"\n", hash[i], hash)
			wasError = true
		}
	}
	if len(hash) < 4 || len(hash) > 30 {
		t.Errorf(RED_BG+"FAILED - length %s"+NO_COLOR+"\n", hash)
		wasError = true
	}
	if !wasError {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", hash)
	}
}

func TestPasswdHash_2(t *testing.T) {
	var passwd = "admin"
	var hash1 = PasswdHash(passwd)
	var hash2 = PasswdHash(passwd)

	if hash1 != hash2 {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", hash1, hash2)
		return
	}
	t.Logf(GREEN_BG+"SUCCESS: %s %s"+NO_COLOR+"\n", hash1, hash2)
}

func TestTokenWebSocketAuth_1(t *testing.T) {
	// var login = "admin"
	var uid = 1
	var hash = TokenWebSocketAuth(uid)
	var wasError bool

	for i := 0; i < len(hash); i++ {
		if (hash[i] < '0' || hash[i] > '9') &&
			(hash[i] < 'a' || hash[i] > 'z') &&
			(hash[i] < 'A' || hash[i] > 'Z') {
			t.Errorf(RED_BG+"FAILED - wrong char '%c' %s"+NO_COLOR+"\n", hash[i], hash)
			wasError = true
		}
	}
	if len(hash) < 8 || len(hash) > 40 {
		t.Errorf(RED_BG+"FAILED - length %s"+NO_COLOR+"\n", hash)
		wasError = true
	}
	if !wasError {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", hash)
	}
}

func TestTokenHash_2(t *testing.T) {
	// var login = "admin"
	var uid = 1
	var hash1 = TokenWebSocketAuth(uid)
	var hash2 = TokenWebSocketAuth(uid)
	if hash1 == hash2 {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", hash1, hash2)
		return
	}
	t.Logf(GREEN_BG+"SUCCESS: %s %s"+NO_COLOR+"\n", hash1, hash2)
}

func TestTokenEnkoder(t *testing.T) {
	var uid = 1
	var encodedToken string
	var newEncodedToken string
	var err error
	var wasError bool

	encodedToken, err = TokenEncode(uid)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: error was returned at encoding - %s"+NO_COLOR+"\n", err.Error())
		return
	}
	for i := 0; i < len(encodedToken); i++ {
		if (encodedToken[i] < '0' || encodedToken[i] > '9') &&
			(encodedToken[i] < 'a' || encodedToken[i] > 'z') &&
			(encodedToken[i] < 'A' || encodedToken[i] > 'Z') &&
			encodedToken[i] != '-' && encodedToken[i] != '_' && encodedToken[i] != '.' &&
			encodedToken[i] != '!' && encodedToken[i] != '~' && encodedToken[i] != '*' &&
			encodedToken[i] != '\'' && encodedToken[i] != '(' && encodedToken[i] != ')' && encodedToken[i] != '=' {
			t.Errorf(RED_BG+"FAILED - wrong char '%c' %s"+NO_COLOR+"\n", encodedToken[i], encodedToken)
			wasError = true
		}
	}
	if len(encodedToken) < 8 || len(encodedToken) > 100 {
		t.Errorf(RED_BG+"FAILED - length %d %s"+NO_COLOR+"\n", len(encodedToken), encodedToken)
		wasError = true
	}

	newEncodedToken, err = TokenEncode(uid)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: error was returned at encoding - %s"+NO_COLOR+"\n", err.Error())
		return
	}
	if len(newEncodedToken) < 8 || len(newEncodedToken) > 100 {
		t.Errorf(RED_BG+"FAILED - length %d %s"+NO_COLOR+"\n", len(newEncodedToken), newEncodedToken)
		wasError = true
	}

	if encodedToken == newEncodedToken {
		t.Errorf(RED_BG + "FAILED: tokens should not be identical" + NO_COLOR + "\n")
		wasError = true
	}
	if !wasError {
		t.Logf(GREEN_BG+"SUCCESS: %s %s"+NO_COLOR+"\n", encodedToken, newEncodedToken)
	}
}

func TestTokenDecoder(t *testing.T) {
	var uid = 1
	var encodedToken string
	var expectedUid = uid
	var err error
	var wasError bool

	encodedToken, err = TokenEncode(uid)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: error was returned at encoding - %s"+NO_COLOR+"\n", err.Error())
		return
	}
	if len(encodedToken) < 8 || len(encodedToken) > 100 {
		t.Errorf(RED_BG+"FAILED - length %d %s"+NO_COLOR+"\n", len(encodedToken), encodedToken)
		wasError = true
	}

	uid, err = TokenDecode(encodedToken)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: error was returned at decoding - %s"+NO_COLOR+"\n", err.Error())
		return
	}
	if uid != expectedUid {
		t.Errorf(RED_BG+"FAILED: uid after encoding/decoding is spoiled. Expected %d Received %d"+NO_COLOR+"\n", expectedUid, uid)
		wasError = true
	}
	if !wasError {
		t.Logf(GREEN_BG+"SUCCESS: %d %d"+NO_COLOR+"\n", uid, expectedUid)
	}
}
