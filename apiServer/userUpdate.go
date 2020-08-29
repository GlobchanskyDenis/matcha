package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func fillUserStruct(request map[string]interface{}, user User) (User, string, error) {
	var usefullFieldsExists, ok, isExist bool
	var message string
	var err error
	var tmpFloat float64
	var interfaceArr []interface{}
	var interestsStr string

	message = "request for UPDATE was recieved: "

	arg, isExist := request["mail"]
	if isExist {
		usefullFieldsExists = true
		user.Mail, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type")
		}
		err = handlers.CheckMail(user.Mail)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " mail=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["pass"]
	if isExist {
		usefullFieldsExists = true
		user.Pass, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type")
		}
		user.EncryptedPass = handlers.PassHash(user.Pass)
		err = handlers.CheckPass(user.Pass)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " password=hidden"
	}
	arg, isExist = request["fname"]
	if isExist {
		usefullFieldsExists = true
		user.Fname, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле fname имеет неверный тип", "fname field has wrong type")
		}
		err = handlers.CheckName(user.Fname)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " fname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["lname"]
	if isExist {
		usefullFieldsExists = true
		user.Lname, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле lname имеет неверный тип", "lname field has wrong type")
		}
		err = handlers.CheckName(user.Lname)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " lname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["birth"]
	if isExist {
		usefullFieldsExists = true
		birth, ok := arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле birth имеет неверный тип", "birth field has wrong type")
		}
		user.Birth, err = time.Parse("2006-01-02", birth)
		if err != nil {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле birth имеет неверный формат", "birth field has wrong format")
		}
		user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
		if user.Age > 80 || user.Age < 16 {
			return user, message, errDef.InvalidArgument.WithArguments("Значение поля birth недопустимо", "birth field has wrong value")
		}
		message += " birth=" + BLUE + birth + NO_COLOR + " age=" + BLUE + strconv.Itoa(user.Age) + NO_COLOR
	}
	arg, isExist = request["gender"]
	if isExist {
		usefullFieldsExists = true
		user.Gender, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле gender имеет неверный тип", "gender field has wrong type")
		}
		err = handlers.CheckGender(user.Gender)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " gender=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["orientation"]
	if isExist {
		usefullFieldsExists = true
		user.Orientation, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле orientation имеет неверный тип", "orientation field has wrong type")
		}
		err = handlers.CheckOrientation(user.Orientation)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " orientation=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["bio"]
	if isExist {
		usefullFieldsExists = true
		user.Bio, ok = arg.(string)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле bio имеет неверный тип", "bio field has wrong type")
		}
		err = handlers.CheckBio(user.Bio)
		if err != nil {
			// handlers - привести все ошибки к типу errDef.ApiErrorArgument
			return user, message, errDef.InvalidArgument.WithArguments(err)
		}
		message += " bio=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["avaID"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		user.AvaID = int(tmpFloat)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле avaID имеет неверный тип", "avaID field has wrong type")
		}
		if user.AvaID < 0 {
			return user, message, errDef.InvalidArgument.WithArguments("Значение поля avaID недопустимо", "avaID field has wrong value")
		}
		message += " avaID=" + BLUE + strconv.Itoa(user.AvaID) + NO_COLOR
	}
	arg, isExist = request["latitude"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле latitude имеет неверный тип", "latitude field has wrong type")
		}
		user.Latitude = float32(tmpFloat)
		message += " latitude=" + BLUE + strconv.FormatFloat(tmpFloat, 'E', -1, 32) + NO_COLOR
	}
	arg, isExist = request["longitude"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле longitude имеет неверный тип", "longitude field has wrong type")
		}
		user.Longitude = float32(tmpFloat)
		message += " longitude=" + BLUE + strconv.FormatFloat(tmpFloat, 'E', -1, 32) + NO_COLOR
	}
	arg, isExist = request["interests"]
	if isExist {
		usefullFieldsExists = true
		interfaceArr, ok = arg.([]interface{})
		if !ok {
			return user, message,
				errDef.InvalidArgument.WithArguments("Поле interests имеет неверный тип", "interests field has wrong type")
		}
		// вытираю старые интересы - чтобы не было дублирования
		user.Interests = nil
		for _, item := range interfaceArr {
			tmpStr, ok := item.(string)
			if !ok {
				return user, message, errDef.InvalidArgument.WithArguments("Поле interests имеет неверный тип", "interests field has wrong type")
			}
			err = handlers.CheckInterest(tmpStr)
			if err != nil {
				// handlers - привести все ошибки к типу errDef.ApiErrorArgument
				return user, message, errDef.InvalidArgument.WithArguments(err)
			}
			user.Interests = append(user.Interests, tmpStr)
			interestsStr += tmpStr + ", "
		}
		if len(interestsStr) > 2 {
			interestsStr = string(interestsStr[:len(interestsStr)-2])
		}
		message += " interests=" + BLUE + interestsStr + NO_COLOR
	}

	if !usefullFieldsExists {
		return user, message, errDef.NoArgument //.WithArguments("Нет ни одного полезного поля", "no usefull fields found")
	}
	return user, message, nil
}

