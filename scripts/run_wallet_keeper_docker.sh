#!/usr/bin/env bash

ROOT_PATH=$(cd "$(dirname $BASH_SOURCE[0])/.." && pwd)

cd $ROOT_PATH
VERSION=$(cat ./VERSION)
RELEASE_IMAGE=wallet_keeper:${VERSION}
CONTAINER_NAME=wallet_keeper_${VERSION}

docker run -it --detach --publish localhost:8000:8000 --name $CONTAINER_NAME $RELEASE_IMAGE

