package domain

import "errors"

var (
	ErrServer = errors.New("Ошибка на стороне сервера")
	ErrInvalidInput = errors.New("Не верные входные данные")

	ErrGetAllExcusions = errors.New("Ошибка при получении экскурсий")

	ErrGetUserById = errors.New("Ошибка при получении пользователя")
	ErrUserNotFound      = errors.New("Пользователь не найден")

	ErrSendCodeEmail = errors.New("Ошибка при отправке кода на почту")
	ErrCheckCodeEmail = errors.New("Ошибка при проверке кода")
	ErrCreateUser = errors.New("Ошибка при создании аккаунта пользователя")
	ErrGetUser = errors.New("Ошибка при получении даных пользователя")
	ErrCreateSession = errors.New("Ошибка при создании новой сессии")
	ErrSecretCodeInvalid = errors.New("Код не верный")
	ErrSecretCodeExpired = errors.New("Код просрочен")

	ErrHeaderAuthorizedIsEmpty = errors.New("Пустой заголовок Authorized")
	ErrInvalidHeaderAuthorized = errors.New("Не верный заголовок Authorized")
	ErrTokenIsEmpty            = errors.New("Токен пустой")
	ErrNoAuthorized     = errors.New("Не авторизован")
	ErrParseToken       = errors.New("Ошибка при парсинге токена")
	ErrUnexpectedMethod = errors.New("Неожиданный способ подписи токена")
)
