> 该项目仍处于开发阶段。。。
>
> 项目目录结构参考了 [go-zero](https://github.com/zeromicro/go-zero) 与 [kratos](https://github.com/go-kratos/kratos)
> 组织。
>
> 部分代码参考了 [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin) 的实现。

## 涉及技术

- 服务注册发现：etcd
- rpc 及服务负载均衡：grpc、protobuf
- 数据操作及日志：gorm、go-redis、zap、casbin
- 网关 - RESTful API：gin
- ......

## 运行项目

```shell
# 安装依赖
go mod tidy

# 如果使用 Goland 编辑器，直接执行 app/user/cmd/server 下的 main 文件即可

# 使用命令行执行
cd app/user/cmd/server

go run *.go -f ../../config/config.yaml
```

## 模块列表

- [x] 用户管理
- [x] 角色-权限管理
- [x] 路由管理
- [x] 菜单管理
- [ ] ......
