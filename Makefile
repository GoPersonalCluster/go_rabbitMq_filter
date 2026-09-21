build:
	docker-compose up -d
bash-app:
	docker exec -it app_filter bash
reset-docker:
	@docker stop $$(docker ps -aq) 2>/dev/null || true
	@docker rm $$(docker ps -aq) 2>/dev/null || true
	@docker rmi $$(docker images -aq) 2>/dev/null || true
	@docker volume rm $$(docker volume ls -q) 2>/dev/null || true
	@docker network prune -f

swagger-init:
	docker exec -u root -w /app/app/internal/handler  -it app_filter bash -c 'swag init -g ../../cmd/api/main.go  --dir . --parseInternal --parseDependency -o ../docs' \
	&& sudo chown -R walter:walter app/internal/docs

run:
	docker exec -u root -it app_filter bash -c 'go run ./app/cmd/api'
debug:
	docker exec -u root -it app_filter bash -c \
	'dlv debug ./app/cmd/api --headless --listen=0.0.0.0:2345 --api-version=2 --accept-multiclient --build-flags="-buildvcs=false"'
