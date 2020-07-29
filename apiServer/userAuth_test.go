package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/database/fakeSql"
	// "MatchaServer/database/postgres"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserAuthenticate(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	/////////// INITIALIZE ///////////

	server, err := New(fakeSql.New())
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot start test server - " + err.Error() + NO_COLOR + "\n")
		return
	}
	testUser := server.TestTestUserCreate(t, mail, passwd)
	defer server.Db.DeleteUser(testUser.Uid)

	/////////// TESTING ///////////

	testCases := []struct {
		name           string
		requestBody		*strings.Reader
		expectedStatus int
	}{
		{
			name: "invalid mail",
			requestBody: strings.NewReader(`{"mail":"` + mailFail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwdFail + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":""}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid empty mail",
			requestBody: strings.NewReader(`{"mail":"","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid no mail",
			requestBody: strings.NewReader(`{"passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid no passwd",
			requestBody: strings.NewReader(`{"mail":"` + mail + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`[{"mail":"` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid broken json",
			requestBody: strings.NewReader(`{"mail":` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusBadRequest,
		}, {
			name: "invalid wrong password",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwdNew + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "invalid not existing mail",
			requestBody: strings.NewReader(`{"mail":"` + mailNew + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusUnprocessableEntity,
		}, {
			name: "valid",
			requestBody: strings.NewReader(`{"mail":"` + mail + `","passwd":"` + passwd + `"}`),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			url := "http://localhost:3000/user/auth/"
			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", url, tc.requestBody)
			server.HttpHandlerUserAuth(rec, req)
			if rec.Code != tc.expectedStatus {
				t_.Errorf(RED_BG+"ERROR: wrong StatusCode: got %d, expected %d"+NO_COLOR+"\n", rec.Code, tc.expectedStatus)
			} else {
				if tc.name != "valid" {
					t.Logf(GREEN_BG + "SUCCESS: user authentication was failed as it expected" + NO_COLOR + "\n")
				} else {
					var response map[string]interface{}
					err := json.NewDecoder(rec.Body).Decode(&response)
					if err != nil {
						t.Errorf(RED_BG+"ERROR: decoding response body error: %s, response body %s"+NO_COLOR+"\n", err.Error(), rec.Body)
						return
					}
					item, isExist := response["x-auth-token"]
					if !isExist {
						t.Errorf(RED_BG + "ERROR: token not found in response" + NO_COLOR + "\n")
						return
					}
					_, ok := item.(string)
					if !ok {
						t.Errorf(RED_BG + "ERROR: token have wrong type" + NO_COLOR + "\n")
						return
					}
					t_.Logf(GREEN_BG + "SUCCESS: user was updated" + NO_COLOR + "\n")
				}
			}
		})
	}
}