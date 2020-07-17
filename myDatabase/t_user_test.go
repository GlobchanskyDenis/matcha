package myDatabase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"encoding/json"
	"strconv"
	. "MatchaServer/config"
)

var (
	conn = ConnDB{}
	token string

	mail = "user_mail@gmail.com"
	passwd = "AsdVar34!A"

	mailNew = "newUser@mail.com"
	passwdNew = "DFe2*FDsd"
	fnameNew = "Денис"
	lnameNew = "Глобчанский"
	ageNew = 21
	genderNew = "male"
	orientationNew = "getero"
	biographyNew = "born, suffered, died"
	avaPhotoIDNew = 42

	mailFail = "mail@gmail@yandex.ru"
	passwdFail = "12345678"
	fnameFail = "@Денис"
	lnameFail = "qweкий   "
	ageFail = 217
	genderFail = "thing"
	orientationFail = "люблю всех"
	biographyFail = `фвыфв ывфывфщзшзщольджук  йлофыдлвоы фыдлвоыдвлффды дл 
	ывофыдлвоыфлдвоы оыфво фылдво л ыовлывфвфыовфыд офыл офвд лфывыфлво фв флдв офлвдофы лфо фдылов`
	avaPhotoIDFail = -1
)

func TestRegUser(t *testing.T) {

	conn.Connect()

	requestData := strings.NewReader(`{"mail":"`+mail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusCreated
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestAuthUser(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"mail":"`+mail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus := http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: decoding response body error: %s, response body %s" + NO_COLOR + "\n", err.Error(), w.Body)
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
	requiredStatus := http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwdNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"fname":"`+fnameNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"lname":"`+lnameNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"age":`+strconv.Itoa(ageNew)+`}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"gender":"`+genderNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"orientation":"`+orientationNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"biography":"`+biographyNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")

	////////////// UPDATE //////////////////
	requestData = strings.NewReader(`{"avaPhotoID":`+strconv.Itoa(avaPhotoIDNew)+`}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
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
	requiredStatus := http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestCreateUserForFailTests(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// USER CREATE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwd+`","mail":"`+mail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusCreated
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")

	////////////// USER AUTH //////////////////
	requestData = strings.NewReader(`{"mail":"`+mail+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus = http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	var response map[string]interface{}
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: decoding response body error: %s, response body %s" + NO_COLOR + "\n", err.Error(), w.Body)
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
	fmt.Print("TESTS FOR FAIL. IF YOU SEE RED COLOR IN LOGS - ITS ALL RIGHT!!!" + NO_COLOR + "\n")

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"mail":"`+mailFail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":"`+mailFail+`","passwd":"`+passwdFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":"`+mail+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusNotAcceptable
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailRegUser_NotCompleteForms(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":"`+mail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":"","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":"`+mail+`","passwd":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	// ////////////// FAIL REGISTRATION //////////////////
	// requestData = strings.NewReader(`{"fname":""}`)
	// url = "http://localhost:3000/user/"
	// r = httptest.NewRequest("POST", url, requestData)
	// w = httptest.NewRecorder()
	// conn.HttpHandlerUser(w, r)
	// requiredStatus = http.StatusBadRequest
	// if w.Code != requiredStatus {
	// 	t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
	// 	return
	// }

	// ////////////// FAIL REGISTRATION //////////////////
	// requestData = strings.NewReader(`{"lname":""}`)
	// url = "http://localhost:3000/user/"
	// r = httptest.NewRequest("POST", url, requestData)
	// w = httptest.NewRecorder()
	// conn.HttpHandlerUser(w, r)
	// requiredStatus = http.StatusBadRequest
	// if w.Code != requiredStatus {
	// 	t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
	// 	return
	// }

	// ////////////// FAIL REGISTRATION //////////////////
	// requestData = strings.NewReader(`{"gender":""}`)
	// url = "http://localhost:3000/user/"
	// r = httptest.NewRequest("POST", url, requestData)
	// w = httptest.NewRecorder()
	// conn.HttpHandlerUser(w, r)
	// requiredStatus = http.StatusBadRequest
	// if w.Code != requiredStatus {
	// 	t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
	// 	return
	// }

	// ////////////// FAIL REGISTRATION //////////////////
	// requestData = strings.NewReader(`{"orientation":""}`)
	// url = "http://localhost:3000/user/"
	// r = httptest.NewRequest("POST", url, requestData)
	// w = httptest.NewRecorder()
	// conn.HttpHandlerUser(w, r)
	// requiredStatus = http.StatusBadRequest
	// if w.Code != requiredStatus {
	// 	t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
	// 	return
	// }

	// ////////////// FAIL REGISTRATION //////////////////
	// requestData = strings.NewReader(`{"biography":""}`)
	// url = "http://localhost:3000/user/"
	// r = httptest.NewRequest("POST", url, requestData)
	// w = httptest.NewRecorder()
	// conn.HttpHandlerUser(w, r)
	// requiredStatus = http.StatusBadRequest
	// if w.Code != requiredStatus {
	// 	t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
	// 	return
	// }
	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailReg_BrokenJson(t *testing.T) {
	fmt.Print("\033[m")

	////////////// FAIL REGISTRATION //////////////////
	requestData := strings.NewReader(`[{"mail":"`+mailNew+`","passwd":"`+passwdNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL REGISTRATION //////////////////
	requestData = strings.NewReader(`{"mail":`+mailNew+`","passwd":"`+passwdNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user creation was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailUpd_InvalidData(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwdFail+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"mail":"`+mailFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"fname":"`+fnameFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"lname":"`+lnameFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"age":`+strconv.Itoa(ageFail)+`}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"gender":"`+genderFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"orientation":"`+orientationFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"biography":"`+biographyFail+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"avaPhotoID":`+strconv.Itoa(avaPhotoIDFail)+`}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"mail":"`+mailNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", "BLAbla")
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusUnauthorized
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailUpd_NotCompliteForms(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`{"mail":"`+mailNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	// r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusUnauthorized
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"mail":"`+mailNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", "ATATAAGSFDKSALDJdssadfrSFASF")
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusUnauthorized
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"fname":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"lname":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"gender":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"orientation":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"asddd":"asdsaddsdds"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailUpd_BrokenJson(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL UPDATE //////////////////
	requestData := strings.NewReader(`[{"mail":"`+mailNew+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL UPDATE //////////////////
	requestData = strings.NewReader(`{"mail":`+mailNew+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("PATCH", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
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
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"passwd":""}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"dasds":"ABA@DAKADAVR"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	// r.Header.Add("x-auth-token", token)
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusUnauthorized
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL DELETE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/user/"
	r = httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", "BLAbla")
	w = httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus = http.StatusUnauthorized
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	
	t.Logf(GREEN_BG + "SUCCESS: user removing was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}


func TestFailAuth(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// FAIL AUTHENTICATE //////////////////
	requestData := strings.NewReader(`{"mail":"`+mailFail+`","passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/auth/"
	r := httptest.NewRequest("POST", url, requestData)
	w := httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus := http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"mail":"","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`{"mail":"`+mail+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	////////////// FAIL AUTHENTICATE //////////////////
	requestData = strings.NewReader(`[{"mail":"`+mail+`","passwd":"`+passwd+`"}`)
	url = "http://localhost:3000/auth/"
	r = httptest.NewRequest("POST", url, requestData)
	w = httptest.NewRecorder()
	conn.HttpHandlerAuth(w, r)
	requiredStatus = http.StatusBadRequest
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS: user authentication was failed as it expected" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}

func TestDelUserAgain(t *testing.T) {
	fmt.Print(NO_COLOR)

	////////////// DELETE //////////////////
	requestData := strings.NewReader(`{"passwd":"`+passwd+`"}`)
	url := "http://localhost:3000/user/"
	r := httptest.NewRequest("DELETE", url, requestData)
	r.Header.Add("x-auth-token", token)
	w := httptest.NewRecorder()
	conn.HttpHandlerUser(w, r)
	requiredStatus := http.StatusOK
	if w.Code != requiredStatus {
		t.Errorf(RED_BG + "ERROR: wrong StatusCode: got %d, expected %d" + NO_COLOR + "\n", w.Code, requiredStatus)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
	fmt.Print(YELLOW)
}
