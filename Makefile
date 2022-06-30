run:
	go run ./...

docker/build:
	docker build -t henrod/simple-http-prometheus:$(TAG) .