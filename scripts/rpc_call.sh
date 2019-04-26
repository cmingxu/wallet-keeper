#!/usr/bin/env sh

# docker run --net host -it ruimarinho/bitcoin-core bitcoin-cli --testnet  --rpcuser=foo --rpcpassword=qDDZdeQ5vw9XXFeVnXT4PZ--tGN2xNjjR4nrtyszZx0= --rpcport=8332 "$@"
docker run --net host -it ruimarinho/bitcoin-core bitcoin-cli --testnet  --rpcuser=xcm --rpcpassword=rpcpassword --rpcport=18332 "$@"
