swag-go:
		swag init && swag fmt -d ./ --exclude ./internal && go run .\main.go

DEFAULT_GOAL := swag-go