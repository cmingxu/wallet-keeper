#!/usr/bin/env sh

CONTAINER_NAME=bitcoind-node
DAEMON=${1:yes}

if [[ x${DAEMON} == x"yes" ]]; then
  DOCKER_ARGS=-d
else
  DOCKER_ARGS=
fi

USER=
PASS=
BITCOIND_ARGS="--testnet"
if [[ "x$USER" != "x" ]]; then
  BITCOIND_ARGS="$BITCOIND_ARGS --rpcuser=$USER"
fi

if [[ "x$PASS" != "x" ]]; then
  BITCOIND_ARGS="$BITCOIND_ARGS --rpcpassword=$PASS"
fi

docker rm -f ${CONTAINER_NAME} 2>&1 >/dev/null

docker run -v $(pwd)/bitcoind-data:/bitcoin \
  -v $(pwd)/bitcoin.conf:/bitcoin/.bitcoin/bitcoin.conf \
  --name=${CONTAINER_NAME} ${DOCKER_ARGS} \
  -p 8333:8333 \
  -p 127.0.0.1:8332:8332 \
  kylemanna/bitcoind
