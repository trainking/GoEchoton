# GoEchoton

- [GoEchoton](#goechoton)
	- [1. 概述](#1-概述)
	- [2. 目录结构](#2-目录结构)
	- [3. 关键实现](#3-关键实现)
		- [3.1 灵活的配置路由](#31-灵活的配置路由)
		- [3.2 Arpc的连接池](#32-arpc的连接池)
		- [3.3 参数验证器](#33-参数验证器)
		- [3.4 自定义工具包](#34-自定义工具包)
	- [4. 说明](#4-说明)

## 1. 概述

该项目基于`golang Echo`框架，搭建日常开发使用项目骨架，提高开发效率，统一开发语言

## 2. 目录结构

* `/bin` 主要存放编译后文件
* `/pkg` 开发一些通用功能包
* `/config` 配置处理
* `/bootstrap/bootstrap.go` 项目启动文件，其中包括初始化操作，如环境配置，启动服务，注册路由等等
* `database` 数据库连接创建的文件
* `handler` 处理`http`请求和返回，相当于**MVC**的**Controller**
* `log` 记录日志
* `middleware` 存放自定义请求中间件
* `model` 模型定义文件
* `repository` 存放模型数据操作逻辑文件，通过interface方式，方便注入不同的数据库
* `service` 业务处理逻辑代码，一般`handler`调用
* `router` 路由定义文件
* `utils` 自定义工具
* `test` 单元测试代码
* `main.go` server的入口文件

## 3. 关键实现

### 3.1 灵活的配置路由

`Echo` 原生的路由配置方式，是调用`Echo#Add()`方法，增加路由。组路由，则是增加组对象之后，再增加路由。

这样的好处是开发简单，但是坏处是结构混乱，容易写在一坨代码里面，难以有结构层次感。

在`./router/types.go` 中，我抽象了两个重要的结构体`Router`和`Group`，将`Echo`隐藏属性，开放出来，在包中定义两个包变量`Routers`，`Groups`。通过`init()`函数的包加载初始化的特性，给这两个包变量复制，然后再`Add`到`Echo`的路由之中。

定义路由，便只需要在`init()`函数中加入即可:

```golang
func init() {
	Routers = append(Routers, []Router{
		{
			Method:  http.MethodGet,
			Path:    "/",
			Handler: handler.User.Index,
		},
		{
			Method:  http.MethodPost,
			Path:    "/login",
			Handler: handler.User.Login,
		},
	}...)
}
```

### 3.2 Arpc的连接池

在`./service/arpc.go`中，实现了一个连接池控制，`arpc`推荐长连接的`tcp`方式实现，所以，使用连接池可以更高服用。

需要注意的是，每次获取的链接，通过手动释放，放回池中：

```

client, err := this.GetClient()

...

defer this.Release(client)

```

### 3.3 参数验证器

`./bootstrap/struct_validateor.go`中，实现了一个自定义的结构体验证其，通过`json`参数与结构体映射的定义语句，实现验证，要验证时，需要调用方法:

```
c.Validate(&param)
```

### 3.4 自定义工具包

- ~~`pkg/aprcpool` arpc的连接池~~
- `pkg/cache` 内存缓存
- `pkg/etcd` Etcd http客户端
- `pkg/flatbuf` flatbuffer解析
- `pkg/sqlt` sql生成模板
- `pkg/upload` 文件上传示例
- `pkg/wechatpayx` 微信支付实现


## 4. 说明

持续更新会在`web-layout`分支中，有些东西，可能随着开发递进，可能再做优化
