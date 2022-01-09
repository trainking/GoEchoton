# 约束

- [约束](#约束)
  - [命名约束](#命名约束)
    - [api](#api)
      - [请求参数](#请求参数)
      - [返回数据](#返回数据)
      - [惯例](#惯例)
        - [列表](#列表)
        - [映射](#映射)
        - [多语言](#多语言)
  - [gitflow](#gitflow)
    - [提交约束](#提交约束)
      - [模板](#模板)
      - [type](#type)
    - [分支约束](#分支约束)
      - [远程分支](#远程分支)
      - [本地分支](#本地分支)

## 命名约束

### api

一个api的结构：

```
├─api           // api的handler包
│  └─login      // 具体业务分包
├─apply         // 请求参数包，Request
├─config        // 配置结构体
├─reply         // 返回参数包, Response
├─svc           // ServerContext
└─authserver.go // main 入口
```

#### 请求参数

请求参数使用`XXXXApply`命名：

```go
type LoginOneApply struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}
```

支持的tag:

- `json` 请求body使用`appliction/json`格式
- `query` 解析路径后`?`的`query params`
- `param` 解析路由定义中`/:id`的参数
- `header` 解析请求头中的`header`
- `form` 解析请求表单中的参数
- `xml` 解析`application/xml`格式
- `validate` 验证参数，[参考](https://pkg.go.dev/github.com/go-playground/validator/v10)

#### 返回数据

返回数据以`XXXXReply`命名:

```go
type (
	LoginOneReply struct {
		Id   int64  `json:"id"`
		Sign string `json:"sign"`
	}
)
```

#### 惯例

##### 列表

- 分页参数使用`page`表示页码，`size`表示每页显示数目
- `count`表示总数目
- `total` 字段表示总计，这是个对象，启用对应列数据，以列的key命名
- `subtotal` 表示小计，结构与`total`一样
- `list` 表示列表数据

##### 映射

- 下拉列表使用映射形式，结构`[{"value": 1, "title":"dddd"}]`
- `value`表示传输后台的值
- `title`表示显示值

##### 多语言

- 需要其他语言返回，需在请求头中加入`language`字段
- 值使用ISO 639标准[ISO 639](https://zh.wikipedia.org/wiki/ISO_639-1)

## gitflow

### 提交约束

#### 模板

```
type:[(scope)] subject
// 空行
[body]
// 空行
[Footer]
```

#### type

|类型|类别|说明|
|:--|--|--|
|feat|代码类|新增功能|
|fix|代码类|bug修复|
|perf|代码类|提高代码性能变更|
|style|代码类|格式化代码，美化代码|
|merge|代码类|合并代码，解决冲突|
|refator|代码类|其他代码变更，口袋选项，不属于上述选项时，选择refator|
|test|非代码类|提交测试代码，单元测试案例等|
|ci|非代码类|部署内容的相关改动，如修改部署文件，ci配置文件|
|docs|非代码类|文档变更|
|chore|非代码类|其他变更，口袋选项，不属于上述选项时，选择chore变更|

### 分支约束

#### 远程分支

- master/main 该分支下是经过测试发布的正式内容，运行项目中静态内容
- develop 开发分支，拥有最全最新的功能，未测试

#### 本地分支

- feature 新特性分支，命名格式`feature/xxxx`，xxx使用特性简称，或模块名如`feature/login`
- release 测试分支，明明格式`release/v1.10`, 使用版本名，测试完结之后将所有修改合并到`develop`和`master`
- hotfix 热修复分支，需要热修复，以`hotfix/21212`, 命名以bug编号作为标记