// USER UPDATE BY PATCH METHOD
// REQUEST BODY IS JSON
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
// RESPONSE BODY IS JSON ONLY IN CASE OF ERROR. IN OTHER CASE - NO RESPONSE BODY
func (server *Server) userUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		uid            int
		err            error
		user           User
		message, token string
		request        map[string]interface{}
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		server.LogError(r, "/user/update/", "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["x-auth-token"]
	if !isExist {
		server.LogWarning(r, "/user/update/", "x-auth-token not exists")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
		return
	}

	token, ok := arg.(string)
	if !ok {
		server.LogWarning(r, "/user/update/", "token have wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		return
	}

	if token == "" {
		server.LogWarning(r, "/user/update/", "token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	arg, isExist = request["interests"]
	if isExist {
		var interestsNameArr []string
		knownInterests, err := server.Db.GetInterests()
		if err != nil {
			server.LogWarning(r, "/user/update/", "GetInterests returned error - "+err.Error())
			server.error(w, errDef.DatabaseError)
			return
		}
		interfaceArr, ok := arg.([]interface{})
		if !ok {
			server.LogWarning(r, "/user/update/", "wrong argument type (interests)")
			server.error(w, errDef.InvalidArgument.WithArguments("Поле interests имеет неверный тип", "interests field has wrong type"))
			return
		}
		for _, item := range interfaceArr {
			interest, ok := item.(string)
			if !ok {
				server.LogWarning(r, "/user/update/", "wrong argument type (interests item)")
				server.error(w, errDef.InvalidArgument.WithArguments("Поле interests (item) имеет неверный тип", "interests (item) field has wrong type"))
				return
			}
			err = handlers.CheckInterest(interest)
			if err != nil {
				server.LogWarning(r, "/user/update/", "invalid interest - "+err.Error())
				server.error(w, errDef.InvalidArgument.WithArguments("Значение поля interests (item) недопустимо",
					"interests (item) field has wrong value"))
				return
			}
			interestsNameArr = append(interestsNameArr, interest)
		}
		unknownInterests := handlers.FindUnknownInterests(knownInterests, interestsNameArr)
		err = server.Db.AddInterests(unknownInterests)
		if err != nil {
			server.LogError(r, "/user/update/", "AddInterests returned error - "+err.Error())
			server.error(w, errDef.DatabaseError)
			return
		}
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		server.LogWarning(r, "/user/update/", "TokenUidDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	if !server.session.IsUserLoggedByUid(uid) {
		server.LogWarning(r, "/user/update/", "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" is not logged")
		server.error(w, errDef.UserNotLogged)
		return
	}

	user, err = server.Db.GetUserByUid(uid)
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		server.LogWarning(r, "/user/update/", "GetUserByUid - record not found")
		server.error(w, errDef.UserNotExist)
		return
	} else if err != nil {
		server.LogError(r, "/user/update/", "GetUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	user, message, err = fillUserStruct(request, user)
	if err != nil {
		server.LogWarning(r, "/user/update/", err.Error())
		server.error(w, err.(errDef.ApiError))
		return
	}

	server.Log(r, "/user/update/", message)

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.LogError(r, "/user/update/", "UpdateUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	// Проверить - принадлежит ли фото юзеру

	w.WriteHeader(http.StatusOK) // 200
	server.LogSuccess(r, "/user/update/", "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/
// UPDATE USER BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		server.Log(r, "/user/update/", "client wants to know what methods are allowed")

	} else if r.Method == "PATCH" {

		server.userUpdate(w, r)

	} else {
		// ALL OTHER METHODS

		server.LogWarning(r, "/user/update/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
