# build:
# 	@go build -o bin/los_logger cmd/main.go

dev:
	@go run cmd/main.go

gen:
	@templ generate

build: gen
	@docker compose up --build -d