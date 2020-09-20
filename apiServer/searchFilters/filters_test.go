package searchFilters

import (
	"MatchaServer/apiServer"
	"encoding/json"
	"strings"
	"testing"
)

// func Test1(t *testing.T) {
// 	Fs := Filters{}
// 	F1 := ageFilter{minAge: 30, maxAge: 30}
// 	F2 := onlineFilter{uidSlice: append([]int{}, 1, 2, 3)}
// 	// F2 := onlineFilter{uidSlice: append([]int{}, 1)}
// 	// F2 := onlineFilter{uidSlice: []int{}}
// 	F3 := ratingFilter{minRating: 0, maxRating: 5}
// 	F4 := interestsFilter{interests: append([]string{}, "fun", "football")}
// 	F5 := locationFilter{minLongitude: 32.5, maxLongitude: 46.4, minLatitude: 3.1415, maxLatitude: 15.45}

// 	Fs.filters = append(Fs.filters, F1, F2, F3, F4, F5)
// 	println(Fs.PrepareQuery(""))
// 	fmt.Println("online filter type ", F2.getFilterType())
// 	fmt.Println("age filter type ", F1.getFilterType())
// 	fmt.Println("rating filter type ", F3.getFilterType())
// 	fmt.Println("interest filter type ", F4.getFilterType())
// 	fmt.Println("location filter type ", F5.getFilterType())
// 	println(Fs.Print())
// }

