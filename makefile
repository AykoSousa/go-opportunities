.PHONY: default run build test docs clean

# Variables
default:
	@swag init
	@go run main.go

run:
	@go run main.go

build:
	@go build -o $(APP_NAME) main.go

test:
	@go test ./ ...

docs:
	@swag init

clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs