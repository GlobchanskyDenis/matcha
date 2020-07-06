package handlers

import (
	"fmt"
	"testing"
)

// ------------- LOGIN TESTING OF VALID CASES --------------------------

func TestLoginOK_1(t *testing.T) {
	var login = "admin"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_2(t *testing.T) {
	var login = "bsabre-c"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_3(t *testing.T) {
	var login = "0dsaas99sdsdad"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_4(t *testing.T) {
	var login = "Денис Г"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_5(t *testing.T) {
	var login = "new User89"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_6(t *testing.T) {
	var login = "__USER___"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

func TestLoginOK_7(t *testing.T) {
	var login = "Денис skinny"
	err := CheckLogin(login)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", login, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", login)
	}
}

// ------------- LOGIN TESTING OF INVALID CASES --------------------------

func TestLoginInvalid_1(t *testing.T) {
	var login = " Денис"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

func TestLoginInvalid_2(t *testing.T) {
	var login = "Денис "
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

func TestLoginInvalid_3(t *testing.T) {
	var login = "Денис *"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

func TestLoginInvalid_4(t *testing.T) {
	var login = "________"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

func TestLoginInvalid_5(t *testing.T) {
	var login = "A"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

func TestLoginInvalid_6(t *testing.T) {
	var login = "Денис Глобчанский"
	err := CheckLogin(login)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", login)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", login, fmt.Sprintf("%s", err))
	}
}

// ------------- PASSWORD TESTING OF VALID CASES --------------------------

func TestPasswordOK_1(t *testing.T) {
	var passwd = "password654321!!!"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", passwd, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", passwd)
	}
}

func TestPasswordOK_2(t *testing.T) {
	var passwd = "qwerty23@"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", passwd, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", passwd)
	}
}

func TestPasswordOK_3(t *testing.T) {
	var passwd = "m42_new_!pass"
	err := CheckPasswd(passwd)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", passwd, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", passwd)
	}
}

// ------------- PASSWORD TESTING OF INVALID CASES --------------------------

func TestPasswordInvalid_1(t *testing.T) {
	var passwd = "Денис Глобчанский"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_2(t *testing.T) {
	var passwd = "asd2!"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_3(t *testing.T) {
	var passwd = "asdasdad"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_4(t *testing.T) {
	var passwd = "password"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_5(t *testing.T) {
	var passwd = "qwerty"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_6(t *testing.T) {
	var passwd = "123546212"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_7(t *testing.T) {
	var passwd = "asdasdsad34543145"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

func TestPasswordInvalid_8(t *testing.T) {
	var passwd = "admin"
	err := CheckPasswd(passwd)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", passwd)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", passwd, fmt.Sprintf("%s", err))
	}
}

// ------------- MAIL TESTING OF VALID CASES --------------------------

func TestMailOK_1(t *testing.T) {
	var mail = "admin@mail.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

func TestMailOK_2(t *testing.T) {
	var mail = "globchansky.denis@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

func TestMailOK_3(t *testing.T) {
	var mail = "globchansky.denis@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

func TestMailOK_4(t *testing.T) {
	var mail = "globchansky.denis@gmail.msk.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

func TestMailOK_5(t *testing.T) {
	var mail = "g@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

func TestMailOK_6(t *testing.T) {
	var mail = "skinnyman23@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", mail, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", mail)
	}
}

// ------------- MAIL TESTING OF INVALID CASES --------------------------

func TestMailInvalid_1(t *testing.T) {
	var mail = "admin"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_2(t *testing.T) {
	var mail = "денис@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_3(t *testing.T) {
	var mail = "@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_4(t *testing.T) {
	var mail = "a@gm@ail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_5(t *testing.T) {
	var mail = "a@gm.a.il.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_6(t *testing.T) {
	var mail = "myMail@.gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_7(t *testing.T) {
	var mail = "myMailgmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_8(t *testing.T) {
	var mail = "my Mail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_9(t *testing.T) {
	var mail = "myMailmyMailmyMailmyMail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_10(t *testing.T) {
	var mail = "myMail@gmail.com."
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

func TestMailInvalid_11(t *testing.T) {
	var mail = "myMail.@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", mail)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", mail, fmt.Sprintf("%s", err))
	}
}

// ------------- PHONE TESTING OF VALID CASES --------------------------

func TestPhoneOK_1(t *testing.T) {
	var phone = "8(963)648-23-23"
	err := CheckPhone(phone)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", phone, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", phone)
	}
}

func TestPhoneOK_2(t *testing.T) {
	var phone = "+7(123)123-123-12"
	err := CheckPhone(phone)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", phone, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", phone)
	}
}

func TestPhoneOK_3(t *testing.T) {
	var phone = "8 963 648 23 23"
	err := CheckPhone(phone)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", phone, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", phone)
	}
}

func TestPhoneOK_4(t *testing.T) {
	var phone = "+7 963 648 23 23"
	err := CheckPhone(phone)
	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s %s\n\n", phone, fmt.Sprintf("%s", err))
	} else {
		t.Logf("\033[32mDONE\033[m - %s\n\n", phone)
	}
}

// ------------- PHONE TESTING OF INVALID CASES --------------------------

func TestPhoneInvalid_1(t *testing.T) {
	var phone = "admin23"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_2(t *testing.T) {
	var phone = "39185"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_3(t *testing.T) {
	var phone = "1651311321355351321"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_4(t *testing.T) {
	var phone = "8(964)123-213--2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_5(t *testing.T) {
	var phone = "8(964 123-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_6(t *testing.T) {
	var phone = "8((964)123-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_7(t *testing.T) {
	var phone = "8(964)123-213-2 "
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_8(t *testing.T) {
	var phone = " 8(964)123-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_9(t *testing.T) {
	var phone = "8(9+64)123-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_10(t *testing.T) {
	var phone = "8(964)1(23-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}

func TestPhoneInvalid_11(t *testing.T) {
	var phone = "8(964)1)23-213-2"
	err := CheckPhone(phone)
	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - '%s' should be error\n\n", phone)
	} else {
		t.Logf("\033[32mDONE\033[m - '%s' %s -- as it expected\n\n", phone, fmt.Sprintf("%s", err))
	}
}