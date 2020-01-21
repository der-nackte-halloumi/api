#!/usr/bin/env bash

set -eo pipefail

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

function image_exists () {
  if test ! -z "$(docker images -q $1)"; then
    echo "image $1 does exist"
    return 0
  fi
  
  echo "image $1 does not exist"
  return 1
}

image_exists "dernacktehalloumi/api:$1"
if [ "$?" -eq 0 ]; then
  echo "--> pushing dernacktehalloumi/api:$1"
  docker push dernacktehalloumi/api:$1
fi

