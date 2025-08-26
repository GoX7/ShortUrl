# ShortUrl 🔗

ShortUrl is a URL shortening service built with Golang, Gin, PostgreSQL, and REST API. It features user authentication, logging, Docker support, and encrypted tokens for secure operations.

ShortUrl — это сервис сокращения URL-адресов, созданный с использованием Golang, Gin, PostgreSQL и REST API. Он включает аутентификацию пользователей, ведение журналов, поддержку Docker и зашифрованные токены для безопасной работы.

---

## Features / Функционал

| Endpoint           | Method | Description (EN)                   | Описание (RU)                     | Example / Пример                                                                                                                                      |
| ------------------ | ------ | ---------------------------------- | --------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| `/status/server`   | GET    | Check server status                | Проверка статуса сервера          | `curl http://localhost:8080/status/server`                                                                                                            |
| `/status/postgres` | GET    | Check PostgreSQL connection        | Проверка подключения к PostgreSQL | `curl http://localhost:8080/status/postgres`                                                                                                          |
| `/auth/register`   | POST   | Register a new user                | Регистрация нового пользователя   | `bash curl -X POST http://localhost:8080/auth/register -H "Content-Type: application/json" -d '{"username":"test","password":"1234"}'`                |
| `/auth/login`      | POST   | Log in                             | Вход пользователя                 | `bash curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d '{"username":"test","password":"1234"}'`                   |
| `/l`               | GET    | Get all links for the current user | Получить все ссылки пользователя  | `bash curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/l`                                                                                |
| `/l/:id`           | GET    | Search link by ID                  | Получить ссылку по ID             | `bash curl -H "Authorization: Bearer <TOKEN>" http://localhost:8080/l/123`                                                                            |
| `/l`               | POST   | Register a new link                | Зарегистрировать новую ссылку     | `bash curl -X POST -H "Authorization: Bearer <TOKEN>" -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/l` |

---

## Key Features / Основные особенности

* Middleware for **authentication** and **logging**
  Middleware для **авторизации** и **логирования**
* **AES-encrypted tokens** for secure authentication
  **AES-шифрование токенов** для безопасной аутентификации
* Fully **RESTful API**
  Полностью **REST API**
* **Docker** and **Docker Compose** ready
  Поддержка **Docker** и **Docker Compose**
* Automation support via **Makefile**
  Автоматизация через **Makefile**

---

## Tech Stack / Технологии

* **Golang**
* **Gin** (Web framework / Фреймворк)
* **PostgreSQL** (Database / База данных)
* **Docker / Docker Compose**
* **Makefile** for automation

---

## Quick Start / Быстрый старт

1. Clone the repository / Клонируем репозиторий:

```bash
git clone https://github.com/YourUsername/ShortUrl.git
cd ShortUrl
```

2. Configure `docker compose env` / Настройте `docker compose env` с параметрами PostgreSQL и AES ключом.

3. Run the service using make / Запуск сервиса через make:

```bash
make build
```

4. Server will be available at `http://localhost:8080` / Сервер будет доступен по адресу `http://localhost:8080`.

---

## Logging / Логирование

All requests and errors are logged, making it easy to debug and monitor the service.
Все запросы и ошибки логируются для удобного дебага и мониторинга сервиса.
