package apiServer

import (
	. "MatchaServer/common"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMessageGet(t *testing.T) {
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
		var id int
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
		user2, err = server.CreateTestUser("user2@gmail.com", pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user3, err = server.CreateTestUser("user3@gmail.com", pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		/*
		**	Creating messages in database
		 */
		id, err = server.Db.SetNewMessage(user1.Uid, user2.Uid, "message from user1 to user2")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		messageIds = append(messageIds, id)
		id, err = server.Db.SetNewMessage(user2.Uid, user1.Uid, "message from user2 to user1")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		messageIds = append(messageIds, id)
		id, err = server.Db.SetNewMessage(user1.Uid, user2.Uid, "re message from user1 to user2")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		messageIds = append(messageIds, id)
		id, err = server.Db.SetNewMessage(user1.Uid, user3.Uid, "message from user1 to user3")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		messageIds = append(messageIds, id)
	})

	/*
	**	Test cases. Main part of testing
	 */
	testCases := []struct {
		name           string
		uid            int
		payload        map[string]interface{}
		expectedAmount int
		expectedStatus int
	}{
		{
			name: "valid - uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user2.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user2.Uid),
			},
			expectedAmount: 3,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user1.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user1.Uid),
			},
			expectedAmount: 3,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - uid#" + strconv.Itoa(user1.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user1.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedAmount: 1,
			expectedStatus: http.StatusOK,
		}, {
			name: "valid - uid#" + strconv.Itoa(user2.Uid) + " and uid#" + strconv.Itoa(user3.Uid),
			uid:  user2.Uid,
			payload: map[string]interface{}{
				"otherUid": float64(user3.Uid),
			},
			expectedAmount: 0,
			expectedStatus: http.StatusOK,
		}, {
			name:           "invalid no otherUid",
			uid:            user2.Uid,
			payload:        map[string]interface{}{},
			expectedAmount: 0,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx      context.Context
				url      = "http://localhost:" + strconv.Itoa(server.Port) + "/message/get/"
				rec      = httptest.NewRecorder()
				response []interface{}
				req      *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)
			ctx = context.WithValue(ctx, "uid", tc.uid)

			// start test
			server.MessageGet(rec, req.WithContext(ctx))
			if rec.Code == tc.expectedStatus && tc.expectedStatus == http.StatusOK {
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t_.Errorf(RED_BG+"ERROR in unmarshal: %s"+NO_COLOR, err.Error())
				}
				fmt.Printf("%T\n", response)
				messageLen := len(response)
				if messageLen == tc.expectedAmount {
					t_.Logf(GREEN_BG+"SUCCESS: message amount #%d status code #%d"+NO_COLOR, messageLen, rec.Code)
				} else {
					t_.Errorf(RED_BG+"ERROR: wrong message amount: got %d, expected %d"+NO_COLOR, messageLen, tc.expectedAmount)
				}
			} else if rec.Code == tc.expectedStatus {
				var errorInterface map[string]interface{}
				err := json.NewDecoder(rec.Body).Decode(&errorInterface)
				if err != nil {
					t_.Errorf(RED_BG+"ERROR in unmarshal: %s"+NO_COLOR, err.Error())
				} else {
					t_.Logf(GREEN_BG+"SUCCESS: error found as it expected - %s"+NO_COLOR, errorInterface)
				}
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

		// Drop user notifications
		err = server.Db.DropUserNotifs(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DropUserNotifs(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DropUserNotifs(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}

		// Drop user ignores
		err = server.Db.DropUserIgnores(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DropUserIgnores(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}
		err = server.Db.DropUserIgnores(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
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

		server.Db.Close()
	})
}
