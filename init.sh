#!/bin/bash

echo "Running migrations"

./cmd/fizzbuzzctl/migrations_up.sh

echo "Building server"

CGO_ENABLED=0 go build -o fizz_buzz ./cmd/fizzbuzzd/main.go

echo "Running server"

./fizz_buzz