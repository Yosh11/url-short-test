swag-go:
		swag init && swag fmt -d ./ --exclude ./internal && go run .\main.go
up:
		docker compose up --build
down:
		docker compose down -v --rmi local

DEFAULT_GOAL := swag-go