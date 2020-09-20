package apiServer

import (
	. "MatchaServer/common"
	"testing"
	"net/http"
	"context"
	"strconv"
	"net/http/httptest"
)

func TestSearch(t *testing.T) {
	/////////// INITIALIZE ///////////

	server, err := New("../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, pass)
	defer server.Db.DeleteUser(testUser.Uid)
	testUser.Longitude = 21.0
	testUser.Latitude = 42.0
	testUser.Gender = "female"
	testUser.Orientation = ""
	err = server.Db.UpdateUser(testUser)
	if err != nil {
			t.Errorf(RED_BG + "Cannot start test - token error: " + err.Error() + NO_COLOR + "\n")
			return
		}
	_ = server.TestTestUserAuthorize(t, testUser)
	uid := testUser.Uid

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        map[string]interface{}
		expectedStatus int
	}{
		{
			name: "valid - radius",
			payload: map[string]interface{}{
				"radius": map[string]interface{}{
					"radius": 111.0,
					"latitude": 23.0,
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
				"radius": -111.0,
				"latitude": 23.0,
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
			ctx = context.WithValue(ctx, "uid", uid)

			// start test
			server.Search(rec, req.WithContext(ctx))
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if rec.Code != http.StatusOK {
				t_.Logf(GREEN_BG + "SUCCESS: search was failed as it expected" + NO_COLOR + "\n")
			} else {
				t_.Logf(GREEN_BG + "SUCCESS: search is done" + NO_COLOR + "\n")
			}
		})
	}
}