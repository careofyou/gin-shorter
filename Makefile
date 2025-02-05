build:
	go build -o myapp main.go

start:
	go run main.go

restart: build start

up:
	sudo docker-compose up --build
