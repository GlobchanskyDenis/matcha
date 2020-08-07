package errDef

import (
	"errors"
)

var (
	RecordNotFound = errors.New("Такой записи не существует в базе данных")
)

func IsRecordNotFoundError(err error) bool {
	if err == RecordNotFound {
		return true
	}
	return false
}