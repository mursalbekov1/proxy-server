DOCKER_COMPOSE_FILE = docker-compose.yml

SERVICE_NAME = app

build:
	docker-compose -f "./docker-compose.yml" build

up:
	docker-compose -f "./docker-compose.yml" up -d

down:
	docker-compose -f "./docker-compose.yml" down

logs:
	docker-compose -f "./docker-compose.yml" logs -f $(SERVICE_NAME)

restart:
	docker-compose -f "./docker-compose.yml" restart $(SERVICE_NAME)

clean:
	docker-compose -f "./docker-compose.yml" down -v --rmi all --remove-orphans

run:
	make build
	make up

rebuild:
	make clean
	make build
	make up

.PHONY: build up down logs restart clean run rebuild
