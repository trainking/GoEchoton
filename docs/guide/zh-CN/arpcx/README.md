# arpcx

- [arpcx](#arpcx)
  - [概述](#概述)
  - [快速开始](#快速开始)
    - [服务端](#服务端)
    - [rpc定义](#rpc定义)
    - [客户端](#客户端)
  - [注意事项](#注意事项)

## 概述

基于`arpc`包装出来的rpc包，实现了以下功能：

- rpc的`server`和`client`的包装事件
- 基于Etcd的服务发现和服务注册

> arpcx 依赖etcdx的实现

## 快速开始

一个实例的目录结构:

```
├─skeleton      // 服务端核心代码
│  ├─handler    // 服务入口Handler
│  └─service    // 业务逻辑的执行service, 一般实现userrpc包中的接口
├─svc           // ServerContext的实现 
├─userstub      // 服务客户端实现，实现了userrpc包中的接口
├─userrpc       // 定义客户端和服务端交互的接口和参数及返回结构体
└─userrpc.go    // 服务启动main入口
```

### 服务端

```go
var listenAddr = flag.String("addr", ":8001", "listen address")
var etcdGateway = flag.String("etcd", "127.0.0.1:2379", "etcdGateway")

func main() {
	flag.Parse()
	svcCtx := svc.New()
	server := arpcx.NewServer(svcCtx)

	// 注册服务
	server.RegisterToEtcd(userrpc.UserRpcTarget, *listenAddr, strings.Split(*etcdGateway, ","))

	server.Start(*listenAddr)
}
```

`handler`的路由定义在svc.SvcContext中:

```
// GetHandlers 定义的Handler
func (s *SvcContext) GetHandlers() []arpcx.Handler {
	return []arpcx.Handler{
		{
			Path:   userrpc.UserCheckPasswdPath,
			Handle: handler.UserCheckPasswd,
		},
	}
}
```

### rpc定义

需要定义个`XXXRpc`结构的接口，定义传入调用函数，和传入返回结构体：

```go
type (

	// UserRpc 定义接口
	UserRpc interface {
		//CheckPasswd 检查密码
		CheckPasswd(ctx context.Context, p *CheckPasswd) error
	}

	CheckPasswd struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}
)
```

同时需要定义两个常量:

```go
const UserRpcTarget = "/user.rpc/"  // 服务名，用于Etcd做前缀

const UserCheckPasswdPath = "/user/checkpasswd"  // 调用接口路径，接口标识
```

### 客户端

客户端实现:

```go
type userRpc struct {
	client *arpcx.ClientEtcdPod
}

func NewUserRpc(etcdGateway []string) userrpc.UserRpc {
	client, err := arpcx.NewClientPool(userrpc.UserRpcTarget, etcdGateway)
	if err != nil {
		panic(err)
	}
	return &userRpc{
		client: client,
	}
}

// CheckPasswd 检查密码
func (u *userRpc) CheckPasswd(ctx context.Context, p *userrpc.CheckPasswd) error {
	return u.client.GetNode().CallWith(ctx, userrpc.UserCheckPasswdPath, p, &struct{}{})
}
```

## 注意事项

- 服务端包命名为`skeleton`
- 客户端包命名一般为`XXXstub`
- 依赖`etcdx`包用于服务发现
- 当无需返回和无需参数的`Call`函数中，建议传`struct{}{}`
