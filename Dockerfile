FROM alpine:latest

ARG BINARY
ADD ./bin/${BINARY} /wallet-keeper


EXPOSE 8000
WORKDIR /

CMD ["/wallet-keeper", "run"]
