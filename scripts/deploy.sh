#!/usr/bin/env bash

set -eo pipefail

./scripts/docker_push.sh $1

chmod 400 der-nackte-halloumi-api-travis

echo $SERVER_PUBLIC_KEY >> $HOME/.ssh/known_hosts

# restart service
ssh -i der-nackte-halloumi-api-travis $SERVER_ADDRESS "docker-compose up --build -d api"
