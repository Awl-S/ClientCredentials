# OAuth Server

OAuth сервер на основе Go и фреймворка Gin для выдачи access token и доступа к защищённым ресурсам с проверкой токена.

## Описание

Сервер позволяет клиенту запрашивать access token с помощью метода POST и использовать его для получения доступа к защищённым ресурсам через метод GET. Основная функциональность реализована в следующих компонентах:

- **TokenHandler** — проверяет `client_id` и `client_secret`, и если они совпадают с зарегистрированными данными, выдает access token.
- **ResourceHandler** — проверяет наличие валидного access token в заголовке `Authorization` и предоставляет доступ к защищённым данным, если токен корректен.

## Структура проекта

1. **`main.go`** — Основной файл с определением обработчиков запросов и логикой OAuth сервера.
2. **`OAuthServer`** — Структура для хранения зарегистрированных клиентов и активных токенов доступа.

## Установка

### Требования

- Go 1.18+.
- Gin Framework для работы с HTTP запросами.

### Шаги для установки

1. Клонируйте репозиторий:
    ```bash
    git clone https://github.com/yourcompany/oauth-server.git
    ```

2. Перейдите в директорию проекта:
    ```bash
    cd oauth-server
    ```

3. Установите зависимости:
    ```bash
    go mod tidy
    ```

4. Запустите сервер:
    ```bash
    go run main.go
    ```

Сервер будет доступен по адресу `http://localhost:8080`.

## Использование

### 1. Получение access token

Запросите токен с помощью POST-запроса на `/token`, указав `client_id` и `client_secret`:
```powershell
Invoke-WebRequest -Uri "http://localhost:8080/token" -Method POST -Body "client_id=my_client_id&client_secret=my_client_secret"
