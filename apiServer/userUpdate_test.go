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

func TestUserUpdate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New("../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, mail)
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
			name: "valid mail",
			payload: map[string]interface{}{
				"mail": mailNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid password",
			payload: map[string]interface{}{
				"pass": passNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid fname",
			payload: map[string]interface{}{
				"fname": fnameNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid lname",
			payload: map[string]interface{}{
				"lname": lnameNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid birth date",
			payload: map[string]interface{}{
				"birth": "1989-10-23",
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid gender",
			payload: map[string]interface{}{
				"gender": genderNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid orientation",
			payload: map[string]interface{}{
				"orientation": orientationNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid biography",
			payload: map[string]interface{}{
				"bio": bioNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid avaPhotoID",
			payload: map[string]interface{}{
				"avaID": avaIDNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid latitude",
			payload: map[string]interface{}{
				"latitude": latitudeNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid longitude",
			payload: map[string]interface{}{
				"longitude": longitudeNew,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid interests #1",
			payload: map[string]interface{}{
				"interests": interests1New,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid interests #2",
			payload: map[string]interface{}{
				"interests": interests2New,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "invalid mail",
			payload: map[string]interface{}{
				"mail": mailFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid password",
			payload: map[string]interface{}{
				"pass": passFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid fname",
			payload: map[string]interface{}{
				"fname": fnameFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid lname",
			payload: map[string]interface{}{
				"lname": lnameFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid birth date",
			payload: map[string]interface{}{
				"birth": "2020-08-23",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid birth date - bad parsing",
			payload: map[string]interface{}{
				"birth": "198910-23",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid gender",
			payload: map[string]interface{}{
				"gender": genderFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid orientation",
			payload: map[string]interface{}{
				"orientation": orientationFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid biography",
			payload: map[string]interface{}{
				"bio": bioFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid avaPhotoID",
			payload: map[string]interface{}{
				"avaID": avaIDFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid latitude",
			payload: map[string]interface{}{
				"latitude": latitudeFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid longitude",
			payload: map[string]interface{}{
				"longitude": longitudeFail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid interests #1",
			payload: map[string]interface{}{
				"interests": interests1Fail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid interests #2",
			payload: map[string]interface{}{
				"interests": interests2Fail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid interests #3",
			payload: map[string]interface{}{
				"interests": interests3Fail,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			payload: map[string]interface{}{
				"mail": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			payload: map[string]interface{}{
				"pass": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty fname",
			payload: map[string]interface{}{
				"fname": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty lname",
			payload: map[string]interface{}{
				"lname": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty gender",
			payload: map[string]interface{}{
				"gender": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty orientation",
			payload: map[string]interface{}{
				"orientation": "",
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid update no usefull fields at all",
			payload: map[string]interface{}{
				"asd": "asddasda",
			},
			expectedStatus: http.StatusBadRequest,
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
			req = httptest.NewRequest("PATCH", url, nil)

			// put info from middlewares into context
			ctx = context.WithValue(req.Context(), "requestParams", tc.payload)
			ctx = context.WithValue(ctx, "uid", uid)

			// start test
			server.UserUpdate(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
			}
		})
	}
}
