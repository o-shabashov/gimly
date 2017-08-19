#!/usr/bin/env bash
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/gimly_static .
env GOOS=linux GOARCH=amd64 go build -o bin/gimly_dynamic