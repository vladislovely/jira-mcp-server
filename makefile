# Makefile для сборки и управления Docker-контейнерами,
# основанными на предоставленных Dockerfile и docker-compose.yml.

APP_NAME := mcp-server
DOCKER_IMAGE := vladislove2k/jira-mcp-server:latest
DOCKET_TAG := v0.0.2
.PHONY: build docker-build docker-push docker-push compose-up compose-down clean

# Локальная сборка Go-приложения (без Docker).
build:
	go build -o $(APP_NAME) ./cmd/...

docker-build:
	docker compose build $(APP_NAME)

# Сборка Docker-образа (используя Dockerfile в текущей директории).
docker-build-and-up:
	docker compose build $(APP_NAME) && docker compose up -d $(APP_NAME)

# Запуск сервисов из docker-compose.yml в фоне (детаче-режим).
docker-up:
	docker compose up -d

# Остановка и удаление контейнеров, запущенных через docker-compose.
docker-down:
	docker compose down

docker-push:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)

# Очистка локально собранного бинарника.
clean:
	rm -f $(APP_NAME)

format:
	golines --max-len=100 -w ./

lint:
	 golangci-lint run ./...

compose-update:
	$(MAKE) compose-down && docker compose pull openwebui mcpo && docker compose build $(APP_NAME) && $(MAKE) compose-up

git-tag-and-push:
	git tag $(DOCKET_TAG)
	git push origin $(DOCKET_TAG)