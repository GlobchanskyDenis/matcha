package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/handlers"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
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
	defer server.Db.DeleteUser(testUser.Uid)
	token := server.TestTestUserAuthorize(t, testUser)
	uid, err := handlers.TokenUidDecode(token)
	if err != nil {
		t.Errorf(RED_BG + "Cannot start test - token error: " + err.Error() + NO_COLOR + "\n")
		return
	}

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        map[string]interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid password",
			payload: map[string]interface{}{
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty password",
			payload: map[string]interface{}{
				"pass": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no useful fields at all",
			payload: map[string]interface{}{
				"Abrakadabra": "asdsad",
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name:           "invalid only token in body",
			payload:        map[string]interface{}{},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "valid",
			payload: map[string]interface{}{
				"pass": pass,
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx context.Context
				url = "http://localhost:" + strconv.Itoa(server.Port) + "/user/delete/"
				rec = httptest.NewRecorder()
				req *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("DELETE", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)
			ctx = context.WithValue(ctx, "uid", uid)

			// start test
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
