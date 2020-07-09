package myDatabase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"encoding/json"
)

var (
	login = "TestUser"
	loginNew = "newTestUser"
	passwd = "AsdVar34!A"
	passwdNew = "DFe2*FDsd"
	mail = "user_mail@gmail.com"
	mailNew = "newUser@mail.com"
	phone = "8-968-646-0102"
	phoneNew = "+7(976)456-4567"
	conn = ConnDB{}
	token string

	loginFail = " AAAA   "
	passwdFail = "12345678"
	mailFail = "mail@gmail@yandex.ru"
	phoneFail = "4654)78954--4"
)

func TestRegUser(t *testing.T) {

	conn.Connect()

	requestData := strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusCreated)
	} else {
		t.Logf("\033[32mDONE\033[m - user was created\n")
	}
	fmt.Print("\033[33m")
}



func TestAuthUser(t *testing.T) {

	fmt.Print("\033[m")

	////////////// AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was authenticated\n")
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("\033[31mError\033[m - error in decoding response body: %s\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf("\033[31mError\033[m - token not found in response\n")
		return
	}
	token = item.(string)
	fmt.Print("\033[33m")
}




func TestUpdUser(t *testing.T) {

	fmt.Print("\033[m")

	////////////// UPDATE //////////////////
	requestData := strings.NewReader(`{"mail":"`+mailNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was updated\n")
	}

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"phone":"`+phoneNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was updated\n")
	}

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was updated\n")
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("\033[31mError\033[m - error in decoding response body: %s\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf("\033[31mError\033[m - token not found in response\n")
		return
	}
	token = item.(string)
	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwdNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was updated\n")
	}
	fmt.Print("\033[33m")
}



func TestDelUser(t *testing.T) {

	fmt.Print("\033[m")

	////////////// DELETE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwdNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was removed successfully\n")
	}
	fmt.Print("\033[33m")
}





func TestCreateUserForFailTests(t *testing.T) {
	fmt.Print("\033[m")

	requestData := strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusCreated {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusCreated)
	} else {
		t.Logf("\033[32mDONE\033[m - user was created\n")
	}

	requestData = strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was authenticated\n")
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf("\033[31mError\033[m - error in decoding response body: %s\n", err)
		return
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf("\033[31mError\033[m - token not found in response\n")
		return
	}
	token = item.(string)

	fmt.Print("\033[33m")
}





func TestFailRegUser_InvalidData(t *testing.T) {
	fmt.Print("TESTS FOR FAIL. IF YOU SEE RED COLOR IN LOGS - ITS ALL RIGHT!!!\n\n\033[m")

	var wasNoError bool

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwdFail+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mailFail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phoneFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	if !wasNoError {
		t.Logf("\033[32mDONE\033[m - user creation was failed as it expected\n")
	}
	fmt.Print("\033[33m")
}





func TestFailRegUser_NotCompleteForms(t *testing.T) {
	fmt.Print("\033[m")

	var wasNoError bool

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":"","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code == http.StatusCreated {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		wasNoError = true	
	}

	if !wasNoError {
		t.Logf("\033[32mDONE\033[m - user creation was failed as it expected\n")
	}
	fmt.Print("\033[33m")
}





func TestFailReg_BrokenJson(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`[{"login":"`+loginFail+`","passwd":"`+passwd+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		return	
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"login":`+login+`","passwd":"`+passwdFail+`","mail":"`+mail+`","phone":"`+phone+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("\033[31mError\033[m - user should not be created - its an error\n")
		return	
	}

	t.Logf("\033[32mDONE\033[m - user creation was failed because of broken json - as it expected\n")
	fmt.Print("\033[33m")
}





func TestFailUpd_InvalidData(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
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
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
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
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
		return	
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"phone":"`+phoneFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
		return	
	}

	t.Logf("\033[32mDONE\033[m - user update was failed because of invalid data - as it expected\n")
	fmt.Print("\033[33m")
}




func TestFailUpd_NotCompliteForms(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	// r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
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
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
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
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
		return
	}

	t.Logf("\033[32mDONE\033[m - user update was failed because of invalid data - as it expected\n")
	fmt.Print("\033[33m")
}

func TestFailUpd_BrokenJson(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`[{"login":"`+loginNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
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
		t.Errorf("\033[31mError\033[m - user should not be updated - its an error\n")
		return
	}

	t.Logf("\033[32mDONE\033[m - user update was failed because of invalid data - as it expected\n")
	fmt.Print("\033[33m")
}



func TestFailDelUser(t *testing.T) {

	fmt.Print("\033[m")

	////////////// FAIL DELETE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwdFail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
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
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
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
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
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
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusInternalServerError)
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
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}
	t.Logf("\033[32mDONE\033[m - user was not removed - as it expected\n")
	fmt.Print("\033[33m")
}

func TestFailAuth(t *testing.T) {

	fmt.Print("\033[m")

	////////////// FAIL AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginFail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"login":"","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"login":"`+login+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusNonAuthoritativeInfo {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusNonAuthoritativeInfo)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`[{"login":"`+login+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusInternalServerError)
		return
	}

	t.Logf("\033[32mDONE\033[m - user was not authenticated - as it expected\n")
	fmt.Print("\033[33m")
}

func TestDelUserAgain(t *testing.T) {

	fmt.Print("\033[m")

	////////////// DELETE //////////////////
	requestData := strings.NewReader(`{"login":"`+loginNew+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("\033[31mError\033[m - wrong StatusCode: got %d, expected %d\n", w.Code, http.StatusOK)
		return
	} else {
		t.Logf("\033[32mDONE\033[m - user was removed successfully\n")
	}
	fmt.Print("\033[33m")
}