func TestAge(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid - full params",
			payload:   `{"age":{"min":18,"max":35}}`,
			isInvalid: false,
		}, {
			name:      "valid - first param exist",
			payload:   `{"age":{"min":18}}`,
			isInvalid: false,
		}, {
			name:      "valid - second param exist",
			payload:   `{"age":{"max":35}}`,
			isInvalid: false,
		}, {
			name:      "invalid - no params",
			payload:   `{"age":{}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #1",
			payload:   `{"age":{"min":35,"max":18}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #2",
			payload:   `{"age":{"min":-18,"max":35}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #3",
			payload:   `{"age":{"min":18,"max":-35}}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			err = f.Parse(params, uid, server.Db, &server.Session)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}

func TestRating(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid - full params",
			payload:   `{"rating":{"min":18,"max":35}}`,
			isInvalid: false,
		}, {
			name:      "valid - first param exist",
			payload:   `{"rating":{"min":18}}`,
			isInvalid: false,
		}, {
			name:      "valid - second param exist",
			payload:   `{"rating":{"max":35}}`,
			isInvalid: false,
		}, {
			name:      "invalid - no params",
			payload:   `{"rating":{}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #1",
			payload:   `{"rating":{"min":35,"max":18}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #2",
			payload:   `{"rating":{"min":-18,"max":35}}`,
			isInvalid: true,
		}, {
			name:      "invalid - wrong params #3",
			payload:   `{"rating":{"min":18,"max":-35}}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			err = f.Parse(params, uid, server.Db, &server.Session)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}

func TestInterests(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid - full params",
			payload:   `{"interests":["football","starcraft"]}`,
			isInvalid: false,
		}, {
			name:      "valid - one param",
			payload:   `{"interests":["football"]}`,
			isInvalid: false,
		}, {
			name:      "invalid - no params",
			payload:   `{"interests":[]}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			err = f.Parse(params, uid, server.Db, &server.Session)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}

func TestLocation(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid - full params",
			payload:   `{"location":{"minLatitude":23,"maxLatitude":24,"minLongitude":-54.43,"maxLongitude":3.1415}}`,
			isInvalid: false,
		}, {
			name:      "valid - only latitude",
			payload:   `{"location":{"minLatitude":23,"maxLatitude":24}}`,
			isInvalid: false,
		}, {
			name:      "valid - only longitude",
			payload:   `{"location":{"minLongitude":-54.43,"maxLongitude":3.1415}}`,
			isInvalid: false,
		}, {
			name:      "valid - only min",
			payload:   `{"location":{"minLongitude":-54.43,"minLatitude":23}}`,
			isInvalid: false,
		}, {
			name:      "valid - only max",
			payload:   `{"location":{"maxLongitude":3.1415,"maxLatitude":24}}`,
			isInvalid: false,
		}, {
			name:      "valid - only one param",
			payload:   `{"location":{"maxLongitude":3.1415}}`,
			isInvalid: false,
		}, {
			name:      "invalid - no params",
			payload:   `{"location":{}}`,
			isInvalid: true,
		}, {
			name:      "invalid - latitude #1",
			payload:   `{"location":{"minLatitude":-230,"maxLatitude":24}}`,
			isInvalid: true,
		}, {
			name:      "invalid - latitude #2",
			payload:   `{"location":{"minLatitude":23,"maxLatitude":240}}`,
			isInvalid: true,
		}, {
			name:      "invalid - latitude #3",
			payload:   `{"location":{"minLatitude":24,"maxLatitude":23}}`,
			isInvalid: true,
		}, {
			name:      "invalid - longitude #1",
			payload:   `{"location":{"minLongitude":-230,"maxLongitude":24}}`,
			isInvalid: true,
		}, {
			name:      "invalid - longitude #2",
			payload:   `{"location":{"minLongitude":23,"maxLongitude":240}}`,
			isInvalid: true,
		}, {
			name:      "invalid - longitude #3",
			payload:   `{"location":{"minLongitude":24,"maxLongitude":23}}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			err = f.Parse(params, uid, server.Db, &server.Session)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}

func TestRadius(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid - full params #1",
			payload:   `{"radius":{"latitude":23,"longitude":-54.43,"radius":111}}`,
			isInvalid: false,
		}, {
			name:      "valid - full params #2",
			payload:   `{"radius":{"latitude":23,"longitude":-54.43,"radius":222}}`,
			isInvalid: false,
		}, {
			name:      "valid - no location",
			payload:   `{"radius":{"radius":222}}`,
			isInvalid: false,
		}, {
			name:      "invalid radius",
			payload:   `{"radius":{"radius":-222}}`,
			isInvalid: true,
		}, {
			name:      "invalid - only latitude in coordinates",
			payload:   `{"radius":{"latitude":23,"radius":222}}`,
			isInvalid: true,
		}, {
			name:      "invalid - only longitude in coordinates",
			payload:   `{"radius":{"longitude":23,"radius":222}}`,
			isInvalid: true,
		}, {
			name:      "invalid - no radius",
			payload:   `{"radius":{"latitude":-2,"longitude":23}}`,
			isInvalid: true,
		}, {
			name:      "invalid - invalid latitude",
			payload:   `{"radius":{"latitude":-242,"longitude":23,"radius":222}}`,
			isInvalid: true,
		}, {
			name:      "invalid - no params at all",
			payload:   `{"radius":{}}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			err = f.Parse(params, uid, server.Db, &server.Session)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}

func TestOnline(t *testing.T) {
	uid := 1
	server, err := apiServer.New("../../config/")
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}
	_, err = server.Session.AddUserToSession(23)
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	_, err = server.Session.AddUserToSession(42)
	if err != nil {
		t.Errorf("Cannot start test: "+err.Error())
		return
	}

	testCases := []struct {
		name      string
		payload   string
		isInvalid bool
	}{
		{
			name:      "valid",
			payload:   `{"online":{}}`,
			isInvalid: false,
		}, {
			name:      "invalid",
			payload:   `{"online":{}}`,
			isInvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t_ *testing.T) {
			var params map[string]interface{}
			var f = &Filters{}
			reader := strings.NewReader(tc.payload)
			err := json.NewDecoder(reader).Decode(&params)
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if !tc.isInvalid {
				err = f.Parse(params, uid, server.Db, &server.Session)
			} else {
				err = f.Parse(params, uid, server.Db, nil)
			}
			if err != nil {
				if tc.isInvalid {
					t.Log("Success: error found as it expected - " + err.Error())
				} else {
					t.Errorf("Error: " + err.Error())
				}
				return
			}

			if tc.isInvalid {
				t.Errorf("Error: it should be error, but it expected " + tc.name)
				return
			}

			println(f.Print())
			println(f.PrepareQuery(""))
		})
	}
}
