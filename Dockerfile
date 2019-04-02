FROM alpine3.9

ADD ./bin/wallet-keeper /

WORKDIR /

CMD wallet-keeper

