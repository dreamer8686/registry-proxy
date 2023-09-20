#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/registry-proxy main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/arm64/registry-proxy main.go

docker buildx use gobuilder || docker buildx create --use --name gobuilder
docker buildx build --push --platform linux/amd64,linux/arm64 -t ketches/registry-proxy:latest .