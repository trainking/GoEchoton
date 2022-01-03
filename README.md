# GoEchoton - 标准api项目示范

- [GoEchoton - 标准api项目示范](#goechoton---标准api项目示范)
  - [概述](#概述)
  - [项目规范](#项目规范)
    - [目录](#目录)
    - [规范](#规范)
  - [包实现](#包实现)

## 概述

构造Api项目开发规范

## 项目规范

### 目录

```
├─configs
├─docs
│  ├─devel
│  │  └─zh-CN
│  ├─guide
│  │  ├─en-US
│  │  └─zh-CN
│  │      ├─api
│  │      └─apiserver
│  └─images
├─internal
│  ├─authserver
│  │  ├─api
│  │  │  └─login
│  │  ├─apply
│  │  ├─config
│  │  ├─reply
│  │  └─svc
│  ├─pkg
│  │  ├─apply
│  │  └─reply
│  └─userrpc
│      ├─handler
│      ├─service
│      ├─svc
│      ├─types
│      └─userclient
│-pkg
│   ├─apiserver
│   ├─arpcclient
│   ├─arpcserver
│   └─etcdx
│-gitignore
│-CHANGELOG
│-go.mod
│-Makefile
└─README.md
```

### 规范

- [约束](./docs/devel/zh-CN/roule.md)

## 包实现

- apiserver: 定义api服务器实践[apiserver](./docs/guide/zh-CN/apiserver/README.md)
- arpcserver: 定义arpc服务器实践[arpcserver](./docs/guide/zh-CN/arpcserver/README.md)
- arpcclient: 定义arpc客户端实践[arpcclient](./docs/guide/zh-CN/arpcclient/README.md)
