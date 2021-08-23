
full: envs up wait scheme wait2 data

fulldb: envs db wait scheme wait2 data

envs:
	@[ -f ./.env ] && echo "==========> .env file already exists, their values will be used"|| cp env.example .env

wait:
	@echo "Waiting ..."
	@sleep 10

wait2:
	@echo "Waiting ..."
	@sleep 2

up:
	docker compose up -d

echo:
	@echo "Envs"
	@echo $(shell docker compose ps -q db)
	@echo $(shell docker compose ps -q app)

scheme:
	docker exec -i $(shell docker compose ps -q db) mysql -uroot todosdb < ./resources/mysql/schema.sql

data:
	docker exec -i $(shell docker compose ps -q db) mysql -uroot todosdb < ./resources/mysql/data.sql

clear:
	@echo "WARNING! docker compose services will be removed, also volumes data"
	docker compose down -v --remove-orphans

db:
	docker compose up -d db

#app:
#	docker compose up -d app

start:
	docker compose start