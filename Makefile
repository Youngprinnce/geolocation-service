.PHONY: run test docker-up

# Run the application locally
run:
	go run main.go server --config config-local.yaml

# Test the application
test:
	go test -v ./...

# Run application using Docker Compose
docker-up:
	docker-compose up --build
