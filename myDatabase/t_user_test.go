package myDatabase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"encoding/json"
	. "MatchaServer/config"
)

var (
	login = "TestUser"
	loginNew = "newTestUser"
	passwd = "AsdVar34!A"
	passwdNew = "DFe2*FDsd"
	mail = "user_mail@gmail.com"
	mailNew = "newUser@mail.com"
	conn = ConnDB{}
	token string

	loginFail = " AAAA   "
	passwdFail = "12345678"
	mailFail = "mail@gmail@yandex.ru"
)

func TestRegUser(t *testing.T) {

	conn.Connect()

	requestData := strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusCreated)
	}
	t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}



func TestAuthUser(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: decoding response body error: %s" + NO_COLOR + "\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR + "\n")
		return
	}
	token = item.(string)
	t.Logf(GREEN_BG + "SUCCESS: user was authenticated" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}




func TestUpdUser(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// UPDATE //////////////////
	requestData := strings.NewReader(`{"mail":"`+mailNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: error in decoding response body: %s" + NO_COLOR + "\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR + "\n")
		return
	}
	token = item.(string)
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwdNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}



func TestDelUser(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// DELETE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwdNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}





func TestCreateUserForFailTests(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// USER CREATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusCreated)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")

	////////////// USER AUTH //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: decoding response body error -  %s" + NO_COLOR + "\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR + "\n")
		return
	}
	token = item.(string)
	t.Logf(GREEN_BG + "SUCCESS: user was authenticated" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}





func TestFailRegUser_InvalidData(t *testing.T) {
	fmt.Print("TESTS FOR FAIL. IF YOU SEE RED COLOR IN LOGS - ITS ALL RIGHT!!!\n\n" + NO_COLOR)

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwdFail+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mailFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return	
	}

	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}





func TestFailRegUser_NotCompleteForms(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return	
	}

	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}





func TestFailReg_BrokenJson(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`[{"login":"`+loginFail+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":`+login+`","passwd":"`+passwdFail+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: user should not be created - its an error" + NO_COLOR + "\n")
		return	
	}

	
	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}





func TestFailUpd_InvalidData(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return	
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwdFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return	
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"mail":"`+mailFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return	
	}

	t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}




func TestFailUpd_NotCompliteForms(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	// r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", "ATATAAGSFDKSALDJdssadfrSFASF")
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL UPDATE //////////////////
	// requestData := strings.NewReader(`{"login":"`+loginNew+`"}`)
	requestData = strings.NewReader(`{"asddd":"asdsaddsdds"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}



func TestFailUpd_BrokenJson(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`[{"login":"`+loginNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"login":`+loginNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: user should not be updated - its an error" + NO_COLOR + "\n")
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}



func TestFailDelUser(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL DELETE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwdFail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"passwd":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"login":"ABADAKADAVR"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`[{"passwd":"asdsadas"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusInternalServerError)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	// r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}
	
	t.Logf(GREEN_BG + "SUCCESS: user removing was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}

func TestFailAuth(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"login":"","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`[{"login":"`+login+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusInternalServerError)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user authentication was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}

func TestDelUserAgain(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// DELETE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, http.StatusOK)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}
