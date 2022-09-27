FROM golang:1.18.3-alpine as builder 
WORKDIR /go/src/github.com/grpc-admin
# 由于单独的服务里依赖了 grpc-admin/common 包，所以需要将 grpc-admin 下的所有文件复制到容器中
# 执行 docker build 命令时，需要带上 -f <Docker路径>
COPY . . 
RUN go env -w CGO_ENABLED=0 \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o user app/user/cmd/server/*.go \
    && mv user app/user/cmd/server

FROM alpine:latest as runner
LABEL MAINTAINER="example@gmail.com"
WORKDIR /app
COPY --from=0 /go/src/github.com/grpc-admin/app/user/cmd/server/user ./
COPY --from=0 /go/src/github.com/grpc-admin/app/user/config/config.yaml ./
EXPOSE 8001
ENTRYPOINT ["./user", "-f", "config.yaml"]