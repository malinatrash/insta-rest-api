.PHONY: run migrate docs

run:
	go run cmd/app/main.go

migrate:
	go run cmd/app/main.go migrate

docs:
	swag init -g cmd/app/main.go

go:
	swag init -g cmd/app/main.go
	go run cmd/app/main.go migrate
	swag init -g cmd/app/main.go
