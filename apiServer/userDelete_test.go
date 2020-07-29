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

func TestUserDelete(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New(fakeSql.New())
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, passwd)
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
				"passwd": passwd,
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
}