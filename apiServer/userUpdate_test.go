package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/database/fakeSql"
	// "MatchaServer/database/postgres"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"bytes"
	"strings"
	"testing"
)

func TestUserUpdate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New(fakeSql.New())
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, mail)
	token := server.TestTestUserAuthorize(t, testUser)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

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
}