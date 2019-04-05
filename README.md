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
