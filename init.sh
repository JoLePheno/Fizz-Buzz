#!/bin/bash

echo "Running database migration"

fizz_buzzctl init
fizz_buzzctl version
fizz_buzzctl up

echo "Running server"

fizz_buzzd
