FROM golang:1.18-alpine AS build
WORKDIR /src
COPY . .
RUN go build

FROM alpine
COPY --from=build /src/simple-http-prometheus /usr/local/bin/simple-http-prometheus
CMD ["simple-http-prometheus"]
