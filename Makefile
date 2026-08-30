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
	docker exec -u root -it app_filter bash -c 'go install github.com/swaggo/swag/cmd/swag@latest || swag init -g app/cmd/api/main.go -o app/internal/docs' \
	&& sudo chown -R walter:walter app/internal/docs

go-run:
	docker exec -u root -it app_filter bash -c 'go run ./app/cmd/api'

