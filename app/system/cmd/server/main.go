package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc-admin/app/system/internal/conf"
	"grpc-admin/app/system/internal/server"
	"grpc-admin/app/system/system"
	"grpc-admin/common/pkg"
	"log"
	"net"
)

var configFile = flag.String("f", "app/system/config/config.yaml", "The conf file")

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

	system.RegisterSystemServer(s, server.NewSystemServer())

	// 注册到 Etcd
	registerService()

	log.Printf("Server [%s] listning at %s...\n", conf.AppConf.Service.Name, conf.AppConf.Service.Host)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
