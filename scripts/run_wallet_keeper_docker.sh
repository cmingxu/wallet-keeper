#!/usr/bin/env bash

ROOT_PATH=$(cd "$(dirname $BASH_SOURCE[0])/.." && pwd)

cd $ROOT_PATH
VERSION=$(cat ./VERSION)
RELEASE_IMAGE=wallet_keeper:${VERSION}
CONTAINER_NAME=wallet_keeper_${VERSION}


docker run -it  \
  --publish 127.0.0.1:8000:8000/tcp \
  --volume /data/wallet-keeper/eth-accounts.json:/data/eth-accounts.json \
  --name $CONTAINER_NAME $RELEASE_IMAGE

