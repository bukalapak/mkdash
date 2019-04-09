.PHONY: test dep compile

include .env

export VERSION            ?= $(shell git show -q --format=%h)
export $(shell sed 's/=.*//' .env)

test:
	go test ./...

dep:
	go mod tidy

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-X main.GrafanaAuth=$$GRAFANA_AUTH" -o mkdash app/cmd/main.go;)
