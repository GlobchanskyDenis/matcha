package httpHandlers

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	// "MatchaServer/database/fakeSql"
	"MatchaServer/database/postgres"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"bytes"
	"strings"
	"testing"
)

var (
	server = Server{}
	token  string

	mail   = "test@gmail.com"
	passwd = "AsdVar34!A"

	mailNew        = "test_new@gmail.com"
	passwdNew      = "DFe2*FDsd"
	fnameNew       = "Денис"
	lnameNew       = "Глобчанский"
	ageNew         = 21
	genderNew      = "male"
	orientationNew = "hetero"
	biographyNew   = `born, suffered, died`
	avaPhotoIDNew  = 42

	mailFail        = "mail@gmail@yandex.ru"
	passwdFail      = "12345678"
	fnameFail       = "@Денис"
	lnameFail       = "qweкий   "
	ageFail         = 217
	genderFail      = "thing"
	orientationFail = "люблю всех"
	biographyFail   = `фвыфв ывфывфщзшзщольджук  йлофыдлвоы фыдлвоыдвлффды дл 
	ывофыдлвоыфлдвоы оыфво фылдво л ыовлывфвфыовфыд офыл офвд лфывыфлво фв флдв офлвдофы лфо фдылов
	sdsadasdsa sadasdasdasd asd asdsadas as asdasdsad as`
	avaPhotoIDFail = -1
)

