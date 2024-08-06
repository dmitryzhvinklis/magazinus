# Magazinus

**Magazinus** - это микросервисное приложение для магазина, написанное на Go. В текущей версии приложение включает базовый сервис для подключения к базе данных PostgreSQL и инициализации необходимых таблиц.

## Описание

Проект состоит из следующих компонентов:

- **База данных**: PostgreSQL.
- **Микросервис**: Подключается к базе данных, выполняет миграции и инициализирует таблицы.

## Установка

### Требования

- [Go](https://golang.org/dl/) версии 1.18 и выше.
- [Docker](https://www.docker.com/get-started) (опционально, для запуска базы данных).

### Клонирование репозитория

```bash
git clone https://github.com/yourusername/magazinus.git
cd magazinus
```

Сборка и запуск контейнера:

bash

    make

Сборка Docker-образа:

bash

    make docker-build

Запуск Docker-контейнера:

bash

    make docker-run

Остановка и удаление Docker-контейнера:

bash

    make docker-stop

Очистка Docker-ресурсов:

bash

    make clean



ЗАПУСКАЕТСЯ НА 8000