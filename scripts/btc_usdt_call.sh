#!/usr/bin/env sh

docker run --net host \ 
  -it ruimarinho/bitcoin-core bitcoin-cli \
  --testnet \
  --rpcconnect=<ask me> \
  --rpcuser=omnicore \
  --rpcpassword=7zmZ*gV6sQK \
  --rpcport=18332 "$@"
