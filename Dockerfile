FROM golang:1.17-alpine as builder
LABEL maintainer="hetal<hetao@hetao.name>"
LABEL version="1.0"

WORKDIR /data/qrcode/

COPY . .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \ 
	&& apk update && apk add tree \
	&& tree -L 5 && export GOPROXY=https://goproxy.cn && go build -o bin/qrcode \
	&& rm -rf /var/lib/apk/*

FROM alpine:3.14 as prod

RUN apk --no-cache add ca-certificates

WORKDIR /data/qrcode/

RUN mkdir bin
COPY --from=0 /data/qrcode/bin/qrcode bin/

HEALTHCHECK --interval=5s --timeout=5s --retries=3 \
    CMD ps aux | grep "qrcode" | grep -v "grep" > /dev/null; if [ 0 != $? ]; then exit 1; fi

CMD ["/data/qrcode/bin/qrcode", "-b" , "0.0.0.0:80"]
