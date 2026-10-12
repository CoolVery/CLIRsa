package apperr

import "errors"

var (
	ErrRsaLen = errors.New("Длина RSA ключа не составляет 40 символов")
	ErrNotFilter = errors.New("Не задано ни одного фильтра для запроса")
	ErrNotExistInDB = errors.New("Не найдено записи в базе данных")
)