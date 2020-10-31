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

func TestNotifications(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		server   *Server
		user1    User
		user2    User
		user3    User
		notifIds []int
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
		**	Creating notifications in database
		 */
		id, err = server.Db.SetNewNotif(user1.Uid, user2.Uid, "user1 liked user2")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		notifIds = append(notifIds, id)
		id, err = server.Db.SetNewNotif(user2.Uid, user1.Uid, "user2 liked user1")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		notifIds = append(notifIds, id)
		id, err = server.Db.SetNewNotif(user3.Uid, user1.Uid, "user3 liked user1")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set message - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		notifIds = append(notifIds, id)
	})

	/*
	**	Test cases. Main part of testing
	 */
	testCasesGet := []struct {
		name           string
		uid            int
		expectedAmount int
		expectedStatus int
	}{
		{
			name:           "valid - uid#" + strconv.Itoa(user1.Uid),
			uid:            user1.Uid,
			expectedAmount: 2,
			expectedStatus: http.StatusOK,
		}, {
			name:           "valid - uid#" + strconv.Itoa(user2.Uid),
			uid:            user2.Uid,
			expectedAmount: 1,
			expectedStatus: http.StatusOK,
		}, {
			name:           "valid - uid#" + strconv.Itoa(user3.Uid),
			uid:            user3.Uid,
			expectedAmount: 0,
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCasesGet {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx      context.Context
				url      = "http://localhost:" + strconv.Itoa(server.Port) + "/notification/get/"
				rec      = httptest.NewRecorder()
				response []interface{}
				req      *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(ctx, "uid", tc.uid)

			// start test
			server.NotificationGet(rec, req.WithContext(ctx))
			if rec.Code == tc.expectedStatus && tc.expectedStatus == http.StatusOK {
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t_.Errorf(RED_BG+"ERROR in unmarshal: %s"+NO_COLOR, err.Error())
				}
				fmt.Printf("%#v\n", response)
				notifLen := len(response)
				if notifLen == tc.expectedAmount {
					t_.Logf(GREEN_BG+"SUCCESS: notif amount #%d status code #%d"+NO_COLOR, notifLen, rec.Code)
				} else {
					t_.Errorf(RED_BG+"ERROR: wrong notif amount: got %d, expected %d"+NO_COLOR, notifLen, tc.expectedAmount)
				}
			} else {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			}
		})
	}

	// Delete notifications of our test
	for _, nid := range notifIds {
		t.Run("delete notif #"+strconv.Itoa(nid), func(t_ *testing.T) {
			var (
				ctx           context.Context
				url           = "http://localhost:" + strconv.Itoa(server.Port) + "/notification/delete/"
				rec           = httptest.NewRecorder()
				requestParams map[string]interface{}
				req           *http.Request
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			requestParams = map[string]interface{}{
				"nid": float64(nid),
			}
			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", requestParams)
			// ctx = context.WithValue(ctx, "uid", tc.uid)

			// start test
			server.NotificationDelete(rec, req.WithContext(ctx))
			if rec.Code == http.StatusOK {
				t_.Logf(GREEN_BG+"SUCCESS: notification #%d was deleted"+NO_COLOR, nid)
			} else {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, http.StatusOK)
			}
		})
	}

	t.Run("invalid delete", func(t_ *testing.T) {
		var (
			ctx           context.Context
			url           = "http://localhost:" + strconv.Itoa(server.Port) + "/notification/delete/"
			rec           = httptest.NewRecorder()
			requestParams map[string]interface{}
			req           *http.Request
		)
		// all request params should be handled in middlewares
		// so new request body is nil
		req = httptest.NewRequest("POST", url, nil)

		requestParams = map[string]interface{}{
			"nid": float64(0),
		}
		// put info from middlewares into context
		ctx = context.WithValue(req.Context(), "requestParams", requestParams)
		// ctx = context.WithValue(ctx, "uid", tc.uid)

		// start test
		server.NotificationDelete(rec, req.WithContext(ctx))
		if rec.Code == http.StatusUnprocessableEntity {
			t_.Logf(GREEN_BG+"SUCCESS: notification #%d was not deleted as it expected"+NO_COLOR, 0)
		} else {
			t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, http.StatusUnprocessableEntity)
		}
	})

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
