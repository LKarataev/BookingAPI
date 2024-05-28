# BookingAPI

В папке **task** находится описание задания и исходный код.

### Как запустить

Для запуска сервера команда `make run`.

Для запуска тестов команда `make test`.

### Слои приложения

1) **cmd/app/main.go** - входная точка в приложение
2) **internal/service** - предоставляет контракт сервиса типа http (обработчик http запросов - конвертирует входные данные из http в объекты данных, независимых от транспорта)
3) **internal/handlers** - бизнесс-логика, независимая от транспорта
4) **internal/dao** - слой доступа к данным - знает детали как устроено хранение
