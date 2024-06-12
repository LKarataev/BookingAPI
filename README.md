# BookingAPI

BookingAPI — это прототип бэкэнд сервиса для бронирования отелей, написанный на языке Go. Этот проект демонстрирует мои навыки в разработке на Go и может служить основой для реального приложения для управления бронированиями. В настоящее время приложение использует память для хранения данных, однако архитектура легко позволяет перейти на использование внешнего хранилища.

## Запуск

Чтобы запустить сервер, выполните команду:

```sh
make run
```

Для запуска тестов используйте:

```sh
make test
```

## Пример использования

Чтобы создать бронирование, выполните команду:

```sh
curl --location --request POST 'localhost:8080/orders' \
--header 'Content-Type: application/json' \
--data-raw '{
    "hotel_id": "reddison",
    "room_id": "lux",
    "email": "guest@mail.ru",
    "from": "2024-01-02T00:00:00Z",
    "to": "2024-01-04T00:00:00Z"
}'
```

## Структура проекта

- **cmd/app/main.go** — входная точка приложения.
- **internal/service** — реализация сервисного слоя HTTP, обработка запросов и преобразование данных.
- **internal/handlers** — бизнес-логика, независимая от транспортного слоя.
- **internal/dao** — доступ к данным и работа с хранилищем.

## Зависимости

Проект использует стандартную библиотеку Go и HTTP роутер Chi.

## Установка

Клонируйте репозиторий:

```sh
git clone https://github.com/LKarataev/BookingAPI.git
```

Перейдите в директорию проекта:

```sh
cd BookingAPI
```

Установите зависимости:
```sh
go mod tidy
```
