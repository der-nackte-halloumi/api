#!/bin/bash

set -e

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

function image_with_tag_exists () {
  if test ! -z "$(docker images -q $1)"; then
    return 0
  fi
  return 1
}

image_with_tag_exists "dernacktehalloumi/api:latest"
if [ "$?" -eq 0 ]; then
  echo "--> pushing dernacktehalloumi/api:latest"
  docker push dernacktehalloumi/api:latest
fi

image_with_tag_exists "dernacktehalloumi/api:$TRAVIS_BUILD_NUMBER"
if [ "$?" -eq 0 ]; then
  echo "--> pushing dernacktehalloumi/api:$TRAVIS_BUILD_NUMBER"
  docker push dernacktehalloumi/api:$TRAVIS_BUILD_NUMBER
fi
