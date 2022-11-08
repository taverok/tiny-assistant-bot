SERVICE=assistbot
HOST=tw1

build:
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/${SERVICE} ./cmd/${SERVICE}

push:
	docker buildx build --platform=linux/amd64 --tag tinas .
	docker tag tinas:latest taverok/tinas:latest
	docker push taverok/tinas:latest

restart:
	ssh $(HOST) " cd /opt/services/tinas/ && \
				docker compose down && \
				docker compose rm -f && docker compose pull && \
				docker compose up -d --no-deps --build --force-recreate"

push_db:
	scp ./db.sqlite $(HOST):/opt/services/tinas/db

pull_db:
	scp  $(HOST):/opt/services/tinas/db/db.sqlite ./db.sqlite