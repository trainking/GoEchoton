# GoEchoton
Golang的项目开发骨架


## web-layout Web项目骨架

* `/bin` 主要存放编译后文件
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
* `server.go` server的入口文件
