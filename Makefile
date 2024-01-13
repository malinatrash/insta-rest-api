.PHONY: run docs go

run:
	go run cmd/app/main.go

docs:
	swag init -g cmd/app/main.go

go:
	swag init -g cmd/app/main.go
	go run cmd/app/main.go migrate
	swag init -g cmd/app/main.go
