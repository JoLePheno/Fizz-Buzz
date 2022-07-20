#!/bin/bash

echo "Building server"

CGO_ENABLED=0 GOOS=linux go build -o fizz_buzz .

echo "Running server"

./fizz_buzz