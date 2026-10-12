package apperr

import "errors"

var (
	ErrRsaLen = errors.New("Длина RSA ключа не составляет 40 символов")
	ErrNotFilter = errors.New("Не задано ни одного фильтра для запроса")
	ErrUserNotExistInDB = errors.New("Не найден пользователь в базе данных")
	ErrContextCanceled = errors.New("Котнекст был отменен")
	ErrContextTimeout = errors.New("Конекст не успел выполниться за таймер")
	ErrContextDeadline = errors.New("Контекст не успел выполнить до дедлайна")
)