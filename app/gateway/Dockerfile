FROM golang:1.18.3-alpine as builder
WORKDIR /go/src/github.com/grpc-admin
COPY . .
RUN go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o gateway app/gateway/*.go \
    && mv gateway app/gateway

FROM alpine:latest as runner
LABEL MAINTAINER="example@gmail.com"
WORKDIR /app
COPY --from=0 /go/src/github.com/grpc-admin/app/gateway/gateway ./
COPY --from=0 /go/src/github.com/grpc-admin/app/gateway/config.yaml ./
EXPOSE 8000
ENTRYPOINT ["./gateway", "-f", "config.yaml"]