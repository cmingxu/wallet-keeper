#!/usr/bin/env sh

docker run --net host \
  -it ethereum/client-go \
  --testnet \
  --rpcconnect=154.8.201.160 \
  -rpcport=18332 "$@"
