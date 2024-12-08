#!/bin/bash

version=${1:-dev}
now=`date -u +%Y-%m-%dT%H:%M:%S`

set -e

echo "Running JS unit tests..."
jest test --config jest.config.js

echo "Downloading GO dependencies..."
go get ./...

go mod tidy

echo "Running GO unit tests..."
go test ./...

echo "Building application..."
go build -a -ldflags "-X main.BuildVersion=$version -X main.BuildTime=$now -extldflags '-static'" -tags netgo -installsuffix netgo cmd/sensoroni.go
