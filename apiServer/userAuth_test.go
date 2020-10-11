package apiServer

import (
	. "MatchaServer/common"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestUserAuthenticate(t *testing.T) {
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
			name: "invalid mail",
			payload: map[string]interface{}{
				"mail": mailFail,
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			payload: map[string]interface{}{
				"mail": "",
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no mail",
			payload: map[string]interface{}{
				"pass": pass,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid no passwd",
			payload: map[string]interface{}{
				"mail": mail,
			},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid wrong password",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": passNew,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid not existing mail",
			payload: map[string]interface{}{
				"mail": mailNew,
				"pass": pass,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "valid",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": pass,
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx context.Context
				url = "http://localhost:" + strconv.Itoa(server.Port) + "/user/auth/"
				rec = httptest.NewRecorder()
				req *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)

			// start test
			server.UserAuth(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			} else {
				if tc.name != "valid" {
					t_.Logf(GREEN_BG + "SUCCESS: user authentication was failed as it expected" + NO_COLOR)
				} else {
					var response map[string]interface{}
					err := json.NewDecoder(rec.Body).Decode(&response)
					if err != nil {
						t_.Errorf(RED_BG+"ERROR: decoding response body error: %s, response body %s"+NO_COLOR, err.Error(), rec.Body)
						return
					}
					item, isExist := response["x-auth-token"]
					if !isExist {
						t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR)
						return
					}
					_, ok := item.(string)
					if !ok {
						t.Errorf(RED_BG + "ERROR: token have wrong type" + NO_COLOR)
						return
					}
					t_.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR)
				}
			}
		})
	}

	/*
	**	Delete test user. Returning the original state of database. Before deleting user,
	**	I should satisfy constraints and delete all data for this user from other tables
	 */
	t.Run("delete test user", func(t_ *testing.T) {

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

		//	Delete user
		err = server.Db.DeleteUser(user.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
			return
		}
	})
}
