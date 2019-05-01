FROM golang:1.12.1-alpine3.9 as builder

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

WORKDIR /go/src/app

RUN apk add git && apk add make && apk add gcc && apk add libc-dev  \
  && apk add --update gcc musl-dev

ENV GOPROXY=https://goproxy.io
ADD . .

RUN make



FROM alpine:latest

COPY --from=builder /go/src/app/bin/wallet-keeper /

EXPOSE 8000
WORKDIR /

CMD ["/wallet-keeper", "run"]
