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
	conn = ConnDB{}
	token string
)

func TestRegUser(t *testing.T) {

	conn.Connect()

	requestData := strings.NewReader(`{"login":"`+login+`","passwd":"`+passwd+`","mail":"user_mail@gmail.com","phone":"8-968-646-0102"}`)
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
	token = (response["x-auth-token"]).(string)
	fmt.Print("\033[33m")
}

func TestUpdUser(t *testing.T) {

	fmt.Print("\033[m")

	////////////// UPDATE //////////////////
	requestData := strings.NewReader(`{"mail":"changed@mail.ru"}`)
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
	requestData = strings.NewReader(`{"phone":"123456789"}`)
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
}
