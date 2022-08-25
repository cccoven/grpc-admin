package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"grpc-admin/app/thirdparty/internal/conf"
	"log"
	"time"
)

var (
	cli    *clientv3.Client
	prefix = "grpc_app"
)

func keepAlive(lease *clientv3.LeaseGrantResponse) {
	ctx := context.TODO()
	// 调用 KeepAlive 保持服务心跳
	ch, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Fatal(err)
	}
	// 在测试中，KeepAlive 会出现失败的情况（queue is null），所以还是需要异步地检测 keepalive 的情况
	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					// 续约失败，让这个 key 立刻过期并重新注册服务
					cli.Revoke(ctx, lease.ID)
					register()
					return
				}
			}
		}
	}()
}

func register() {
	ctx := context.TODO()
	ttl := 5
	lease, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		log.Fatal(err)
	}
	// 可以使用 target 作为一个应用的服务目录名，用于命名隔离
	em, err := endpoints.NewManager(cli, prefix)
	if err != nil {
		log.Fatal(err)
	}

	// 服务的 key 必须带上一个唯一性的 ID，否则多个服务注册时会造成相同的 key 覆盖，从而导致客户端永远只能发现一个服务
	fullKey := fmt.Sprintf("%s/%s/%d", prefix, conf.AppConf.Service.Name, lease.ID)
	err = em.AddEndpoint(ctx, fullKey, endpoints.Endpoint{Addr: conf.AppConf.Service.Host}, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Fatal(err)
	}

	keepAlive(lease)
}

func registerService() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   conf.AppConf.Etcd.Hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	register()
}
