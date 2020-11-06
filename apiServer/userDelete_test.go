package apiServer

import (
	. "MatchaServer/common"
	// "MatchaServer/handlers"
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

	var (
		server *Server
		user   User
	)

	/*
	**	Initialize server and test user
	 */
	t.Run("Initialize", func(t_ *testing.T) {
		var err error
		server, err = New("../config/")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user, err = server.CreateTestUser(mail, pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		err = server.AuthorizeTestUser(user)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot authorize test user - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
	})

	/*
	**	Test cases. Main part of testing
	 */
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
			ctx = context.WithValue(ctx, "uid", user.Uid)

			// start test
			server.UserDelete(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)

				/*
				**	Delete test user. In case of test case that should be valid is failed
				 */
				if tc.expectedStatus == http.StatusOK {
					t_.Run("delete test user after fail", func(t_ *testing.T) {
						//	Delete devices of test user
						devices, err := server.Db.GetDevicesByUid(user.Uid)
						if err != nil {
							t_.Errorf(RED_BG + "Error: cannot get devices of user that i trying to delete - " + err.Error() + NO_COLOR)
							return
						}
						for _, device := range devices {
							err = server.Db.DeleteDevice(device.Id)
							if err != nil {
								t_.Errorf(RED_BG + "Error: cannot delete device of user - " + err.Error() + NO_COLOR)
								return
							}
						}
						// Drop user ignores
						err = server.Db.DropUserIgnores(user.Uid)
						if err != nil {
							t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
						}
						//	Delete user
						err = server.Db.DeleteUser(user.Uid)
						if err != nil {
							t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
							return
						}
					})
				} else if rec.Code != http.StatusOK {
					t_.Logf(GREEN_BG + "SUCCESS: user removing was failed as it expected" + NO_COLOR + "\n")
				} else {
					t_.Logf(GREEN_BG + "SUCCESS: user was removed" + NO_COLOR + "\n")
				}
			}
		})
	}
}
