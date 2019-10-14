FROM golang:1.13-alpine3.10 as builder
LABEL maintainer="hetal<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/qrcode/

COPY Makefile .
COPY go.mod .
COPY src src

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add git tree \
	&& tree -L 3 && export GOPATH=/data/qrcode/ && make build \
	&& rm -rf /var/lib/apk/*

FROM alpine:3.9 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /data/qrcode/

COPY --from=0 /data/qrcode/qrcode .

HEALTHCHECK --interval=5m --timeout=3s CMD curl -f http://localhost/ || exit 1

EXPOSE 80/tcp

CMD ["/data/qrcode/qrcode"]
