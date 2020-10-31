package apiServer

import (
	. "MatchaServer/common"
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestLikes(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		server     *Server
		user1      User
		user2      User
		user3      User
		messageIds []int
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
		/*
		**	Creating users
		 */
		user1, err = server.CreateTestUser("user1@gmail.com", pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user1.AvaID = 1
		err = server.Db.UpdateUser(user1)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user2, err = server.CreateTestUser("user2@gmail.com", pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user2.AvaID = 2
		err = server.Db.UpdateUser(user2)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user3, err = server.CreateTestUser("user3@gmail.com", pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user3.AvaID = 3
		err = server.Db.UpdateUser(user3)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
	})

	/*
	**	Test cases. Set likes
	 */
	testCasesSet := []struct {
		name           string
		uid            int
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "valid - set like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user2.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user2.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - set like from uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user1.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user1.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - set like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - set like from uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name:           "invalid no otherUid",
			uid:            user2.Uid,
			payload:        map[string]interface{}{},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid - repeating like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user2.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user2.Uid),
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}

	for _, tc := range testCasesSet {
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
			ctx = context.WithValue(ctx, "uid", tc.uid)

			// start test
			server.LikeSet(rec, req.WithContext(ctx))
			if rec.Code == tc.expectedStatus && tc.expectedStatus == http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: like was set" + NO_COLOR)
			} else if rec.Code == tc.expectedStatus {
				t_.Logf(GREEN_BG + "SUCCESS: test was failed as it expected" + NO_COLOR)
			} else {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			}
		})
	}

	/*
	**	Test cases. Unset likes
	 */
	testCasesUnset := []struct {
		name           string
		uid            int
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "valid - unset like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user2.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user2.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - unset like from uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user1.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user1.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - unset like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - unset like from uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedStatus: http.StatusOK,
		}, {
			name:           "invalid no otherUid",
			uid:            user2.Uid,
			payload:        map[string]interface{}{},
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid - repeating unset like from uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user2.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user2.Uid),
			},
			expectedStatus: http.StatusNotAcceptable,
		},
	}

	for _, tc := range testCasesUnset {
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
			ctx = context.WithValue(ctx, "uid", tc.uid)

			// start test
			server.LikeUnset(rec, req.WithContext(ctx))
			if rec.Code == tc.expectedStatus && tc.expectedStatus == http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: like was set" + NO_COLOR)
			} else if rec.Code == tc.expectedStatus {
				t_.Logf(GREEN_BG + "SUCCESS: test was failed as it expected" + NO_COLOR)
			} else {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			}
		})
	}

	/*
	**	Delete test user. Returning the original state of database. Before deleting user,
	**	I should satisfy constraints and delete all data for this user from other tables
	 */
	t.Run("delete test user", func(t_ *testing.T) {

		//	Delete devices of test user
		devices, err := server.Db.GetDevicesByUid(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot get devices of user that i trying to delete - " + err.Error() + NO_COLOR)
		}
		for _, device := range devices {
			err = server.Db.DeleteDevice(device.Id)
			if err != nil {
				t_.Errorf(RED_BG + "Error: cannot delete device of user - " + err.Error() + NO_COLOR)
			}
		}

		// Delete messages of our test
		for _, mid := range messageIds {
			err = server.Db.DeleteMessage(mid)
			if err != nil {
				t_.Errorf(RED_BG + "Error: cannot delete message - " + err.Error() + NO_COLOR)
			}
		}

		//	Delete user
		err = server.Db.DeleteUser(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DeleteUser(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DeleteUser(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
		}
	})
}
