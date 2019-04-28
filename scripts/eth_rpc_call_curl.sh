#!/usr/bin/env sh

METHOD=$1
PARAM=$2

curl -X POST -H 'Content-Type: application/json' \
--data "{\"jsonrpc\":\"2.0\",\"method\":\"${METHOD}\",\"params\":[${PARAM}],\"id\":1}" \
http://154.8.201.160:8545/
