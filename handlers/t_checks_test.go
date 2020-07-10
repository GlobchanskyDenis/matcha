package handlers

import (
	"testing"
	. "MatchaServer/config"
)

// ------------- LOGIN TESTING OF VALID CASES --------------------------

func TestLoginOK_1(t *testing.T) {
	var login = "admin"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_2(t *testing.T) {
	var login = "bsabre-c"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_3(t *testing.T) {
	var login = "0dsaas99sdsdad"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_4(t *testing.T) {
	var login = "Денис Г"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_5(t *testing.T) {
	var login = "new User89"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_6(t *testing.T) {
	var login = "__USER___"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

func TestLoginOK_7(t *testing.T) {
	var login = "Денис skinny"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", login, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", login)
	}
}

// ------------- LOGIN TESTING OF INVALID CASES --------------------------

func TestLoginInvalid_1(t *testing.T) {
	var login = " Денис"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

func TestLoginInvalid_2(t *testing.T) {
	var login = "Денис "
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

func TestLoginInvalid_3(t *testing.T) {
	var login = "Денис *"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

func TestLoginInvalid_4(t *testing.T) {
	var login = "________"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

func TestLoginInvalid_5(t *testing.T) {
	var login = "A"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

func TestLoginInvalid_6(t *testing.T) {
	var login = "Денис Глобчанский"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", login)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", login, err.Error())
	}
}

// ------------- PASSWORD TESTING OF VALID CASES --------------------------

func TestPasswordOK_1(t *testing.T) {
	var passwd = "password654321!!!"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", passwd)
	}
}

func TestPasswordOK_2(t *testing.T) {
	var passwd = "qwerty23@"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", passwd)
	}
}

func TestPasswordOK_3(t *testing.T) {
	var passwd = "m42_new_!pass"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", passwd)
	}
}

// ------------- PASSWORD TESTING OF INVALID CASES --------------------------

func TestPasswordInvalid_1(t *testing.T) {
	var passwd = "Денис Глобчанский"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_2(t *testing.T) {
	var passwd = "asd2!"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_3(t *testing.T) {
	var passwd = "asdasdad"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_4(t *testing.T) {
	var passwd = "password"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_5(t *testing.T) {
	var passwd = "qwerty"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_6(t *testing.T) {
	var passwd = "123546212"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_7(t *testing.T) {
	var passwd = "asdasdsad34543145"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_8(t *testing.T) {
	var passwd = "admin"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", passwd)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", passwd, err.Error())
	}
}

// ------------- MAIL TESTING OF VALID CASES --------------------------

func TestMailOK_1(t *testing.T) {
	var mail = "admin@mail.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

func TestMailOK_2(t *testing.T) {
	var mail = "globchansky.denis@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

func TestMailOK_3(t *testing.T) {
	var mail = "globchansky.denis@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

func TestMailOK_4(t *testing.T) {
	var mail = "globchansky.denis@gmail.msk.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

func TestMailOK_5(t *testing.T) {
	var mail = "g@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

func TestMailOK_6(t *testing.T) {
	var mail = "skinnyman23@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s %s" + NO_COLOR + "\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s" + NO_COLOR + "\n", mail)
	}
}

// ------------- MAIL TESTING OF INVALID CASES --------------------------

func TestMailInvalid_1(t *testing.T) {
	var mail = "admin"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_2(t *testing.T) {
	var mail = "денис@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_3(t *testing.T) {
	var mail = "@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_4(t *testing.T) {
	var mail = "a@gm@ail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_5(t *testing.T) {
	var mail = "a@gm.a.il.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_6(t *testing.T) {
	var mail = "myMail@.gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_7(t *testing.T) {
	var mail = "myMailgmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_8(t *testing.T) {
	var mail = "my Mail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_9(t *testing.T) {
	var mail = "myMailmyMailmyMailmyMail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_10(t *testing.T) {
	var mail = "myMail@gmail.com."
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}

func TestMailInvalid_11(t *testing.T) {
	var mail = "myMail.@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG + "FAILED: '%s' should be error" + NO_COLOR + "\n", mail)
	} else {
		t.Logf(GREEN_BG + "SUCCESS: '%s' %s -- as it expected" + NO_COLOR + "\n", mail, err.Error())
	}
}
