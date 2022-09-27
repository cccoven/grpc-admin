package main

import (
	"flag"
	"grpc-admin/app/gateway/bootstrap"
	"grpc-admin/app/gateway/conf"
	"grpc-admin/common/pkg"
)

var configFile = flag.String("f", "app/gateway/config.yaml", "The conf file")

func main() {
	flag.Parse()
	pkg.LoadConfig(*configFile, &conf.AppConf)

	bootstrap.Run()
}
