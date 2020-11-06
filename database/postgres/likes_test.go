package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"fmt"
	"strconv"
	"testing"
)

func TestLikes(t *testing.T) {
	print(NO_COLOR)
	defer print(YELLOW)

	var (
		conn  ConnDB
		user1 User
		user2 User
		user3 User
	)

	/*
	**	Initialize connection and test users
	 */
	t.Run("Initialize", func(t_ *testing.T) {
		conf, err := config.Create("../../config/")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		err = conn.Connect(&conf.Sql)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		t_.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)

		user1, err = conn.SetNewUser("user1@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user2, err = conn.SetNewUser("user2@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		user3, err = conn.SetNewUser("user3@gmail.com", "qwerty")
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot set new user for tests - " + err.Error() + NO_COLOR)
			t.FailNow()
		}
		t_.Log(GREEN_BG + "SUCCESS: test users was created" + NO_COLOR)
	})

	/*
	**	Test cases - main part of testing
	 */
	testCasesSetLike := []struct {
		name    string
		uid1    int
		uid2    int
		isValid bool
	}{
		{
			name:    "Set like uid=" + strconv.Itoa(user1.Uid) + " to uid=" + strconv.Itoa(user2.Uid),
			uid1:    user1.Uid,
			uid2:    user2.Uid,
			isValid: true,
		}, {
			name:    "Set like uid=" + strconv.Itoa(user1.Uid) + " to uid=" + strconv.Itoa(user3.Uid),
			uid1:    user1.Uid,
			uid2:    user3.Uid,
			isValid: true,
		}, {
			name:    "Set like uid=" + strconv.Itoa(user3.Uid) + " to uid=" + strconv.Itoa(user1.Uid),
			uid1:    user3.Uid,
			uid2:    user1.Uid,
			isValid: true,
		}, {
			name:    "Set like uid=" + strconv.Itoa(user2.Uid) + " to uid=" + strconv.Itoa(user1.Uid),
			uid1:    user2.Uid,
			uid2:    user1.Uid,
			isValid: true,
		}, {
			name:    "Set like uid=" + strconv.Itoa(user2.Uid) + " to uid=" + strconv.Itoa(user3.Uid),
			uid1:    user2.Uid,
			uid2:    user3.Uid,
			isValid: true,
		}, {
			name:    "Set like uid=" + strconv.Itoa(user3.Uid) + " to uid=" + strconv.Itoa(user2.Uid),
			uid1:    user3.Uid,
			uid2:    user2.Uid,
			isValid: true,
		}, {
			name:    "Invalid set like uid=" + strconv.Itoa(user1.Uid) + " to uid=" + strconv.Itoa(user2.Uid),
			uid1:    user1.Uid,
			uid2:    user2.Uid,
			isValid: false,
		},
	}

	/*
	**	Test cases - main part of testing
	 */
	for _, tc := range testCasesSetLike {
		t.Run(tc.name, func(t_ *testing.T) {
			err := conn.SetNewLike(tc.uid1, tc.uid2)
			if tc.isValid {
				if err != nil {
					t_.Errorf(RED_BG+"Error: %s"+NO_COLOR, err.Error())
				} else {
					t_.Logf(GREEN_BG+"Success: like was set from uid=%d to uid=%d"+NO_COLOR, tc.uid1, tc.uid2)
				}
			} else {
				if err == nil {
					t_.Errorf(RED_BG + "Error: not found, but it should be" + NO_COLOR)
				} else {
					t_.Log(GREEN_BG + "Success: like was not set as it expected - " + err.Error() + NO_COLOR)
				}
			}
		})
	}

	testCasesUnsetLike := []struct {
		name    string
		uid1    int
		uid2    int
		isValid bool
	}{
		{
			name:    "Unset like uid=" + strconv.Itoa(user2.Uid) + " to uid=" + strconv.Itoa(user3.Uid),
			uid1:    user2.Uid,
			uid2:    user3.Uid,
			isValid: true,
		}, {
			name:    "Unset like uid=" + strconv.Itoa(user3.Uid) + " to uid=" + strconv.Itoa(user2.Uid),
			uid1:    user3.Uid,
			uid2:    user2.Uid,
			isValid: true,
		}, {
			name:    "Invalid unset like uid=" + strconv.Itoa(user2.Uid) + " to uid=" + strconv.Itoa(user3.Uid),
			uid1:    user2.Uid,
			uid2:    user3.Uid,
			isValid: false,
		},
	}

	for _, tc := range testCasesUnsetLike {
		t.Run(tc.name, func(t_ *testing.T) {
			err := conn.UnsetLike(tc.uid1, tc.uid2)
			if tc.isValid {
				if err != nil {
					t_.Errorf(RED_BG+"Error: %s"+NO_COLOR, err.Error())
				} else {
					t_.Logf(GREEN_BG+"Success: like was unset from uid=%d to uid=%d"+NO_COLOR, tc.uid1, tc.uid2)
				}
			} else {
				if err == nil {
					t_.Errorf(RED_BG + "Error: not found, but it should be" + NO_COLOR)
				} else {
					t_.Log(GREEN_BG + "Success: like was not unset as it expected: " + err.Error() + NO_COLOR)
				}
			}
		})
	}

	testCasesGetUsers := []struct {
		name           string
		uid            int
		expectedAmount int
	}{
		{
			name:           "get users that can speak with user #" + strconv.Itoa(user1.Uid),
			uid:            user1.Uid,
			expectedAmount: 2,
		}, {
			name:           "get users that can speak with user #" + strconv.Itoa(user2.Uid),
			uid:            user2.Uid,
			expectedAmount: 1,
		}, {
			name:           "get users that can speak with user #" + strconv.Itoa(user3.Uid),
			uid:            user3.Uid,
			expectedAmount: 1,
		},
	}

	for _, tc := range testCasesGetUsers {
		t.Run(tc.name, func(t_ *testing.T) {
			users, err := conn.GetFriendUsers(tc.uid)
			if err != nil {
				t_.Errorf(RED_BG+"Error: %s"+NO_COLOR, err.Error())
				t_.FailNow()
			}
			if len(users) != tc.expectedAmount {
				for nbr, usr := range users {
					fmt.Printf("nbr %d uid %d mail %s\n", nbr, usr.Uid, usr.Mail)
				}
				t_.Errorf(RED_BG+"Error: wrong number of detected users. Expected %d found %d"+NO_COLOR, tc.expectedAmount, len(users))
				t_.FailNow()
			}
			t_.Logf(GREEN_BG+"Success: %d users was found"+NO_COLOR, len(users))
		})
	}

	testCasesIsCanSpeak := []struct {
		name           string
		uid1           int
		uid2           int
		expectedResult bool
	}{
		{
			name:           "is can speak user #" + strconv.Itoa(user1.Uid) + " with user #" + strconv.Itoa(user2.Uid),
			uid1:           user1.Uid,
			uid2:           user2.Uid,
			expectedResult: true,
		}, {
			name:           "is can speak user #" + strconv.Itoa(user2.Uid) + " with user #" + strconv.Itoa(user3.Uid),
			uid1:           user2.Uid,
			uid2:           user3.Uid,
			expectedResult: false,
		}, {
			name:           "is can speak user #" + strconv.Itoa(user2.Uid) + " with user #" + strconv.Itoa(user1.Uid),
			uid1:           user2.Uid,
			uid2:           user1.Uid,
			expectedResult: true,
		}, {
			name:           "is can speak user #" + strconv.Itoa(user3.Uid) + " with user #" + strconv.Itoa(user2.Uid),
			uid1:           user3.Uid,
			uid2:           user2.Uid,
			expectedResult: false,
		},
	}

	for _, tc := range testCasesIsCanSpeak {
		t.Run(tc.name, func(t_ *testing.T) {
			result, err := conn.IsICanSpeakWithUser(tc.uid1, tc.uid2)
			if err != nil {
				t_.Errorf(RED_BG+"Error: %s"+NO_COLOR, err.Error())
			} else if tc.expectedResult != result {
				t_.Errorf(RED_BG+"Unexpected result: expected %t got %t"+NO_COLOR, tc.expectedResult, result)
			} else {
				t_.Log(GREEN_BG + "Success" + NO_COLOR)
			}
		})
	}

	/*
	**	Delete test users and close connection
	 */
	t.Run("Delete test users and close connection", func(t_ *testing.T) {
		var wasError bool
		err := conn.DropUserIgnores(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DropUsersLikes(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop likes - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user1.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}

		err = conn.DropUserIgnores(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DropUsersLikes(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop likes - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user2.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}

		err = conn.DropUserIgnores(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop ignores - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DropUsersLikes(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot drop likes - " + err.Error() + NO_COLOR)
			wasError = true
		}
		err = conn.DeleteUser(user3.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: cannot delete - " + err.Error() + NO_COLOR)
			wasError = true
		}
		if !wasError {
			t_.Log(GREEN_BG + "SUCCESS: all test users was deleted" + NO_COLOR)
		}

		conn.Close()
	})
}
