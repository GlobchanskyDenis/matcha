package apiServer

import (
	. "MatchaServer/config"
	// "MatchaServer/database/fakeSql"

	"MatchaServer/database/postgres"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserCreate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New(postgres.New())
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, pass)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid mail",
			payload: map[string]string{
				"mail": mailFail,
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid password",
			payload: map[string]string{
				"mail": mail,
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "same user already exists",
			payload: map[string]string{
				"mail": mail,
				"pass": pass,
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
				"pass": pass,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "mail is empty",
			payload: map[string]string{
				"mail": "",
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "password is empty",
			payload: map[string]string{
				"mail": mail,
				"pass": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name:           "broken json 1",
			requestBody:    strings.NewReader(`[{"mail":"` + mailNew + `","pass":"` + passNew + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name:           "broken json 2",
			requestBody:    strings.NewReader(`{"mail":` + mailNew + `","pass":"` + passNew + `"}`),
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
			server.HandlerUserCreate(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user create was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
			}
		})
	}
}
