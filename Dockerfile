FROM golang:alpine AS builder

WORKDIR /go/src/check-in
COPY . .

RUN go generate && go env && go build -o client .

FROM alpine:latest
LABEL MAINTAINER="386139859@qq.com"

WORKDIR /go/src/wailian-client

COPY --from=builder /go/src/client ./

EXPOSE 8080

ENTRYPOINT ./client -c config.yaml
