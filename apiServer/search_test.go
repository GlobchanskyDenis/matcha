package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		server  *Server
		myUser  User
		user1   User
		user2   User
		photoID int
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
		myUser, err = server.CreateTestUser(mail, pass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		photoID, err = server.Db.SetNewPhoto(myUser.Uid, "test photo src")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot create test photo - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		longitudeMyUser := 21.0
		latitudeMyUser := 44.0
		myUser.Longitude = &longitudeMyUser
		myUser.Latitude = &latitudeMyUser
		myUser.AvaID = &photoID
		myUser.Fname = "Fname"
		myUser.Lname = "Lname"
		myUser.Gender = "female"
		myUser.Orientation = ""
		err = server.Db.UpdateUser(myUser)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user1, err = server.CreateTestUser("testUser1@gmail.com", "pass")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		longitude1 := 22.0
		latitude1 := 42.0
		user1.Longitude = &longitude1
		user1.Latitude = &latitude1
		user1.AvaID = &photoID
		user1.Fname = "Fname"
		user1.Lname = "Lname"
		date, err := time.Parse("2006-01-02", "1989-10-23")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot parse time - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user1.Birth.Time = &date
		user1.Gender = "male"
		user1.Orientation = "hetero"
		err = server.Db.UpdateUser(user1)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user2, err = server.CreateTestUser("testUser2@gmail.com", "pass")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		longitude2 := 21.0
		latitude2 := 42.0
		user2.Longitude = &longitude2
		user2.Latitude = &latitude2
		user2.AvaID = &photoID
		user2.Fname = "Fname"
		user2.Lname = "Lname"
		user2.Gender = "male"
		user2.Orientation = "homo"
		err = server.Db.UpdateUser(user2)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		err = server.AuthorizeTestUser(myUser)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot authorize test user - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		err = server.Db.SetNewLike(myUser.Uid, 1)
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
		expectedAmount int
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
			expectedAmount: 0,
		}, {
			name: "valid - online",
			payload: map[string]interface{}{
				"online": map[string]interface{}{},
			},
			expectedStatus: http.StatusOK,
			expectedAmount: 0,
		}, {
			name: "valid - age",
			payload: map[string]interface{}{
				"age": map[string]interface{}{
					"min": 17.0,
					"max": 38.0,
				},
			},
			expectedStatus: http.StatusOK,
			expectedAmount: 1,
		}, {
			name: "valid - interests",
			payload: map[string]interface{}{
				"interests": []interface{}{
					"starcraft",
					"football",
				},
			},
			expectedStatus: http.StatusOK,
			expectedAmount: 0,
		}, {
			name: "invalid",
			payload: map[string]interface{}{
				"radius": map[string]interface{}{
					"radius":    -111.0,
					"latitude":  23.0,
					"longitude": 52.0,
				},
			},
			expectedStatus: errors.InvalidArgument.HttpResponseStatus, //http.StatusUnprocessableEntity,
			expectedAmount: 0,
		}, {
			name: "valid - search users that wasnt liked",
			payload: map[string]interface{}{
				"wasntLiked": map[string]interface{}{},
			},
			expectedStatus: http.StatusOK,
			expectedAmount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var (
				ctx      context.Context
				url      = "http://localhost:" + strconv.Itoa(server.Port) + "/search/"
				rec      = httptest.NewRecorder()
				req      *http.Request
				response []interface{}
			)
			// all request params should be handled in middlewares
			// so new request body is nil
			req = httptest.NewRequest("POST", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)
			ctx = context.WithValue(ctx, "uid", myUser.Uid)

			// start test
			server.Search(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR, rec.Code, tc.expectedStatus)
			} else if rec.Code != http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: search was failed as it expected" + NO_COLOR)
			} else {
				err := json.NewDecoder(rec.Body).Decode(&response)
				if err != nil {
					t_.Errorf(RED_BG+"ERROR in unmarshal: %s"+NO_COLOR, err.Error())
				}
				fmt.Printf("%#v\n", response)
				usersAmount := len(response)
				if usersAmount == tc.expectedAmount {
					t_.Logf(GREEN_BG+"SUCCESS: users amount #%d status code #%d"+NO_COLOR, usersAmount, rec.Code)
				} else {
					t_.Errorf(RED_BG+"ERROR: wrong users amount: got %d, expected %d"+NO_COLOR, usersAmount, tc.expectedAmount)
				}
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
		devices, err := server.Db.GetDevicesByUid(myUser.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot get devices of user that i trying to delete - " + err.Error() + NO_COLOR)
			// return
		}
		for _, device := range devices {
			err = server.Db.DeleteDevice(device.Id)
			if err != nil {
				t_.Errorf(RED_BG + "Error: cannot delete device of user - " + err.Error() + NO_COLOR)
				// return
			}
		}

		// delete test photo
		err = server.Db.DeletePhoto(photoID)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete test photo - " + err.Error() + NO_COLOR)
		}

		//	Unset like
		err = server.Db.UnsetLike(myUser.Uid, 1)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
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
		err = server.Db.DropUserNotifs(myUser.Uid)
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
		err = server.Db.DropUserIgnores(myUser.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot drop user ignores - " + err.Error() + NO_COLOR)
		}

		//	Delete user
		err = server.Db.DeleteUser(myUser.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
			// return
		}
		err = server.Db.DeleteUser(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
			// return
		}
		err = server.Db.DeleteUser(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "Error: cannot delete user - " + err.Error() + NO_COLOR)
			// return
		}

		server.Db.Close()
	})
}
