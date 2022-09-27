package main

import (
	"flag"
	"google.golang.org/grpc"
	"grpc-admin/app/thirdparty/internal/conf"
	"grpc-admin/app/thirdparty/internal/server"
	"grpc-admin/app/thirdparty/thirdparty"
	"grpc-admin/common/pkg"
	"log"
	"net"
)

var configFile = flag.String("f", "../../config/config.yaml", "The conf file")

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

	thirdparty.RegisterThirdPartyServer(s, server.NewThirdPartyServer())
	
	// 注册到 Etcd
	registerService()

	log.Printf("Server [%s] listning at %s:%s...\n", conf.AppConf.Service.Name, conf.AppConf.Service.Host, conf.AppConf.Service.Port)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
