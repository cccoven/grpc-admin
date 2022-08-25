package rpc

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-admin/app/gateway/conf"
	"time"
)

var (
	etcdCli *clientv3.Client
	prefix  = "grpc_app"
)

func Discovery(serviceName string, balanceMode string) *grpc.ClientConn {
	var err error

	if etcdCli == nil {
		etcdCli, err = clientv3.New(clientv3.Config{
			Endpoints:   conf.AppConf.Etcd.Hosts,
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			return nil
		}
	}

	// etcd 官方实现的 builder 对象
	etcdResolver, err := resolver.NewBuilder(etcdCli)
	if err != nil {
		return nil
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("etcd:///%s/%s", prefix, serviceName),
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"loadBalancingConfig":[{"%s":{}}]}`, balanceMode)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil
	}
	// TODO conn.Close()

	return conn
}
