# apiserver

- [apiserver](#apiserver)
  - [概述](#概述)
  - [快速开始](#快速开始)
    - [svc](#svc)
    - [配置](#配置)
    - [路由](#路由)
    - [中间件](#中间件)
    - [验证器](#验证器)
  - [待实现](#待实现)

## 概述

基于`echo`抽象出来的Api服务端实践，依赖以下包:

```go
go get github.com/labstack/echo/v4
go get github.com/go-playground/validator/v10
```

## 快速开始

新开始一个服务:

```go
func main() {
	flag.Parse()

	var conf config.Config
	if err := apiserver.LoadConfigFile("./configs/authserver.yaml", &conf); err != nil {
		panic(err)
	}
	svcCtx := svc.New(&conf)
	server := apiserver.New(svcCtx)
	server.Start(*listenAddr)
}
```

### svc

`svc`定义是启动环境必须的内容，通过依赖注入的方式。在`apiserver`已经定义了其接口`apiserver.ServerContext`。**每一个应用必须自己实现这个接口**，约定使用包名`svc`。

```go
type SvcContext struct {
	conf *config.Config
}

func New(conf *config.Config) apiserver.ServerContext {
	return &SvcContext{
		conf: conf,
	}
}

//GetRouters 获取顶级路由
func (s *SvcContext) GetRouters() []apiserver.Router {
	loginApi := login.New()
	return []apiserver.Router{
		{
			Method:  http.MethodPost,
			Path:    "/v1/login/one",
			Name:    "登录第一步",
			Handler: loginApi.LoginOne,
		},
	}
}

// GetGroups 获取组路由
func (s *SvcContext) GetGroups() []apiserver.Group {
	return []apiserver.Group{}
}

func (s *SvcContext) GetMiddlewares() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{}
}

// GetValidators 获取自定义验证器
func (s *SvcContext) GetValidators() map[string]validator.Func {
	return map[string]validator.Func{}
}
```

### 配置

配置使用`yaml`格式定义，每个应用需要自定一个`Config`的结构体，解析配置结构体，`apiservr`提供`LoadConfigFile`函数加载文件中配置。

### 路由

路由配置分为两种，定义在`svc.GetRouters`中的，为顶级路由，即无组路由：

```go
func (s *SvcContext) GetRouters() []apiserver.Router {
	loginApi := login.New()
	return []apiserver.Router{
		{
			Method:  http.MethodPost,    // 请求Method
			Path:    "/v1/login/one",   // 路径，全路径
			Name:    "登录第一步",      // 路由名称
			Handler: loginApi.LoginOne,   // 调用handler， echo.HandlerFunc的实现
            // Middlwares []echo.MiddlewareFunc{}   // 该路由应用的中间件, 仅在该路由上生效
		},
	}
}
```

组路由，定义在`svc.GetGroups`，将路由分组，该组内路由共用路径前缀和中间件:

```go
// GetGroups 获取组路由
func (s *SvcContext) GetGroups() []apiserver.Group {
	return []apiserver.Group{
		{
			Path:       "/v1",   // 前缀
			Middlwares: []echo.MiddlewareFunc{},  // 组内必须中间件
			Routers: []apiserver.Router{
				{
					Method:  http.MethodPost,   // 请求method
					Path:    "/login/one",   // 路径，相对路径，全路径是 前缀+此属性
					Name:    "登录第一步",
					Handler: loginApi.LoginOne,
				},
			},
		},
	}
}
```

### 中间件

全局中间件定义在`svc.GetMiddlewares`中，此中定义的中间件会应用到整个服务中。

中间见示例:

```go
func Online() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			t := c.Request().Header.Get(echo.HeaderAuthorization)
			op, err := repository.NewHauthorizedOP()
			if err != nil {
				log.Fatal(err)
			}
			r := op.Check(t)
			if !r {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "invalid or expired jwt",
					Internal: nil,
				}
			}
			return next(c)
		}
	}
}
```

### 验证器

验证器是提供参数验证使用，官方提供的使用`https://pkg.go.dev/github.com/go-playground/validator/v10`。也可以提供自定义验证器，自定义的验证器加入到`svc.GetValidators`中。

示例:

```go
// idValid id验证器
func idValid(fl validator.FieldLevel) bool {
	switch fl.Field().Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64:
		return true
	default:
		return regexp.MustCompile(`^[1-9]\d*$`).MatchString(fl.Field().String())
	}
}

func (s *SvcContext) GetValidators() map[string]validator.Func {
	return map[string]validator.Func{}{
        "idValid": idValid,
    }
}
```

## 待实现

- 从Etcd中加载配置

