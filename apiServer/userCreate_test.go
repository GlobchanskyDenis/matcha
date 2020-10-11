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

func TestUserCreate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var server *Server

	/*
	**	Initialize server
	 */
	t.Run("Initialize", func(t_ *testing.T) {
		var err error
		server, err = New("../config/")
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
			name: "invalid password",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "valid",
			payload: map[string]interface{}{
				"mail": mail,
				"pass": pass,
			},
			expectedStatus: http.StatusCreated,
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
			} else if tc.expectedStatus != http.StatusCreated {
				t_.Logf(GREEN_BG + "SUCCESS: user create was failed as it expected" + NO_COLOR + "\n")
			} else {
				t_.Logf(GREEN_BG + "SUCCESS: user was created" + NO_COLOR + "\n")
			}
			/*
			**	Delete user if it was created
			 */
			if tc.expectedStatus == http.StatusCreated {
				/*
				**	Handle response
				 */
				var response map[string]interface{}
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t_.Errorf(RED_BG+"ERROR: decoding response body error: %s, response body %s. Cannot delete user"+NO_COLOR, err.Error(), rec.Body)
					return
				}
				item, isExist := response["uid"]
				if !isExist {
					t_.Errorf(RED_BG + "ERROR: uid not found in response" + NO_COLOR)
					return
				}
				uid64, ok := item.(float64)
				if !ok {
					t_.Errorf(RED_BG + "ERROR: uid have wrong type" + NO_COLOR)
					return
				}
				uid := int(uid64)
				/*
				**	Delete devices of test user
				 */
				devices, err := server.Db.GetDevicesByUid(uid)
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
				/*
				**	Delete user
				 */
				err = server.Db.DeleteUser(uid)
				if err != nil {
					t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
					return
				}
			}
		})
	}
}
