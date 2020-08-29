package handlers

import (
	. "MatchaServer/common"
	"testing"
)

// ------------- PASSWORD TESTING OF VALID CASES --------------------------

func TestPasswordOK_1(t *testing.T) {
	var passwd = "password654321!!!"
	err := CheckPass(passwd)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", passwd)
	}
}

func TestPasswordOK_2(t *testing.T) {
	var passwd = "qwerty23@"
	err := CheckPass(passwd)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", passwd)
	}
}

func TestPasswordOK_3(t *testing.T) {
	var passwd = "m42_new_!pass"
	err := CheckPass(passwd)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", passwd, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", passwd)
	}
}

// ------------- PASSWORD TESTING OF INVALID CASES --------------------------

func TestPasswordInvalid_1(t *testing.T) {
	var passwd = "Денис Глобчанский"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_2(t *testing.T) {
	var passwd = "asd2!"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_3(t *testing.T) {
	var passwd = "asdasdad"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_4(t *testing.T) {
	var passwd = "password"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_5(t *testing.T) {
	var passwd = "qwerty"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_6(t *testing.T) {
	var passwd = "123546212"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_7(t *testing.T) {
	var passwd = "asdasdsad34543145"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

func TestPasswordInvalid_8(t *testing.T) {
	var passwd = "admin"
	err := CheckPass(passwd)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", passwd)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", passwd, err.Error())
	}
}

// ------------- MAIL TESTING OF VALID CASES --------------------------

func TestMailOK_1(t *testing.T) {
	var mail = "admin@mail.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

func TestMailOK_2(t *testing.T) {
	var mail = "globchansky.denis@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

func TestMailOK_3(t *testing.T) {
	var mail = "globchansky.denis@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

func TestMailOK_4(t *testing.T) {
	var mail = "globchansky.denis@gmail.msk.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

func TestMailOK_5(t *testing.T) {
	var mail = "g@gmail.com"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

func TestMailOK_6(t *testing.T) {
	var mail = "skinnyman23@yandex.ru"
	err := CheckMail(mail)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", mail, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", mail)
	}
}

// ------------- MAIL TESTING OF INVALID CASES --------------------------

func TestMailInvalid_1(t *testing.T) {
	var mail = "admin"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_2(t *testing.T) {
	var mail = "денис@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_3(t *testing.T) {
	var mail = "@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_4(t *testing.T) {
	var mail = "a@gm@ail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_5(t *testing.T) {
	var mail = "a@gm.a.il.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_6(t *testing.T) {
	var mail = "myMail@.gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_7(t *testing.T) {
	var mail = "myMailgmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_8(t *testing.T) {
	var mail = "my Mail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_9(t *testing.T) {
	var mail = "myMailmyMailmyMailmyMail@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_10(t *testing.T) {
	var mail = "myMail@gmail.com."
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

func TestMailInvalid_11(t *testing.T) {
	var mail = "myMail.@gmail.com"
	err := CheckMail(mail)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", mail)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", mail, err.Error())
	}
}

// ------------- NAME TESTING OF VALID CASES --------------------------

func TestNameOK_1(t *testing.T) {
	var Name = "admin"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_2(t *testing.T) {
	var Name = "bsabre-c"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_3(t *testing.T) {
	var Name = "dsaas99sdsdad0"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_4(t *testing.T) {
	var Name = "Денис Г"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_5(t *testing.T) {
	var Name = "new User89"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_6(t *testing.T) {
	var Name = "Денис ибн Сергей"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_7(t *testing.T) {
	var Name = "Денис skinny"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

func TestNameOK_8(t *testing.T) {
	var Name = "Ю"
	err := CheckName(Name)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", Name, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", Name)
	}
}

// ------------- NAME TESTING OF INVALID CASES --------------------------

func TestNameInvalid_1(t *testing.T) {
	var Name = " Денис"
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

func TestNameInvalid_2(t *testing.T) {
	var Name = "Денис "
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

func TestNameInvalid_3(t *testing.T) {
	var Name = "Денис *"
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

func TestNameInvalid_4(t *testing.T) {
	var Name = "________"
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

func TestNameInvalid_5(t *testing.T) {
	var Name = "Денис Глобчанский"
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

func TestNameInvalid_6(t *testing.T) {
	var Name = "Денис\tГ"
	err := CheckName(Name)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", Name)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", Name, err.Error())
	}
}

// ------------- GENDER TESTING OF VALID CASES --------------------------

func TestGenderOK_1(t *testing.T) {
	var gender = "male"
	err := CheckGender(gender)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", gender, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", gender)
	}
}

func TestGenderOK_2(t *testing.T) {
	var gender = "female"
	err := CheckGender(gender)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", gender, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", gender)
	}
}

// ------------- GENDER TESTING OF INVALID CASES --------------------------

func TestGenderInvalid_1(t *testing.T) {
	var gender = "Полубог"
	err := CheckGender(gender)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", gender)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", gender, err.Error())
	}
}

func TestGenderInvalid_2(t *testing.T) {
	var gender = ""
	err := CheckGender(gender)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", gender)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", gender, err.Error())
	}
}

// ------------- ORIENTATION TESTING OF VALID CASES --------------------------

func TestOrientationOK_1(t *testing.T) {
	var orientation = "hetero"
	err := CheckOrientation(orientation)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", orientation, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", orientation)
	}
}

func TestOrientationOK_2(t *testing.T) {
	var orientation = "bi"
	err := CheckOrientation(orientation)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", orientation, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", orientation)
	}
}

func TestOrientationOK_3(t *testing.T) {
	var orientation = "homo"
	err := CheckOrientation(orientation)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", orientation, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", orientation)
	}
}

// ------------- ORIENTATION TESTING OF INVALID CASES --------------------------

func TestOrientationInvalid_1(t *testing.T) {
	var orientation = "люблю всех"
	err := CheckOrientation(orientation)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", orientation)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", orientation, err.Error())
	}
}

