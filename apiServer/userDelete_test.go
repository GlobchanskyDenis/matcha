package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"context"
)

func TestUserDelete(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New("../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, pass)
	token := server.TestTestUserAuthorize(t, testUser)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid password",
			payload: map[string]string{
				"pass":         passFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty password",
			payload: map[string]string{
				"pass":         "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no useful fields at all",
			payload: map[string]string{
				"Abrakadabra":  "asdsad",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid only token in body",
			payload: map[string]string{
				"x-auth-token": token,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "valid",
			payload: map[string]string{
				"pass":         pass,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				requestParams     map[string]interface{}
				err error
				ctx		context.Context
				url = "http://localhost:3000/user/create/"
				rec = httptest.NewRecorder()
				req *http.Request
			)
			// put request body into req
			if tc.requestBody == nil {
				requestBody := &bytes.Buffer{}
				json.NewEncoder(requestBody).Encode(tc.payload)
				req = httptest.NewRequest("DELETE", url, requestBody)
			} else {
				req = httptest.NewRequest("DELETE", url, tc.requestBody)
			}
			// get params from request
			err = json.NewDecoder(req.Body).Decode(&requestParams)
			if err != nil {
				t_.Errorf(RED_BG+"Cannot start test because of error: "+ err.Error() + NO_COLOR+"\n")
				return
			}
			// put params in context
			ctx = context.WithValue(req.Context(), "requestParams", requestParams)

			item, isExist := requestParams["x-auth-token"]
			if !isExist {
				t_.Errorf(RED_BG+"Cannot start test: token expected"+ NO_COLOR+"\n")
				return
			}

			token, ok := item.(string)
			if !ok {
				t_.Errorf(RED_BG+"Cannot start test: token has wrong type"+ NO_COLOR+"\n")
				return
			}

			if token == "" {
				t_.Errorf(RED_BG+"Cannot start test: token is empty"+ NO_COLOR+"\n")
				return
			}

			uid, err := handlers.TokenUidDecode(token)
			if err != nil {
				t_.Errorf(RED_BG+"Cannot start test because of error: "+ err.Error() + NO_COLOR+"\n")
				return
			}

			isLogged := server.session.IsUserLoggedByUid(uid)
			if !isLogged {
				t_.Errorf(RED_BG+"Cannot start test: token is empty"+ NO_COLOR+"\n")
				return
			}

			ctx = context.WithValue(ctx, "uid", uid)
			server.UserDelete(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if rec.Code != http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: user removing was failed as it expected" + NO_COLOR + "\n")
			} else {
				t_.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
			}
		})
	}
}
