package errors

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrBadRequest         = errors.New("Некорректный запрос")
	ErrAlreadyAbbreviated = errors.New("Эта ссылка уже сокращена")
	ErrURLNotFound        = errors.New("URL не найден")
	ErrIncorrectStoreType = errors.New("Неизвестный тип хранилища")
)

var errorToCode = map[error]int{
	ErrBadRequest:         http.StatusBadRequest,
	ErrAlreadyAbbreviated: http.StatusBadRequest,
	ErrURLNotFound:        http.StatusNotFound,
}

func ConvertError(err error) int {
	result, ok := errorToCode[err]
	log.Println(ok)
	if ok {
		return result
	}
	return http.StatusInternalServerError
}
