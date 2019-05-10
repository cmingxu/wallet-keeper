#!/usr/bin/env bash

./bin/wallet-keeper \
  --log-level debug \
  --log-dir /tmp \
  --env development \
  run \
  --http-listen-addr 0.0.0.0:8080 \
  --backends btc,usdt,eth \
  --rpc-addr ws://107.150.126.20:9546 \
  --eth-wallet-dir /tmp/wallets
  --eth-rpc-addr  http://192.168.0.101:8545 \
  --eth-account-path /tmp/eth-accounts.json \
  --btc-rpc-addr 192.168.0.101:8332 \
  --btc-rpc-user foo \ 
  --btc-rpc-pass bar \
  --usdt-rpc-addr localhost:18332 \
  --usdt-rpc-user foo \
  --usdt-rpc-pass bar \
  --usdt-property-id 31




