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

## Миграции

Для поднятия миграций необходимо ввести команду:
```
go tool goose postgres "СТРОКА_ПОДКЛЮЧЕНИЯ_К_БД" up -dir migrations
```
Для отката миграций необходимо ввести команду:
```
go tool goose postgres "СТРОКА_ПОДКЛЮЧЕНИЯ_К_БД" down -dir migrations
```

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
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Type string `json:"type"`
}
```
### Brand
```go
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
```
### NewProduct
```go
type NewProduct struct {
	Name       string    `json:"name,omitempty"`
	Image      *string   `json:"image,omitempty"`
	Table      *string   `json:"table,omitempty"`
	Categories *[]string `json:"categories,omitempty"`
	Brand      *int      `json:"brand,omitempty"`
}

```
## Зависимости 

| Название        | Ссылка                                    |
|-----------------|-------------------------------------------|
|Goose            | [Ссылка](https://github.com/pressly/goose)|