#!/usr/bin/env sh

CONTAINER_NAME=bitcoind-node
#IMAGE_NAME=cmingxu/bitcoin-core
IMAGE_NAME=ruimarinho/bitcoin-core
DAEMON=${1:yes}

if [[ x${DAEMON} == x"yes" ]]; then
  DOCKER_ARGS=-d
else
  DOCKER_ARGS=
fi

BITCOIN_DATA=/bitcoind-data


docker rm -f ${CONTAINER_NAME} 2>&1 >/dev/null

docker run --env=BITCOIN_DATA=${BITCOIN_DATA} \
  -v $(pwd)/bitcoind-data:/bitcoind-data \
  -v $(pwd)/bitcoin.conf:/bitcoin.conf \
  --name=${CONTAINER_NAME} \
  ${DOCKER_ARGS} \
  -p 8333:8333 \
  -p 0.0.0.0:8332:8332 \
  -p 0.0.0.0:18332:18332 \
  ${IMAGE_NAME} -conf=/bitcoin.conf -deprecatedrpc=accounts
