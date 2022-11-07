SERVICE=assistbot
HOST=tw1

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/${SERVICE} ./cmd/${SERVICE}

push:
	docker buildx build --platform=linux/amd64 --tag tinas .
	docker tag tinas:latest taverok/tinas:latest
	docker push taverok/tinas:latest

restart:
	ssh $(HOST) "docker compose -f /opt/services/tinas/docker-compose.yml up -d --no-deps --build --remove-orphans"

copy_db:
	scp ./db.sqlite $(HOST):/opt/services/tinas/db