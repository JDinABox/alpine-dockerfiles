build:
	go build -ldflags="-s -w" -o ./app.so ./

docker-build:
	docker build --tag wireguard-alpine:latest . 