Реализован REST-сервис для агрегации данных об онлайн
подписках пользователей.

## Продакшен
1. Запустить `make prod-up`
2. Посетить [api](http://localhost:8080/swagger/index.html)

## Разработка
1. Запустить `make dev-up`
3. Посетить [api](http://localhost:3000/swagger/index.html)

## Переменные окружения
### Логгер
- level
   - debug *(dev по-умолчанию)*
   - info *(prod по-умолчанию)*
   - warn
   - error
   - fatal
