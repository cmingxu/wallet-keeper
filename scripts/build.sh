#!/usr/bin/env bash

ROOT_PATH=$(cd "$(dirname $BASH_SOURCE[0])/.." && pwd)

VERSION=$(cat ./VERSION)
RELEASE_IMAGE=wallet_keeper:${VERSION}


docker build -t $RELEASE_IMAGE ${ENV} --no-cache --rm -f ./Dockerfile .
