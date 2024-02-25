.PHONY: simulator

unit-test:
	go generate ./...
	go test ./...
build:
	docker build -t simulator:latest .  
setup:
	docker-compose up -d
