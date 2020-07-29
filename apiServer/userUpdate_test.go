package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/database/fakeSql"
	// "MatchaServer/database/postgres"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserUpdate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New(fakeSql.New())
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, mail)
	token := server.TestTestUserAuthorize(t, testUser)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		payload        interface{}
		requestBody    *strings.Reader
		expectedStatus int
	}{
		{
			name: "valid mail",
			payload: map[string]string{
				"mail":         mailNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid passwd",
			payload: map[string]string{
				"passwd":       passwdNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid fname",
			payload: map[string]string{
				"fname":        fnameNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid lname",
			payload: map[string]string{
				"lname":        lnameNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid age",
			payload: map[string]interface{}{
				"age":          ageNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid gender",
			payload: map[string]string{
				"gender":       genderNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid orientation",
			payload: map[string]string{
				"orientation":  orientationNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid biography",
			payload: map[string]string{
				"biography":    biographyNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "valid avaPhotoID",
			payload: map[string]interface{}{
				"avaPhotoID":   avaPhotoIDNew,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusOK,
		}, {
			name: "invalid mail",
			payload: map[string]string{
				"mail":         mailFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			payload: map[string]string{
				"passwd":       passwdFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid fname",
			payload: map[string]string{
				"fname":        fnameFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid lname",
			payload: map[string]string{
				"lname":        lnameFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid age",
			payload: map[string]interface{}{
				"age":          ageFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid gender",
			payload: map[string]string{
				"gender":       genderFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid orientation",
			payload: map[string]string{
				"orientation":  orientationFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid biography",
			payload: map[string]string{
				"biography":    biographyFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid avaPhotoID",
			payload: map[string]interface{}{
				"avaPhotoID":   avaPhotoIDFail,
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			payload: map[string]string{
				"mail":         "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			payload: map[string]string{
				"passwd":       "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty fname",
			payload: map[string]string{
				"fname":        "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty lname",
			payload: map[string]string{
				"lname":        "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty gender",
			payload: map[string]string{
				"gender":       "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty orientation",
			payload: map[string]string{
				"orientation":  "",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid update no usefull fields at all",
			payload: map[string]string{
				"asd":          "asddasda",
				"x-auth-token": token,
			},
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid token",
			payload: map[string]string{
				"fname":        fnameNew,
				"x-auth-token": "BlaBla",
			},
			expectedStatus: http.StatusUnauthorized,
		}, {
			name: "invalid token not exists",
			payload: map[string]string{
				"fname": fnameNew,
			},
			expectedStatus: http.StatusUnauthorized,
		}, {
			name:           "invalid broken json",
			requestBody:    strings.NewReader(`[{"mail":"` + mailNew + `","x-auth-token":"` + token + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name:           "invalid broken json",
			requestBody:    strings.NewReader(`{"mail":` + mailNew + `","x-auth-token":"` + token + `"}`),
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var req *http.Request
			var url = "http://localhost:3000/user/update/"
			var rec = httptest.NewRecorder()
			if tc.requestBody == nil {
				requestBody := &bytes.Buffer{}
				json.NewEncoder(requestBody).Encode(tc.payload)
				req = httptest.NewRequest("PATCH", url, requestBody)
			} else {
				req = httptest.NewRequest("PATCH", url, tc.requestBody)
			}
			server.HandlerUserUpdate(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else if tc.expectedStatus != http.StatusOK {
				t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
			} else {
				t.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
			}
		})
	}

	t.Run("invalid without request body at all", func(t_ *testing.T) {
		url := "http://localhost:3000/user/update/"
		rec := httptest.NewRecorder()
		req := httptest.NewRequest("PATCH", url, nil)
		server.HandlerUserUpdate(rec, req)
		expectedStatus := http.StatusBadRequest
		if rec.Code != expectedStatus {
			t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, expectedStatus)
		} else {
			t.Logf(GREEN_BG + "SUCCESS: user update was failed as it expected" + NO_COLOR + "\n")
		}
	})
}
