#!/usr/bin/env bash

# Set environment variables from .env file
if [ -f .env ]
  then
    export $(cat .env | sed 's/#.*//g' | xargs)
fi

go run ./cmd/migrate/main.go up
