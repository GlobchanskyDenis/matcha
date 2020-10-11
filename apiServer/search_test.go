package apiServer

import (
	. "MatchaServer/common"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSearch(t *testing.T) {
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
		user.Longitude = 21.0
		user.Latitude = 42.0
		user.Gender = "female"
		user.Orientation = ""
		err = server.Db.UpdateUser(user)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
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
		expectedStatus int
	}{
		{
			name: "valid - radius",
			payload: map[string]interface{}{
				"radius": map[string]interface{}{
					"radius":    111.0,
					"latitude":  23.0,
					"longitude": 52.0,
				},
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - online",
			payload: map[string]interface{}{
				"online": map[string]interface{}{},
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - age",
			payload: map[string]interface{}{
				"age": map[string]interface{}{
					"min": 17.0,
					"max": 38.0,
				},
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - interests",
			payload: map[string]interface{}{
				"interests": []interface{}{
					"starcraft",
					"football",
				},
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "invalid",
			payload: map[string]interface{}{
				"radius": map[string]interface{}{
					"radius":    -111.0,
					"latitude":  23.0,
					"longitude": 52.0,
				},
			},
			expectedStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx context.Context
				url = "http://localhost:" + strconv.Itoa(server.Port) + "/search/"
				rec = httptest.NewRecorder()
				req *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)
			ctx = context.WithValue(ctx, "uid", user.Uid)

			// start test
			server.Search(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			} else if rec.Code != http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: search was failed as it expected" + NO_COLOR)
			} else {
				t_.Logf(GREEN_BG + "SUCCESS: search is done" + NO_COLOR)
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
