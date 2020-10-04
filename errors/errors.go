package errors

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Code               int    `json:"code"`
	HttpResponseStatus int    `json:"-"`
	RuPattern          string `json:"-"`
	RuToClient         string `json:"ru"`
	EngPattern         string `json:"-"`
	EngToClient        string `json:"eng"`
}

type ApiErrorArgument struct {
	ru  string
	eng string
	err error
}

func NewArg(ru string, eng string) ApiErrorArgument {
	var err ApiErrorArgument

	err.ru = ru
	err.eng = eng
	return err
}

func (err ApiErrorArgument) Error() string {
	if err.err != nil {
		return err.eng + " " + err.err.Error()
	}
	return err.eng
}

func (err ApiErrorArgument) AddOriginalError(newErr error) ApiErrorArgument {
	err.err = newErr
	return err
}

func new(code int, httpStatus int, ruPattern string, engPattern string) ApiError {
	var err ApiError

	err.Code = code
	err.HttpResponseStatus = httpStatus
	err.RuPattern = ruPattern
	err.RuToClient = ruPattern
	err.EngPattern = engPattern
	err.EngToClient = engPattern

	return err
}

func (err ApiError) WithArguments(arg ...interface{}) ApiError {
	if len(arg) == 1 {
		if argument, ok := arg[0].(ApiErrorArgument); ok {
			err.RuToClient = err.RuPattern + " " + argument.ru
			err.EngToClient = err.EngPattern + " " + argument.eng
			return err
		}
	}
	if len(arg) == 2 {
		ruArgument, ruOk := arg[0].(string)
		engArgument, engOk := arg[1].(string)
		if ruOk && engOk {
			err.RuToClient = err.RuPattern + " - " + ruArgument
			err.EngToClient = err.EngPattern + " - " + engArgument
			return err
		}
	}
	println("\033[41m ApiError Arguments() - was found wrong argument \033[m\n")
	return err
}

func (err ApiError) ToJson() []byte {
	dst, _ := json.Marshal(err)
	return dst
}

func (err ApiError) Error() string {
	return err.EngToClient
}

func (e ApiError) IsOverlapWithError(err error) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(ApiError); ok {
		if apiErr.Code == e.Code {
			return true
		}
	}
	return false
}

var (
	// Common errors (Code 1000 - 1999)
	RecordNotFound = new(1000, http.StatusUnprocessableEntity, // 422
		"Такой записи не существует в базе данных",
		"Record not found in database")

	// User errors (Code 2000 - 2999)
	UserNotLogged = new(2000, http.StatusUnauthorized, // 401
		"Пользователь не авторизован",
		"User not authorized")
	NotConfirmedMail = new(2001, http.StatusUnauthorized, // 401 - не нашел более подходящего статуса
		"Пожалуйста, подтвердите вашу почту. Письмо выслано на ваш почтовый ящик",
		"Please confirm your mail. Mail was sent to your email address")
	UserNotExist = new(2002, http.StatusUnauthorized, // 401 - не нашел более подходящего статуса
		"Пользователь не существует",
		"User not exists")
	AuthFail = new(2003, http.StatusUnprocessableEntity, // 422
		"Не могу авторизовать пользователя. Неверная почта или пароль",
		"Cannot authorize user. Wrong mail or password")
	RegFailUserExists = new(2004, http.StatusNotAcceptable, //406
		"Такой пользователь уже существует",
		"Same user already exists")

	// Request errors
	InvalidRequestBody = new(4000, http.StatusBadRequest, // 400
		"Тело запроса содержит ошибку",
		"Request body is invalid")
	NoArgument = new(4002, http.StatusBadRequest, // 400
		"Отстутствует одно из обязательных полей",
		"One of the required fields is missing")
	InvalidArgument = new(4003, http.StatusUnprocessableEntity, // 422
		"Ошибка в аргументе",
		"Argument error")

	// Internal errors
	DatabaseError = new(5000, http.StatusInternalServerError, // 500
		"База данных вернула ошибку",
		"Database returned error")
	WebSocketError = new(5001, http.StatusInternalServerError, // 500
		"Произошла ошибка веб сокета",
		"Websocket error")
	MarshalError = new(5002, http.StatusInternalServerError, // 500
		"Произошла ошибка при упаковке данных",
		"An error occurred while packing data")
	UnmarshalError = new(5003, http.StatusInternalServerError, // 500
		"Произошла ошибка при распаковке данных",
		"An error occurred while unpacking data")
	UnknownInternalError = new(5004, http.StatusInternalServerError, // 500
		"Произошла неизвестная ошибка",
		"An unknown error occurred")

	// Error arguments for the package of database queries
	DatabasePreparingError = NewArg("ошибка во время подготовки к запросу",
		"error during preparing")
	DatabaseExecutingError = NewArg("ошибка во время выполнения запроса",
		"error during executing query")
	DatabaseQueryError = NewArg("ошибка во время выполнения запроса",
		"error during executing query")
	DatabaseScanError = NewArg("ошибка во время парсинга параметров",
		"error during scaning")
)
