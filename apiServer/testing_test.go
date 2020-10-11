package apiServer

import (
	"MatchaServer/common"
	"MatchaServer/handlers"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	// "testing"
)

var (
	mail = "test@gmail.com"
	pass = "AsdVar34!A"

	mailNew                = "test_new@gmail.com"
	passNew                = "DFe2*FDsd"
	fnameNew               = "Денис"
	lnameNew               = "Глобчанский"
	ageNew                 = 21
	genderNew              = "male"
	orientationNew         = "hetero"
	bioNew                 = `born, suffered, died`
	avaIDNew       float64 = 42
	latitudeNew    float64 = 3.1415
	longitudeNew   float64 = 56.1
	interests1New          = append([]interface{}{}, "fun", "other", "football")
	interests2New          = []interface{}{}

	mailFail        = "mail@gmail@yandex.ru"
	passFail        = "12345678"
	fnameFail       = "@Денис"
	lnameFail       = "qweкий   "
	ageFail         = 217
	genderFail      = "thing"
	orientationFail = "люблю всех"
	bioFail         = `фвыфв ывфывфщзшзщольджук  йлофыдлвоы фыдлвоыдвлффды дл 
	ывофыдлвоыфлдвоы оыфво фылдво л ыовлывфвфыовфыд офыл офвд лфывыфлво фв флдв офлвдофы лфо фдылов
	sdsadasdsa sadasdasdasd asd asdsadas as asdasdsad as`
	avaIDFail               = -1
	latitudeFail            = "123"
	longitudeFail           = "asd"
	interests1Fail          = [...]string{"asdasdsadasdasdasdadasdasdadfsdfdsfsdfgfdgfgdfsdfsfsdsadasdasdsadsdasdasdasdasdasdasdasdasdasdasdadsdasddasdasdsadasdasd"}
	interests2Fail          = []string{"", "", "football"}
	interests3Fail []string = nil
)

/*
**	Function that used only in unit testing
 */
func (server *Server) CreateTestUser(mail string, pass string) (common.User, error) {
	user, err := server.Db.SetNewUser(mail, handlers.PassHash(pass))
	if err != nil {
		return user, err
	}
	user.Pass = pass
	user.EncryptedPass = handlers.PassHash(pass)
	user.Status = "confirmed"
	err = server.Db.UpdateUser(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

/*
**	Function that used only in unit testing
 */
func (server *Server) AuthorizeTestUser(user common.User) error {
	var (
		requestParams map[string]interface{}
		err           error
		ctx           context.Context
		url           = "http://localhost:3000/user/auth/"
		rec           = httptest.NewRecorder()
		requestBody   = strings.NewReader(`{"mail":"` + user.Mail + `","pass":"` + user.Pass + `"}`)
		req           = httptest.NewRequest("POST", url, requestBody)
	)

	err = json.NewDecoder(req.Body).Decode(&requestParams)
	if err != nil {
		return errors.New("Cannot start test because of error: " + err.Error())
	}
	ctx = context.WithValue(req.Context(), "requestParams", requestParams)
	server.UserAuth(rec, req.WithContext(ctx))
	if rec.Code != http.StatusOK {
		return errors.New("User auth - wrong response status. Expected 200 got " + strconv.Itoa(rec.Code))
	}
	var response map[string]interface{}
	err = json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		return errors.New("json decode error - " + err.Error())
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		return errors.New("x-auth-token not exist in response of test user auth")
	}
	_, ok := item.(string)
	if !ok {
		return errors.New("x-auth-token have wrong type")
	}
	return nil
}
