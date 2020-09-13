package apiServer

import (
	. "MatchaServer/common"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUserCreate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New("../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, pass)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        map[string]interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid mail",
			payload: map[string]interface{}{
				"mail": mailFail,
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid password",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "same user already exists",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": pass,
			},
			expectedStatus: http.StatusNotAcceptable,
		}, {
			name: "password not exists",
			payload: map[string]interface{}{
				"mail": mail,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "mail not exists",
			payload: map[string]interface{}{
				"pass": pass,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "mail is empty",
			payload: map[string]interface{}{
				"mail": "",
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "password is empty",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx context.Context
				url = "http://localhost:" + strconv.Itoa(server.Port) + "/user/create/"
				rec = httptest.NewRecorder()
				req *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)

			// start test
			server.UserCreate(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: user create was failed as it expected" + NO_COLOR + "\n")
			} else {
				t_.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
			}
		})
	}
}
