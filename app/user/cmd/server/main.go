package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc-admin/app/user/internal/conf"
	"grpc-admin/app/user/internal/server"
	"grpc-admin/app/user/user"
	"grpc-admin/common/pkg"
	"log"
	"net"
)

// 注意配置文件路径
// 开发时是基于当前工作目录查找
// 编译发布时需要使用相对路径找到 config 目录下的配置文件：../../config/config.yaml
var configFile = flag.String("f", "app/user/config/config.yaml", "The config file")

func main() {
	flag.Parse()
	pkg.LoadConfig(*configFile, &conf.AppConf)

	lis, err := net.Listen("tcp", conf.AppConf.Service.Host)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	defer s.GracefulStop()

	user.RegisterUserServer(s, server.NewUserServer())

	// 注册到 Etcd
	registerService()

	log.Printf("Server [%s] listning at %s...\n", conf.AppConf.Service.Name, conf.AppConf.Service.Host)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