func TestOrientationInvalid_2(t *testing.T) {
	var orientation = ""
	err := CheckOrientation(orientation)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", orientation)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", orientation, err.Error())
	}
}

// ------------- BIOGRAPHY TESTING OF VALID CASES --------------------------

func TestBiographyOK_1(t *testing.T) {
	var biography = "Родился, потерпел, умер"
	err := CheckBio(biography)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", biography, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", biography)
	}
}

func TestBiographyOK_2(t *testing.T) {
	var biography = `Родился в 1989 году 
	поступил в университет
	закончил университет
	работал программистом
	изобрел машину времени
	умер в 1875`
	err := CheckBio(biography)
	if err != nil {
		t.Errorf(RED_BG+"FAILED: %s %s"+NO_COLOR+"\n", biography, err.Error())
	} else {
		t.Logf(GREEN_BG+"SUCCESS: %s"+NO_COLOR+"\n", biography)
	}
}

// ------------- BIOGRAPHY TESTING OF INVALID CASES --------------------------

func TestBiographyInvalid_1(t *testing.T) {
	var biography = `фывдофдфыощцййцшойцвлфыолдыфоафыдаорыфдлфыодлыфоыфвдлоыфвлдф
	овфылдвоыфлдвоыфвлдоывфлдвофылдвоывлфывоылфдыовлосчсчтсчсчьстчьсчимч
	ыршыфвоыфвфлдлвоывфрвфылвдфродвлыфвфдлвфдвлфвфл`
	err := CheckBio(biography)
	if err == nil {
		t.Errorf(RED_BG+"FAILED: '%s' should be error"+NO_COLOR+"\n", biography)
	} else {
		t.Logf(GREEN_BG+"SUCCESS: '%s' %s -- as it expected"+NO_COLOR+"\n", biography, err.Error())
	}
}
