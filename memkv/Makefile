build:
	docker build -t fronzec/memkv:latest .

down:
	docker-compose down --remove-orphans

up: down build
	docker-compose up
