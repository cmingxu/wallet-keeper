Jex backend Golang implementation

## Run
```
	make
	./bin/wallet-keeper --http-listen-addr=http://0.0.0.0:8080 \
   --log-level debug --log-path=/tmp/wallet-keeper.log
```

## How to config

```
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

```


### getblockcount
```
$ curl -sSL -H "CoinType:btc" localhost:8000/getblockcount |  jq .
{
  "message": "603443"
}
```

### getaddress
```
$ curl -sSL -H "CoinType:btc" localhost:8000/getaddresses?accounts=barfoox |  jq .
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
```
$ curl -sSL -H "CoinType:btc" localhost:8000/getnewaddress?accounts=barfoox |  jq .
{
  "message": "2MwTYicqDegSuR2MsTeRVFLwkhMYSZKHPiP"
}
```

### listaccounts
```
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
```
$ curl -sSL -H "CoinType:btc" 'localhost:8000/sendtoaddress?address=2N2VJhke2sWspswJKWTFjqfibRY1wfZPbEQ&amount=0.1' |  jq .
{
  "message": "-6: Insufficient funds"
}
```
