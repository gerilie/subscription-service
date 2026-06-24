# Subscription Service

REST API для управления пользовательскими подписками и расчета суммарной стоимости подписок за указанный период.

## Возможности

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

## Модель подписки

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

## API

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

## Валидация

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

## Технологический стек

* Go
* PostgreSQL
* Docker
* Docker Compose
* Goose
* Swaggo / Swagger
* GolangCI-Lint

---

## Структура проекта

```text
.
├── cmd/
│   └── subscription/
├── internal/
├── migrations/
├── docs/
├── docker-compose.yml
├── docker-compose.dev.yml
├── docker-compose.prod.yml
├── Dockerfile
├── Makefile
├── .env.dev
├── .env.prod
└── README.md
```

---

## Конфигурация

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

## Запуск проекта

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

## Сборка

```bash
make build
```

Будет создан бинарный файл:

```text
./subscription
```

---

## Тестирование

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

## Линтинг

```bash
make lint
```

---

## Генерация кода

```bash
make gen
```

---

## Миграции

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

## Swagger

Генерация Swagger-документации:

```bash
make swagger
```

После запуска приложения документация доступна по адресу:

```text
http://localhost:8080/swagger/index.html
```

---

## Git Hooks

Установка git hooks:

```bash
make git-hooks
```

---

## Логирование

Сервис ведет структурированное логирование:

* запуск приложения;
* HTTP-запросы;
* ошибки валидации;
* ошибки базы данных;
* бизнес-события.

---

## Ограничения

* Проверка существования пользователя не выполняется.
* Стоимость подписки хранится в целых рублях.
* Копейки не учитываются.
* Формат дат — `MM-YYYY`.
* Каждая подписка принадлежит одному пользователю.

---

## Полезные команды

```bash
# Development
make dev-up
make dev-down

# Production
make prod-up
make prod-down

# Build
make build

# Tests
make test
make test-unit
make test-integration

# Lint
make lint

# Generate code
make gen

# Generate Swagger
make swagger

# Create migration
make migrate-gen name=create_table

# Install git hooks
make git-hooks
```
