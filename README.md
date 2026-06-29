# Subscription Service

REST API для управления пользовательскими подписками и расчета суммарной стоимости подписок за указанный период.

<details>
<summary><b>📑 Оглавление</b> (нажмите, чтобы развернуть)</summary>
<br>

- [Subscription Service](#subscription-service)
	- [🎯 Что этот проект демонстрирует](#-что-этот-проект-демонстрирует)
	- [📝 Возможности](#-возможности)
	- [⚠️ Ограничения](#️-ограничения)
	- [🛠️ Технологический стек](#️-технологический-стек)
	- [🚀 Запуск проекта](#-запуск-проекта)
		- [Development](#development)
		- [Production](#production)
	- [📄 Swagger](#-swagger)
	- [📸 Демонстрация работы](#-демонстрация-работы)
		- [▶️ Скринкаст (2 минуты)](#️-скринкаст-2-минуты)
		- [📸 Скриншоты](#-скриншоты)
	- [📦 Модель подписки](#-модель-подписки)
		- [Поля](#поля)
	- [📋 API](#-api)
		- [Healthcheck](#healthcheck)
			- [Ping](#ping)
		- [Подписки](#подписки)
			- [Создать подписку](#создать-подписку)
			- [Получить подписку](#получить-подписку)
			- [Получить список подписок](#получить-список-подписок)
			- [Обновить подписку](#обновить-подписку)
			- [Удалить подписку](#удалить-подписку)
		- [Подсчет стоимости подписок](#подсчет-стоимости-подписок)
			- [Получить суммарную стоимость](#получить-суммарную-стоимость)
	- [✅ Валидация](#-валидация)
		- [user\_id](#user_id)
		- [start\_date / end\_date](#start_date--end_date)
		- [price](#price)
	- [🧠 Структура проекта](#-структура-проекта)
	- [📝 Логирование](#-логирование)
	- [⚙️ Конфигурация](#️-конфигурация)
	- [🔨 Сборка](#-сборка)
	- [🧪 Тестирование](#-тестирование)
		- [Все тесты](#все-тесты)
		- [Unit-тесты](#unit-тесты)
		- [Integration-тесты](#integration-тесты)
	- [🧹 Линтинг](#-линтинг)
	- [🔧 Генерация кода](#-генерация-кода)
	- [📦 Миграции](#-миграции)
	- [🪝 Git Hooks](#-git-hooks)
	- [🤝 Вклад в проект](#-вклад-в-проект)
	- [📄 Лицензия](#-лицензия)

</details>

## 🎯 Что этот проект демонстрирует

- **Clean Architecture** в Go (cmd/internal)
- **Graceful shutdown** и обработка сигналов
- **Rate limiting** на уровне middleware
- **Integration tests** с реальной БД (не моками)
- **Structured logging** с уровнями (debug/info/error)
- **Многоокружечный деплой** (dev/prod через docker-compose)
- **Автоматическая генерация Swagger** из кода
- **Makefile как единая точка входа** для всей команды
- **Pre-commit хуки** с автоматическим добавлением локально через `make git-hooks`

---

## 📝 Возможности

* Создание подписки
* Получение подписки по ID
* Получение списка подписок с пагинацией и фильтрацией
* Обновление подписки
* Удаление подписки
* Подсчет суммарной стоимости подписок за выбранный период
* Swagger-документация
* PostgreSQL в качестве СУБД
* SQL-миграции
* Структурированное логирование
* Конфигурация через `.env`
* Docker Compose для локального запуска

---

## ⚠️ Ограничения

* Проверка существования пользователя не выполняется.
* Стоимость подписки хранится в целых рублях.
* Копейки не учитываются.
* Формат дат — `MM-YYYY`.
* Каждая подписка принадлежит одному пользователю.

---

## 🛠️ Технологический стек

* Go
* PostgreSQL
* Docker
* Docker Compose
* Goose
* Swaggo / Swagger
* GolangCI-Lint

---

## 🚀 Запуск проекта

### Development

```bash
make dev-up
```

Остановка:

```bash
make dev-down
```

---

### Production

```bash
make prod-up
```

Остановка:

```bash
make prod-down
```

---

## 📄 Swagger

Генерация Swagger-документации:

```bash
make swagger
```

После запуска приложения документация доступна по адресу:

```text
http://localhost:8080/swagger/index.html (Prod)
http://localhost:3000/swagger/index.html (Dev)
```
---

## 📸 Демонстрация работы

### ▶️ Скринкаст (2 минуты)
https://github.com/user-attachments/assets/dde1a15a-c659-4c1a-9383-d26fd5564385

**Что показано в видео:**
- ✅ Поднятие окружения через `make dev-up`
- ✅ Показ Swagger-документации
- ✅ Создание подписки через программу для тестирования API (Insomnia)
- ✅ Получение списка с фильтрацией и пагинацией
- ✅ Расчет суммарной стоимости за период
- ✅ Структурированное логирование в терминале

> 💡 **Видео без звука**, но с пояснениями в интерфейсе. Все эндпоинты документированы в Swagger.

---

### 📸 Скриншоты

 Создание подписки | Ошибка валидации | Список подписок |
|--------------|----------------------|------------|
| ![Создание подписки](https://raw.githubusercontent.com/gerilie/subscription-service/1dfa36d97f1ef21a9b1733f174e035befa697e06/create.png) | ![Ошибка валидации, Bad Request](https://raw.githubusercontent.com/gerilie/subscription-service/1dfa36d97f1ef21a9b1733f174e035befa697e06/bad-request.png) | ![Список подписок](https://raw.githubusercontent.com/gerilie/subscription-service/1dfa36d97f1ef21a9b1733f174e035befa697e06/list.png) |

---

## 📦 Модель подписки

```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025"
}
```

### Поля

| Поле         | Тип     | Описание                              |
| ------------ | ------- | ------------------------------------- |
| id           | integer | Идентификатор подписки                |
| service_name | string  | Название сервиса                      |
| price        | integer | Стоимость подписки в рублях           |
| user_id      | UUID    | Идентификатор пользователя            |
| start_date   | MM-YYYY | Дата начала подписки                  |
| end_date     | MM-YYYY | Дата окончания подписки (опционально) |

---

## 📋 API

### Healthcheck

#### Ping

```http
GET /ping
```

---

### Подписки

#### Создать подписку

```http
POST /subscriptions
```

Пример запроса:

```json
{
  "service_name": "Yandex Plus",
  "price": 400,
  "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
  "start_date": "07-2025"
}
```

---

#### Получить подписку

```http
GET /subscriptions/{id}
```

---

#### Получить список подписок

```http
GET /subscriptions?page=1&limit=20
```

Поддерживаемые фильтры:

| Параметр     | Описание                   |
| ------------ | -------------------------- |
| page         | Номер страницы             |
| limit        | Количество элементов       |
| service_name | Фильтр по названию сервиса |
| user_id      | Фильтр по пользователю     |

Пример:

```http
GET /subscriptions?page=1&limit=20&service_name=Yandex Plus
```

---

#### Обновить подписку

```http
PATCH /subscriptions/{id}
```

---

#### Удалить подписку

```http
DELETE /subscriptions/{id}
```

---

### Подсчет стоимости подписок

#### Получить суммарную стоимость

```http
GET /subscriptions/sum
```

Параметры:

| Параметр     | Обязательный | Формат  |
| ------------ | ------------ | ------- |
| start_date   | Да           | MM-YYYY |
| end_date     | Да           | MM-YYYY |
| service_name | Нет          | string  |
| user_id      | Нет          | UUID    |

Пример:

```http
GET /subscriptions/sum?start_date=01-2025&end_date=12-2025&user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba
```

Пример ответа:

```json
{
  "total_price": 4800
}
```

---

## ✅ Валидация

### user_id

Должен быть валидным UUID.

### start_date / end_date

Поддерживается формат:

```text
MM-YYYY
```

Примеры:

```text
01-2025
07-2025
12-2030
```

### price

* только целое число;
* значение должно быть больше 0.

---

## 🧠 Структура проекта

<img width="491" height="1775" alt="clean_architecture_diagram" src="https://github.com/user-attachments/assets/0b26eefa-de5d-4c99-937b-3a62ab03fe59" />

- `cmd/subscription` — точка входа
- `internal/` — бизнес-логика (не экспортируется)
- `migrations/` — версионирование схемы
- `docs/` — автогенерируемая Swagger-документация
- `pkg/` - библиотечные пакеты
---

## 📝 Логирование

Сервис ведет структурированное логирование:

* запуск приложения;
* HTTP-запросы;
* ошибки валидации;
* ошибки базы данных;
* бизнес-события.

---

## ⚙️ Конфигурация

Конфигурация хранится в `.env.dev` и `.env.prod`.

Пример:

```env
# App
APP_ENV=dev
APP_HOST=app
APP_PORT=8080

APP_READ_HEADER_TIMEOUT=5s
APP_READ_TIMEOUT=15s
APP_WRITE_TIMEOUT=30s
APP_IDLE_TIMEOUT=120s

APP_RATE_LIMIT_REQUESTS_PER_SECOND=5
APP_RATE_LIMIT_BURST=10
APP_RATE_LIMIT_CLEANUP_INTERVAL=3m
APP_RATE_LIMIT_CLEANUP_MAX_IDLE=15m

# DB
POSTGRES_IMAGE=postgres:18.4-bookworm
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=dev
POSTGRES_PASSWORD=dev
POSTGRES_DB=subscriptions

# Logger
LOGGER_LEVEL=debug
```

---

## 🔨 Сборка

```bash
make build
```

Будет создан бинарный файл:

```text
./subscription
```

---

## 🧪 Тестирование

### Все тесты

```bash
make test
```

### Unit-тесты

```bash
make test-unit
```

### Integration-тесты

```bash
make test-integration
```

---

## 🧹 Линтинг

```bash
make lint
```

---

## 🔧 Генерация кода

```bash
make gen
```

---

## 📦 Миграции

Создание новой миграции:

```bash
make migrate-gen name=create_subscriptions_table
```

Миграции расположены в каталоге:

```text
migrations/
```

Для работы используется Goose.

---

## 🪝 Git Hooks

Установка git hooks:

```bash
make git-hooks
```

---

## 🤝 Вклад в проект

Любой вклад приветствуется!

Перед созданием Pull Request, пожалуйста, ознакомься с [CONTRIBUTING.md](CONTRIBUTING.md).

---

## 📄 Лицензия
Этот проект распространяется под лицензией MIT.
