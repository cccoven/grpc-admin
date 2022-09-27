FROM golang:1.18.3-alpine as builder
WORKDIR /go/src/github.com/grpc-admin
COPY . .
RUN go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o system app/system/cmd/server/*.go \
    && mv system app/system/cmd/server 

FROM alpine:latest as runner
LABEL MAINTAINER="example@gmail.com"
WORKDIR /app
COPY --from=0 /go/src/github.com/grpc-admin/app/system/cmd/server/system ./
COPY --from=0 /go/src/github.com/grpc-admin/app/system/config/config.yaml ./
EXPOSE 8003
ENTRYPOINT ["./system", "-f", "config.yaml"]