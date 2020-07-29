package apiServer

import (
	"testing"
	"MatchaServer/config"
	"MatchaServer/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
)

var (
	mail   = "test@gmail.com"
	passwd = "AsdVar34!A"

	mailNew        = "test_new@gmail.com"
	passwdNew      = "DFe2*FDsd"
	fnameNew       = "Денис"
	lnameNew       = "Глобчанский"
	ageNew         = 21
	genderNew      = "male"
	orientationNew = "hetero"
	biographyNew   = `born, suffered, died`
	avaPhotoIDNew  = 42

	mailFail        = "mail@gmail@yandex.ru"
	passwdFail      = "12345678"
	fnameFail       = "@Денис"
	lnameFail       = "qweкий   "
	ageFail         = 217
	genderFail      = "thing"
	orientationFail = "люблю всех"
	biographyFail   = `фвыфв ывфывфщзшзщольджук  йлофыдлвоы фыдлвоыдвлффды дл 
	ывофыдлвоыфлдвоы оыфво фылдво л ыовлывфвфыовфыд офыл офвд лфывыфлво фв флдв офлвдофы лфо фдылов
	sdsadasdsa sadasdasdasd asd asdsadas as asdasdsad as`
	avaPhotoIDFail = -1
)

func (server *Server) TestTestUserCreate(t *testing.T, mail string, passwd string) config.User {
	t.Helper()

	user, err := server.Db.SetNewUser(mail, handlers.PasswdHash(passwd))
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot create test user - " + err.Error() + config.NO_COLOR + "\n")
		return user
	}
	user.Passwd = handlers.PasswdHash(passwd)
	user.AccType = "confirmed"
	err = server.Db.UpdateUser(user)
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot update test user - " + err.Error() + config.NO_COLOR + "\n")
		return user
	}
	user.Passwd = passwd
	return user
}

func (server *Server) TestTestUserAuthorize(t *testing.T, user config.User) string {
	t.Helper()

	url := "http://localhost:3000/user/auth/"
	rec := httptest.NewRecorder()
	println(user.Mail)
	println(user.Passwd)
	requestBody := strings.NewReader(`{"mail":"` + user.Mail + `","passwd":"` + user.Passwd + `"}`)
	req := httptest.NewRequest("POST", url, requestBody)
	server.HttpHandlerUserAuth(rec, req)
	if rec.Code != http.StatusOK {
		println("Expected 200")
		print("Got ")
		println(rec.Code)
		t.Errorf(config.RED_BG + "ERROR: wrong response status code while user authentication" + config.NO_COLOR + "\n")
		t.Fatal()
	}
	var response map[string]interface{}
	err := json.NewDecoder(rec.Body).Decode(&response)
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: json decode error while user authentication - " + err.Error() + config.NO_COLOR + "\n")
		t.Fatal()
	}
	item, isExist := response["x-auth-token"]
	if !isExist {
		t.Errorf(config.RED_BG + "ERROR: x-auth-token not exists in response of user authentication" + config.NO_COLOR + "\n")
		t.Fatal()
	}
	_, ok := item.(string)
	if !ok {
		t.Errorf(config.RED_BG + "ERROR: x-auth-token have wrong type" + config.NO_COLOR + "\n")
		t.Fatal()
	}
	return item.(string)
}