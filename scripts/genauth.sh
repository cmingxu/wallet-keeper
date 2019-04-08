#!/usr/bin/env sh

if [[ $# != 1 ]];then
  echo "./genauth.sh <username>"
else
  curl -sSL https://raw.githubusercontent.com/bitcoin/bitcoin/master/share/rpcauth/rpcauth.py | python - $1
fi

