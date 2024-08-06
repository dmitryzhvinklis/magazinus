# Установки
DOCKER_CMD=docker
DOCKER_COMPOSE_CMD=docker-compose
IMAGE_NAME=magazinus_db_service
CONTAINER_NAME=magazinus_db_service

# Цели Makefile

.PHONY: all build run stop clean docker-build docker-run docker-stop

# Основная цель: сборка и запуск контейнера
all: docker-build docker-run

# Собирает Docker-образ
docker-build:
	$(DOCKER_CMD) build -t $(IMAGE_NAME) .

# Запускает Docker-контейнер
docker-run:
	$(DOCKER_CMD) run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

# Останавливает и удаляет Docker-контейнер
docker-stop:
	$(DOCKER_CMD) stop $(CONTAINER_NAME)
	$(DOCKER_CMD) rm $(CONTAINER_NAME)

# Очищает Docker-образы и контейнеры
clean:
	$(DOCKER_CMD) system prune -a -f

