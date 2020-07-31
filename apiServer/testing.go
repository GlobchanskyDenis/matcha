package apiServer

import (
	"MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	mail = "test@gmail.com"
	pass = "AsdVar34!A"

	mailNew        = "test_new@gmail.com"
	passNew        = "DFe2*FDsd"
	fnameNew       = "Денис"
	lnameNew       = "Глобчанский"
	ageNew         = 21
	genderNew      = "male"
	orientationNew = "hetero"
	bioNew         = `born, suffered, died`
	avaIDNew       = 42

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
	avaIDFail = -1
)

func (server *Server) TestTestUserCreate(t *testing.T, mail string, pass string) config.User {
	t.Helper()

	user, err := server.Db.SetNewUser(mail, handlers.PassHash(pass))
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot create test user - " + err.Error() + config.NO_COLOR + "\n")
		return user
	}
	user.Pass = pass
	user.EncryptedPass = handlers.PassHash(pass)
	user.Status = "confirmed"
	err = server.Db.UpdateUser(user)
	if err != nil {
		t.Errorf(config.RED_BG + "ERROR: Cannot update test user - " + err.Error() + config.NO_COLOR + "\n")
		return user
	}
	return user
}

func (server *Server) TestTestUserAuthorize(t *testing.T, user config.User) string {
	t.Helper()

	url := "http://localhost:3000/user/auth/"
	rec := httptest.NewRecorder()
	requestBody := strings.NewReader(`{"mail":"` + user.Mail + `","pass":"` + user.Pass + `"}`)
	req := httptest.NewRequest("POST", url, requestBody)
	server.HandlerUserAuth(rec, req)
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
