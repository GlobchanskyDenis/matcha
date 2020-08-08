package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"encoding/json"
	"errors"
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
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckMail(user.Mail)
		if err != nil {
			return user, message, err
		}
		message += " mail=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["pass"]
	if isExist {
		usefullFieldsExists = true
		user.Pass, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		user.EncryptedPass = handlers.PassHash(user.Pass)
		err = handlers.CheckPass(user.Pass)
		if err != nil {
			return user, message, err
		}
		message += " password=hidden"
	}
	arg, isExist = request["fname"]
	if isExist {
		usefullFieldsExists = true
		user.Fname, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckName(user.Fname)
		if err != nil {
			return user, message, err
		}
		message += " fname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["lname"]
	if isExist {
		usefullFieldsExists = true
		user.Lname, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckName(user.Lname)
		if err != nil {
			return user, message, err
		}
		message += " lname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["birth"]
	if isExist {
		usefullFieldsExists = true
		birth, ok := arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		user.Birth, err = time.Parse("2006-01-02", birth)
		if err != nil {
			return user, message, err
		}
		user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
		if user.Age > 80 || user.Age < 16 {
			return user, message, errors.New("forbidden age")
		}
		message += " birth=" + BLUE + birth + NO_COLOR + " age=" + BLUE + strconv.Itoa(user.Age) + NO_COLOR
	}
	arg, isExist = request["gender"]
	if isExist {
		usefullFieldsExists = true
		user.Gender, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckGender(user.Gender)
		if err != nil {
			return user, message, err
		}
		message += " gender=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["orientation"]
	if isExist {
		usefullFieldsExists = true
		user.Orientation, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckOrientation(user.Orientation)
		if err != nil {
			return user, message, err
		}
		message += " orientation=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["bio"]
	if isExist {
		usefullFieldsExists = true
		user.Bio, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckBio(user.Bio)
		if err != nil {
			return user, message, err
		}
		message += " bio=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["avaID"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		user.AvaID = int(tmpFloat)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		if user.AvaID < 0 {
			return user, message, errors.New("this id is forbidden")
		}
		message += " avaID=" + BLUE + strconv.Itoa(user.AvaID) + NO_COLOR
	}
	arg, isExist = request["latitude"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		user.Latitude = float32(tmpFloat)
		message += " latitude=" + BLUE + strconv.FormatFloat(tmpFloat, 'E', -1, 32) + NO_COLOR
	}
	arg, isExist = request["longitude"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		user.Latitude = float32(tmpFloat)
		message += " longitude=" + BLUE + strconv.FormatFloat(tmpFloat, 'E', -1, 32) + NO_COLOR
	}
	arg, isExist = request["interests"]
	if isExist {
		usefullFieldsExists = true
		interfaceArr, ok = arg.([]interface{})
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		for _, item := range interfaceArr {
			tmpStr, ok := item.(string)
			if !ok {
				return user, message, errors.New("wrong type of param")
			}
			err = handlers.CheckInterest(tmpStr)
			if err != nil {
				return user, message, errors.New("invalid interest - " + err.Error())
			}
			user.Interests = append(user.Interests, tmpStr)
			interestsStr += tmpStr + ", "
		}
		if len(interestsStr) > 2 {
			interestsStr = string(interestsStr[:len(interestsStr) - 2])
		}
		message += " interests=" + BLUE + interestsStr + NO_COLOR
	}

	if !usefullFieldsExists {
		return user, message, errors.New("no usefull fields found")
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
		consoleLogError(r, "/user/update/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "json decode failed" + `"}`))
		return
	}

	arg, isExist := request["x-auth-token"]
	if !isExist {
		consoleLogWarning(r, "/user/update/", "token not exists")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "token not exists" + `"}`))
		return
	}

	token, ok := arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/update/", "token have wrong type")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "token have wrong type" + `"}`))
		return
	}

	if token == "" {
		consoleLogWarning(r, "/user/update/", "token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "token is empty" + `"}`))
		return
	}

	arg, isExist = request["interests"]
	if isExist {
		var interestsNameArr []string
		knownInterests, err := server.Db.GetInterests()
		if err != nil {
			consoleLogWarning(r, "/user/update/", "GetInterests returned error - " + err.Error())
			w.WriteHeader(http.StatusInternalServerError) // 500
			w.Write([]byte(`{"error":"` + "database error" + `"}`))
			return
		}
		interfaceArr, ok := arg.([]interface{})
		if !ok {
			consoleLogWarning(r, "/user/update/", "wrong argument type (interests)")
			w.WriteHeader(http.StatusUnprocessableEntity) // 422
			w.Write([]byte(`{"error":"` + "wrong argument type (interests)" + `"}`))
			return
		}
		for _, item := range interfaceArr {
			interest, ok := item.(string)
			if !ok {
				consoleLogWarning(r, "/user/update/", "wrong argument type (interests item)")
				w.WriteHeader(http.StatusUnprocessableEntity) // 422
				w.Write([]byte(`{"error":"` + "wrong argument type (interests item)" + `"}`))
				return
			}
			err = handlers.CheckInterest(interest)
			if err != nil {
				consoleLogWarning(r, "/user/update/", "invalid interest - " + err.Error())
				w.WriteHeader(http.StatusUnprocessableEntity) // 422
				w.Write([]byte(`{"error":"` + "invalid interest - " + err.Error() + `"}`))
				return
			}
			interestsNameArr = append(interestsNameArr, interest)
		}
		unknownInterests := handlers.FindUnknownInterests(knownInterests, interestsNameArr)
		err = server.Db.AddInterests(unknownInterests)
		if err != nil {
			consoleLogError(r, "/user/update/", "AddInterests returned error - " + err.Error())
			w.WriteHeader(http.StatusInternalServerError) // 500
			w.Write([]byte(`{"error":"` + "database error" + `"}`))
			return
		}
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/update/", "TokenUidDecode returned error - "+err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "token decoding error" + `"}`))
		return
	}

	if !server.session.IsUserLoggedByUid(uid) {
		consoleLogWarning(r, "/user/update/", "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" is not logged")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "user is not logged" + `"}`))
		return
	}

	user, err = server.Db.GetUserByUid(uid)
	if errDef.IsRecordNotFoundError(err) {
		consoleLogWarning(r, "/user/update/", "GetUserByUid - record not found")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	} else if err != nil {
		consoleLogError(r, "/user/update/", "GetUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request returned error" + `"}`))
		return
	}

	user, message, err = fillUserStruct(request, user)
	if err != nil {
		consoleLogWarning(r, "/user/update/", err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	consoleLog(r, "/user/update/", message)



	err = server.Db.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/update/", "UpdateUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request returned error" + `"}`))
		return
	}

	// Проверить - принадлежит ли фото юзеру

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/update/", "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
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

		consoleLog(r, "/user/update/", "client wants to know what methods are allowed")

	} else if r.Method == "PATCH" {

		server.userUpdate(w, r)

	} else {
		// ALL OTHER METHODS

		consoleLogWarning(r, "/user/update/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
