Wallet-Keeper simplified access to BTC/OmniProtocol/ETH backends with unified APIs, help cryptocurrency developers focus their bussiness logic instead of RPC details of different backend.

## Run
```
	make
	./bin/wallet-keeper --log-level debug --log-dir=/tmp/  \
  run --http-listen-addr=http://0.0.0.0:8080
```

## Run In Docker(need connectivity golang.org/x/XXXX in docker build)
```bash
./scripts/build.sh
./scripts/run_wallet_keeper_docker.sh
```


## How to config

```bash
$ ./bin/wallet-keeper-0.0.1
NAME:
   wallet-keeper-0.0.1 - A new cli application

USAGE:
   wallet-keeper-0.0.1 [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     run, r   serve api gateway
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value  default log level (default: "info") [$LOG_LEVEL]
   --log-path value    [$LOG_PATH]
   --env value        (default: "production") [$ENV]
   --help, -h         show help
   --version, -v      print the version


$> ./bin/wallet-keeper-0.1.0 run --help
use stdout as default log output
NAME:
   wallet-keeper-0.1.0 run - serve api gateway

USAGE:
   wallet-keeper-0.1.0 run [command options] [arguments...]

OPTIONS:
   --http-listen-addr value  http address of web application (default: "0.0.0.0:8000") [$HTTP_LISTEN_ADDR]
   --btc-rpc-addr value      [NOTICE] testnet and mainnet have different default port (default: "192.168.0.101:8332") [$BTC_RPCADDR]
   --btc-rpc-user value      (default: "foo") [$BTC_RPCUSER]
   --btc-rpc-pass value      password can be generate through scripts/rcpauth.py (default: "qDDZdeQ5vw9XXFeVnXT4PZ--tGN2xNjjR4nrtyszZx0=") [$BTC_PRCPASS]
   --usdt-rpc-addr value     [NOTICE] testnet and mainnet have different default port (default: "localhost:18332") [$USDT_RPCADDR]
   --usdt-rpc-user value     (default: "foo") [$USDT_RPCUSER]
   --usdt-rpc-pass value     password can be generate through scripts/rcpauth.py (default: "usdtpass") [$USDT_PRCPASS]
   --usdt-property-id value  property id of usdt, default is 2 (default: 2) [$USDT_PROPERTY_ID]
```



## BTC - bitcored

### getblockcount

```bash
$ curl -sSL -H "CoinType:btc" localhost:8000/getblockcount |  jq .
{
  "message": "603443"
}
```

### getaddress
```bash
$ curl -sSL -H "CoinType:btc" localhost:8000/getaddress?account=barfoox |  jq .
{
  "message": [
    "2MwFt5ZbGfK2yqCWHb1hyGKkm8K6DUByPj8",
    "2N3Mqnjq9KUnLqUNjRdgkqh1VY4DJBjPoie",
    "2N4MbjrKuBD9KYFitz6nmSZFBJxdSWguC7Y",
    "2NDuD1sWwsuZeVdBCv8pusjSQUCNTbJTR7x"
  ]
}
```

### getnewaddress
```bash
$ curl -sSL -H "CoinType:btc" localhost:8000/getnewaddress?accounts=barfoox |  jq .
{
  "message": "2MwTYicqDegSuR2MsTeRVFLwkhMYSZKHPiP"
}
```

### listaccounts
```bash
$ curl -sSL -H "CoinType:btc" localhost:8000/listaccounts |  jq .
{
  "message": {
    "barfoo": 0,
    "barfoox": 0,
    "duckduck": 0,
    "foobar": 0
  }
}
```

### sendtoaddress
```bash
$ curl -sSL -H "CoinType:btc" 'localhost:8000/sendtoaddress?address=2N2VJhke2sWspswJKWTFjqfibRY1wfZPbEQ&amount=0.1' |  jq .
{
  "message": "-6: Insufficient funds"
}
```

### getaddressesbyaccount
```bash
$ curl -sSL -H "CoinType:btc" 'localhost:8000/getaddressesbyaccount?account=foobar' |  jq .
{
  "message": [
    "2MziFYWKdptgkDn9esKhQnQXt86H6taGM3f",
    "2MziSR87om6fZyknsTMF447Yftt5afQB9GN",
    "2NBEYgwGWTiTaPvCKp64hDeP5xnwckdrYNK",
    "2NGTBdRezf7TNtF3X1ptmWzVr4XWc8MHnnP"
  ]
}
```

### getaccountinfo
```bash
$ curl -sSL -H "CoinType:btc" 'localhost:8000/getaccountinfo?account=barfoo' |  jq .
{
  "message": {
    "account": "barfoo",
    "balance": 0.0015,
    "addresses": [
      "2Mtb1opq1JvzfdLdGRPFSWwbEmvRGBXYdos",
      "2MxK75vSZLABDgqZRmUPMt6kfyabXZD81SJ",
      "2N5ETGDkhFFZqKQWhcbfYMihv9rSorMQLP9",
      "2ND1nUpT3in3HeMjssd7XkPa3j6nHXQgs1G"
    ]
  }
}
```

### createaccount
```bash
$ curl -sSL -H "CoinType:btc" 'localhost:8000/createaccount?account=barfoo1' |  jq .
{
  "message": {
    "account": "barfoo1",
    "balance": 0,
    "addresses": [
      "2N7y28aMq3xfxEZBFDvfvVzQ6aZLXr7np64"
    ]
  }
}
```

## USDT - omnicore

### getblockcount

```bash
$ curl -sSL -H "CoinType:usdt" localhost:8000/getblockcount |  jq .
{
  "message": "1488490"
}
```



### getaddress

```bash
$ curl -sSL -H "CoinType:usdt" localhost:8000/getaddress??account=test |  jq .
{
  "message": "mqyY2dscucMNBHAiUid88mEgkvZ3ADJLeZ"
}
```



### getaddressesbyaccount

```bash
$ curl -sSL -H "CoinType:usdt" 'localhost:8000/getaddressesbyaccount?account=test' |  jq .
{
  "message": [
    "mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4",
    "mvePubvtfpAo18ykLeZ1hD6BMYDismUqxf"
  ]
}
```



### getaccountinfo

```bash
$ curl -sSL -H "CoinType:usdt" 'localhost:8000/getaccountinfo?account=test' |  jq .
{
  "message": {
    "account": "test",
    "balance": 10,
    "addresses": [
      "mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4",
      "mvePubvtfpAo18ykLeZ1hD6BMYDismUqxf"
    ]
  }
}
```



### move

```bash
$ curl -sSL -H "CoinType:usdt" 'localhost:8000/move?from=mvePubvtfpAo18ykLeZ1hD6BMYDismUqxf&to=mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4&amount=3' |  jq .
{
  "message": "success"
}
```



### sendfrom

```bash
$ curl -sSL -H "CoinType:usdt" 'localhost:8000/sendfrom?from=test&address=mkVW5QARPvvAY328y9LMHep5ZhuiTrgdd4&amount=3' |  jq .
{
  "message": "success"
}
```





