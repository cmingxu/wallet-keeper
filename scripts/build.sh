#!/usr/bin/env bash

ROOT_PATH=$(cd "$(dirname $BASH_SOURCE[0])/.." && pwd)

GOPROXY=${GOPROXY:-https://goproxy.io}

BUILD_CONTAINER_NAME=wallet_keeper_build
BUILD_IMAGE=wk_build_base

cd $ROOT_PATH
VERSION=$(cat ./VERSION)
RELEASE_IMAGE=wallet_keeper:${VERSION}

if [[ -z $(docker images | grep "$BUILD_IMAGE") ]]; then
  docker build -t $BUILD_IMAGE --no-cache --rm -f ./Dockerfile.build .
fi

# check if golang.org can be reached
ping -q -W 1 -c 1 golang.org
if [ $? == "0" ]; then
  ENV=''
else
  ENV="--env GOPROXY=${GOPROXY}"
fi


echo ${ENV}
docker run --rm \
  --name $BUILD_CONTAINER_NAME \
  -v "$(pwd)":/go/src/app \
  ${ENV} \
  $BUILD_IMAGE make

binary=./bin/wallet-keeper-${VERSION}
if [[ -f ${binary} ]]; then
  docker build -t $RELEASE_IMAGE --build-arg BINARY=$(basename ${binary}) --no-cache --rm -f ./Dockerfile .
else
  echo "unable to locate binary file '${binary}', stop"
fi

docker build -t $RELEASE_IMAGE ${ENV} --no-cache --rm -f ./Dockerfile .
