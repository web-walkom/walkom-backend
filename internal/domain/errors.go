package domain

import "errors"

var (
	ErrSecretCodeInvalid = errors.New("Код не верный")
	ErrSecretCodeExpired = errors.New("Код просрочен")
	ErrUserNotFound      = errors.New("Пользователь не найден")

	ErrHeaderAuthorizedIsEmpty = errors.New("Пустой заголовок Authorized")
	ErrInvalidHeaderAuthorized = errors.New("Не верный заголовок Authorized")
	ErrTokenIsEmpty            = errors.New("Токен пустой")

	ErrNoAuthorized     = errors.New("Не авторизован")
	ErrParseToken       = errors.New("Ошибка при парсинге токена")
	ErrUnexpectedMethod = errors.New("Неожиданный способ подписи токена")
)
