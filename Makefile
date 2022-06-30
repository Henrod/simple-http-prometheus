run:
	go run ./...

docker/build:
	docker build -t henrod/simple-http-prometheus:$(TAG) .
	@read -p "Push the built image? [y/N] " push; \
  	if [ "$$push" == 'y' ]; then docker push henrod/simple-http-prometheus:$(TAG); fi
