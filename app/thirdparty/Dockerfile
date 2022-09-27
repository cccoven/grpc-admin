FROM golang:1.18.3-alpine as builder 
WORKDIR /go/src/github.com/grpc-admin
COPY . . 
RUN go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o thirdparty app/thirdparty/cmd/server/*.go \
    && mv thirdparty app/thirdparty/cmd/server

FROM alpine:latest as runner
LABEL MAINTAINER="example@gmail.com"
WORKDIR /app
COPY --from=0 /go/src/github.com/grpc-admin/app/thirdparty/cmd/server/thirdparty ./
COPY --from=0 /go/src/github.com/grpc-admin/app/thirdparty/config/config.yaml ./
EXPOSE 8001
ENTRYPOINT ["./thirdparty", "-f", "config.yaml"]