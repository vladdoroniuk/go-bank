include .env
export $(shell sed 's/=.*//' .env)

service = auth

build:
	@go build -o ./bin/${service}

run: build
	@./bin/${service}

test:
	@go test -v ./...