func TestUserCreate(t *testing.T) {
	print(NO_COLOR)

	err := server.New(&postgres.ConnDB{})
	// err := server.New(&fakeSql.ConnFake{})
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR + "\n")
		return
	}

	t.Run("Create test user", func(t_ *testing.T) {
		user, err := server.Db.SetNewUser(mail, handlers.PasswdHash(passwd))
		if err != nil {
			t.Errorf(RED_BG + "ERROR: SetNewUser returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
		user.Passwd = handlers.PasswdHash(passwd)
		user.AccType = "confirmed"
		err = server.Db.UpdateUser(user)
		if err != nil {
			t.Errorf(RED_BG + "ERROR: UpdateUser returned error - " + err.Error() + NO_COLOR + "\n")
			return
		}
		t.Logf(GREEN_BG + "SUCCESS: user confirmed its mail" + NO_COLOR + "\n")
	})

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid mail",
			payload: map[string]string{
				"mail":   mailFail,
				"passwd": passwd,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			payload: map[string]string{
				"mail":   mail,
				"passwd": passwdFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "same user already exists",
			payload: map[string]string{
				"mail":   mail,
				"passwd": passwd,
			},
			expectedStatus: http.StatusNotAcceptable,
		}, {
			name: "password not exists",
			payload: map[string]string{
				"mail": mail,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "mail not exists",
			payload: map[string]string{
				"passwd": passwd,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "mail is empty",
			payload: map[string]string{
				"mail":   "",
				"passwd": passwd,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "passwd is empty",
			payload: map[string]string{
				"mail":   mail,
				"passwd": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name:           "broken json 1",
			requestBody:    strings.NewReader(`[{"mail":"` + mailNew + `","passwd":"` + passwdNew + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name:           "broken json 2",
			requestBody:    strings.NewReader(`{"mail":` + mailNew + `","passwd":"` + passwdNew + `"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var req *http.Request
			var url = "http://localhost:3000/user/create/"
			var rec = httptest.NewRecorder()
			if tc.requestBody == nil {
				requestBody := &bytes.Buffer{}
				json.NewEncoder(requestBody).Encode(tc.payload)
				req = httptest.NewRequest("POST", url, requestBody)
			} else {
				req = httptest.NewRequest("POST", url, tc.requestBody)
			}
			server.HttpHandlerUserReg(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user create was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
			}
		})
	}

	print(YELLOW)
}

func TestUserAuthenticate(t *testing.T) {
	print(NO_COLOR)

	testCases := []struct {
		name           string
		requestBody		*strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid mail",
			requestBody: strings.NewReader(`{"mail":"` + mailFail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwdFail + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":""}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			requestBody: strings.NewReader(`{"mail":"","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no mail",
			requestBody: strings.NewReader(`{"passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid no passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`[{"mail":"` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`{"mail":` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid wrong password",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwdNew + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid not existing mail",
			requestBody: strings.NewReader(`{"mail":"` + mailNew + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "valid",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			url := "http://localhost:3000/user/auth/"
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", url, tc.requestBody)
			server.HttpHandlerUserAuth(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else {
				if tc.name != "valid" {
					t.Logf(GREEN_BG + "SUCCESS: user authentication was failed as it expected" + NO_COLOR + "\n")
				} else {
					var response map[string]interface{}
					err := json.NewDecoder(rec.Body).Decode(&response)
					if err != nil {
						t.Errorf(RED_BG+"ERROR: decoding response body error: %s, response body %s"+NO_COLOR+"\n", err.Error(), rec.Body)
						return
					}
					item, isExist := response["x-auth-token"]
					if !isExist {
						t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR + "\n")
						return
					}
					token = item.(string)
					t_.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
				}
			}
		})
	}
	print(YELLOW)
}

func TestUserUpdate(t *testing.T) {
	print(NO_COLOR)

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		requestHeaderName  string
		requestHeaderValue string
		expectedStatus int
	}{
		{
			name: "valid mail",
			payload: map[string]string{
				"mail": mailNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid passwd",
			payload: map[string]string{
				"passwd": passwdNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid fname",
			payload: map[string]string{
				"fname": fnameNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid lname",
			payload: map[string]string{
				"lname": lnameNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid age",
			payload: map[string]int{
				"age": ageNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid gender",
			payload: map[string]string{
				"gender": genderNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid orientation",
			payload: map[string]string{
				"orientation": orientationNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid biography",
			payload: map[string]string{
				"biography": biographyNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid avaPhotoID",
			payload: map[string]int{
				"avaPhotoID": avaPhotoIDNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		}, {
			name: "invalid mail",
			payload: map[string]string{
				"mail": mailFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			payload: map[string]string{
				"passwd": passwdFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid fname",
			payload: map[string]string{
				"fname": fnameFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid lname",
			payload: map[string]string{
				"lname": lnameFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid age",
			payload: map[string]int{
				"age": ageFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid gender",
			payload: map[string]string{
				"gender": genderFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid orientation",
			payload: map[string]string{
				"orientation": orientationFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid biography",
			payload: map[string]string{
				"biography": biographyFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid avaPhotoID",
			payload: map[string]int{
				"avaPhotoID": avaPhotoIDFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			payload: map[string]string{
				"mail": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			payload: map[string]string{
				"passwd": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty fname",
			payload: map[string]string{
				"fname": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty lname",
			payload: map[string]string{
				"lname": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty gender",
			payload: map[string]string{
				"gender": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty orientation",
			payload: map[string]string{
				"orientation": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid update no usefull fields at all",
			payload: map[string]string{
				"asd": "asddasda",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		},{
			name:               "invalid token",
			payload: map[string]string{
				"fname": fnameNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: "BlaBla",
			expectedStatus:     http.StatusUnauthorized,
		}, {
			name:               "invalid token not exists",
			payload: map[string]string{
				"fname": fnameNew,
			},
			requestHeaderName:  "BlaBla",
			requestHeaderValue: "token",
			expectedStatus:     http.StatusUnauthorized,
		}, {
			name:           "invalid broken json",
			requestBody:    strings.NewReader(`[{"mail":"` + mailNew + `"}`),
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		}, {
			name:           "invalid broken json",
			requestBody:    strings.NewReader(`{"mail":` + mailNew + `"}`),
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var req *http.Request
			var url = "http://localhost:3000/user/update/"
			var rec = httptest.NewRecorder()
			if tc.requestBody == nil {
				requestBody := &bytes.Buffer{}
				json.NewEncoder(requestBody).Encode(tc.payload)
				req = httptest.NewRequest("PATCH", url, requestBody)
			} else {
				req = httptest.NewRequest("PATCH", url, tc.requestBody)
			}
			req.Header.Add(tc.requestHeaderName, tc.requestHeaderValue)
			server.HttpHandlerUserUpdate(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
			}
		})
	}

	t.Run("invalid without request body at all", func(t_ *testing.T) {
		url := "http://localhost:3000/user/update/"
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("PATCH", url, nil)
		req.Header.Add("x-auth-token", token)
		server.HttpHandlerUserUpdate(rec, req)
		expectedStatus := http.StatusBadRequest
		if rec.Code != expectedStatus {
			t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, expectedStatus)
		} else {
			t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
		}
	})

	print(YELLOW)
}

func TestUserDelete(t *testing.T) {
	print(NO_COLOR)

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		requestHeaderName  string
		requestHeaderValue string
		expectedStatus int
	}{
		{
			name: "invalid passwd",
			payload: map[string]string{
				"passwd": passwdFail,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			payload: map[string]string{
				"passwd": "",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no useful fields at all",
			payload: map[string]string{
				"Abrakadabra": "asdsad",
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid no fields at all",
			payload: map[string]string{},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid token",
			payload: map[string]string{
				"passwd": passwd,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: "token123",
			expectedStatus: http.StatusUnauthorized,
		}, {
			name: "invalid empty token",
			payload: map[string]string{
				"passwd": passwd,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: "token123",
			expectedStatus: http.StatusUnauthorized,
		}, {
			name: "invalid no token at all",
			payload: map[string]string{
				"passwd": passwd,
			},
			requestHeaderName:  "tokkkken",
			requestHeaderValue: "token321",
			expectedStatus: http.StatusUnauthorized,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`{"passwd":` + passwdNew + `"}`),
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`[{"passwd":"` + passwdNew + `"}`),
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "valid",
			payload: map[string]string{
				"passwd": passwdNew,
			},
			requestHeaderName:  "x-auth-token",
			requestHeaderValue: token,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var req *http.Request
			var url = "http://localhost:3000/user/delete/"
			var rec = httptest.NewRecorder()
			if tc.requestBody == nil {
				requestBody := &bytes.Buffer{}
				json.NewEncoder(requestBody).Encode(tc.payload)
				req = httptest.NewRequest("DELETE", url, requestBody)
			} else {
				req = httptest.NewRequest("DELETE", url, tc.requestBody)
			}
			req.Header.Add(tc.requestHeaderName, tc.requestHeaderValue)
			server.HttpHandlerUserDelete(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if rec.Code != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user removing was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
			}
		})
	}
	
	print(YELLOW)
}
