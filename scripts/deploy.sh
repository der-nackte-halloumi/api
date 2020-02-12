#!/usr/bin/env bash

set -eo pipefail

echo ">>> pushing latest images for api and migrations"
./scripts/docker_push.sh latest

chmod 400 der-nackte-halloumi-api-travis

echo $SERVER_PUBLIC_KEY >> $HOME/.ssh/known_hosts

# run migrations
echo ">>> run migrations with image dernacktehalloumi/migrations:$TRAVIS_BUILD_NUMBER"
ssh -i der-nackte-halloumi-api-travis $SERVER_ADDRESS "MIGRATIONS_IMAGE_TAG=$TRAVIS_BUILD_NUMBER docker-compose up --build -d migrations && docker wait migrations"

# restart service
echo ">>> restarting api with image dernacktehalloumi/api:$TRAVIS_BUILD_NUMBER"
ssh -i der-nackte-halloumi-api-travis $SERVER_ADDRESS "API_IMAGE_TAG=$TRAVIS_BUILD_NUMBER docker-compose up --build -d api"
