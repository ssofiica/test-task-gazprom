# Тестовое задание
Сервис для поздравлений с днем рождения

## Запуск проекта

```git clone https://github.com/ssofiica/test-task-gazprom.git```

В корне проекта <br>
```
cd integration
docker-compose up
```
Это запустит контейнер с постгрес и редис<br>
Для запуска сервиса ```go run cmd/main.go```

# API
/api/v1

## Авторизация<br>
### POST &nbsp;&nbsp;&nbsp;&nbsp;/signup
Регистрация нового пользователя в системе

**body** - имя, фамилия, почта, пароль, дата рождения (формат - год-месяц-день)<br>
```{"name":"София", "surname":"Валова", "email":"s@mail.ru", "password":"1234567890", "birthday":"2003-04-26"}```

**response:** <br>
200 - ```{"id":1,"name":"София", "surname":"Валова", "email":"s@mail.ru"}```<br>
400 - ```{"error":"Предоставлены некорректные данные"}```<br>
403 - ```{"error":"Вы уже зарегистрированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```

### POST &nbsp;&nbsp;&nbsp;&nbsp;/signin
Авторизация пользователя в системе

**body** - почта, пароль<br>
```{"email":"s@mail.ru", "password":"1234567890"}```

**response:** <br>
200 - ```{"id":1,"name":"София", "surname":"Валова", "email":"s@mail.ru"}```<br>
400 - ```{"error":"Предоставлены некорректные данные"}```<br>
400 - ```{"error":"Неверный адрес почты"}```<br>
400 - ```{"error":"Неверный пароль"}```<br>
403 - ```{"error":"Вы уже авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```

### POST &nbsp;&nbsp;&nbsp;&nbsp;/signout
Деавторизация пользователя в системе

**response:** <br>
200 - ```{"detail":"Сессия успешно завершена""}```<br>
403 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```

## Пользователь
Для доступа к этим методам необходима авторизация
### GET &nbsp;&nbsp;&nbsp;&nbsp;/user/all<br>
Список всех пользователей

**response:** <br>
200
```json
[
    {"id":1,"name":"Иван", "surname":"Иванов", "email":"i@mail.ru", "birthday":"2000-02-06"},
    {"id":2,"name":"София", "surname":"Валова", "email":"s@mail.ru", "birthday":"2003-04-25"}
]
```
401 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```

### GET &nbsp;&nbsp;&nbsp;&nbsp;/user/search<br>
Поиск пользователей по имени и фамилии

**body** - имя, фамилия<br>
```{"name":"София", "surname":"Валова"}```

**response:** <br>
200
```json
[
    {"id":1,"name":"София", "surname":"Валова", "email":"sofia@mail.ru", "birthday":"2000-02-06"},
    {"id":2,"name":"София", "surname":"Валова", "email":"s@mail.ru", "birthday":"2003-04-25"}
]
```
400 - ```{"error":"Предоставлены некорректные данные"}```<br>
401 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```

### POST &nbsp;&nbsp;&nbsp;&nbsp;/user/subscribe/:id<br>
Подписка на оповещение о дне рождении
slug-параметр id - id пользователя, на которых надо подписаться

**response:** <br>
200 - ```{"detail":"Подписка успешно оформлена"}```
400 - ```{"error":"Параметр должен быть числовым"}```<br>
400 - ```{"error":"Такого пользователя нет"}```
401 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```
500 - ```{"error":"Не удалось оформить подписку на день рождение"}```

### DELETE &nbsp;&nbsp;&nbsp;&nbsp;/user/unsubscribe/:id<br>
Отмена подписк на оповещение о дне рождении
slug-параметр id - id пользователя, от которого надо отписаться

**response:** <br>
200 - ```{"detail":"Подписка успешно отменена"}```
400 - ```{"error":"Параметр должен быть числовым"}```<br>
400 - ```{"error":"Такого пользователя нет"}```
401 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```
500 - ```{"error":"Не удалось отменить подписку на день рождение, возможно вы не подписаны"}```

### GET &nbsp;&nbsp;&nbsp;&nbsp;/user/notification<br>
Список людей, на которых подписан пользователь и у которых сегодня день рождения

**response:** <br>
200
```json
[
    {"id":1,"name":"София", "surname":"Валова", "email":"sofia@mail.ru", "birthday":"2000-02-06"},
    {"id":2,"name":"Мария", "surname":"Марьева", "email":"maria@mail.ru", "birthday":"2000-02-06"}
]
```
401 - ```{"error":"Вы не авторизированы"}```<br>
500 - ```{"error":"Ошибка сервера"}```


### ER-диаграмма
```mermaid
erDiagram
    BIRTHDAY_SUBSCRIBING||--o{ USER : subscribes
    BIRTHDAY_SUBSCRIBING||--o{ USER : has-birthday
    BIRTHDAY_SUBSCRIBING {
        int birthday_user_id PK, FK
        int subscribing_user_id PK, FK
    } 
    USER{
        int id PK
        text name
        text surname
        text email
        text password
        date birthday
    }
```
