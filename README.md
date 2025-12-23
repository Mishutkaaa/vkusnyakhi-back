# Vkusnyakhi back
Сервервая часть сайта с вкусняшками
## Запуск

1. Копирование репозитория `git clone https://github.com/Mishutkaaa/vkusnyakhi-back.git`
2. Установка зависимостей `go mod tidy`

## Конфигурация

Конфигурация указывается в переменных окружения или в файле `.env`. Если стандартное значение не указано, значит параметр является обязательным.

| Название        | Описание                          | Стандартное значение |
|-----------------|-----------------------------------|----------------------|
|POSTGRES_PORT    | Порт БД                           | 5432                 |
|POSTGRES_HOST    | Хост БД                           | localhost            |
|POSTGRES_USER    | Пользователь БД                   |                      |
|POSTGRES_PASSWORD| Пароль БД                         |                      |
|POSTGRES_DB      | Имя БД                            |                      |
|PORT             | Порт приложение                   | 8080                 |

## Модели

### Drinks
```go
type Drinks struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Image    *string `json:"image,omitempty"`
	Category *[]int  `json:"category,omitempty"`
	Brand    *string `json:"brand,omitempty"`
}
```
### Food
```go
type Food struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Image    *string `json:"image,omitempty"`
	Category *[]int  `json:"category,omitempty"`
	Brand    *string `json:"brand,omitempty"`
}
```
### Categories
```go
type Categories struct {
	ID    int 
	Name  string
	Table string
}
```
### Brand
```go
type Brand struct {
	ID   int
	Name string
}
```
## Зависимости 

| Название        | Ссылка                                    |
|-----------------|-------------------------------------------|
|Goose            | [Ссылка](https://github.com/pressly/goose)|