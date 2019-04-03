FROM alpine:latest

ARG BINARY
ADD ./bin/${BINARY} /

WORKDIR /

CMD ${BINARY}

