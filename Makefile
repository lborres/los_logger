# build:
# 	@go build -o bin/los_logger cmd/main.go

dev:
	@go run cmd/main.go

build:
	@docker compose up --build -d

gen:
	@templ generate