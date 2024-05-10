.PHONY: up down dev-up dev-down run test test-race clean

up:
	docker-compose --file ./build/docker-compose.yml up -d --remove-orphans

down:
	docker-compose --file ./build/docker-compose.yml down

dev-up:
	docker-compose --file ./build/docker-compose.yml up -d --build --remove-orphans

dev-down:
	docker-compose --file ./build/docker-compose.yml down --rmi all -v

run:
	CGO_ENABLED=0 go build -ldflags='-w -s' -o app ./cmd/qr-coder && HTTP_ADDR=:8080 ./app

test:
	go test -v -shuffle=on -count=2 -short -cover ./...

test-race:
	go test -race ./...

clean:
	rm app
