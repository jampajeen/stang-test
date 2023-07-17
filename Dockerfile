FROM golang:1.19-alpine AS build
RUN apk update && apk --no-cache add gcc g++ make ca-certificates git
RUN mkdir -p /go/src/github.com/jampajeen/stang-test && chmod -R 777 /go/src/github.com/jampajeen/stang-test
RUN mkdir -p /scripts && chmod -R 777 /scripts
WORKDIR /go/src/github.com/jampajeen/stang-test

COPY go.mod .
COPY go.sum .

RUN go mod download

# COPY . .
COPY scripts scripts
COPY monitor-service monitor-service
COPY monitor-service/config.yml config.yml

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/bin .
COPY --from=build /go/src/github.com/jampajeen/stang-test/config.yml config.yml
