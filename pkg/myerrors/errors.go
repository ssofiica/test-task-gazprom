package myerrors

import (
	"errors"
)

var (
	Authorized       = errors.New("Вы уже авторизированы")
	Unauthorized     = errors.New("Вы не авторизованы")
	Registered       = errors.New("Вы уже зарегистрированы")
	InternalServer   = errors.New("Ошибка сервера")
	ParametrIsNumber = errors.New("Параметр должен быть числовым")
	BadCredentials   = errors.New("Предоставлены некорректные данные")

	WrongPassword        = errors.New("Неверный пароль")
	NoUser               = errors.New("Такого пользователя нет")
	NoCreatingUser       = errors.New("Не удалось зарегистрировать пользователя")
	NoSubcribeBdayUser   = errors.New("Не удалось оформить подписку на день рождение")
	NoUnsubcribeBdayUser = errors.New("Не удалось отменить подписку на день рождение, возможно вы не подписаны")
)
