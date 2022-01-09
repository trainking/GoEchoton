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
├─cmd
│  └─initctl
├─configs
├─docs
│  ├─devel
│  │  └─zh-CN
│  ├─guide
│  │  ├─en-US
│  │  └─zh-CN
│  │      ├─api
│  │      ├─apiserver
│  │      ├─arpcclient
│  │      └─arpcserver
│  └─images
├─internal
│  ├─pkg
│  │  ├─apply
│  │  └─reply
│  ├─rpc
│  │  ├─order
│  │  └─user
│  │      ├─skeleton
│  │      │  ├─handler
│  │      │  └─service
│  │      ├─svc
│  │      ├─userclient
│  │      ├─userrpc
│  │      └─userstub
│  └─server
│      ├─auth
│      │  ├─api
│      │  │  └─login
│      │  ├─apply
│      │  ├─config
│      │  ├─reply
│      │  └─svc
│      └─order
│-pkg
│   ├─apiserver
│   ├─arpcclient
│   ├─arpcserver
│   └─etcdx│-gitignore
│-CHANGELOG
│-go.mod
│-Makefile
└─README.md
```

### 规范

- [约束](./docs/devel/zh-CN/roule.md)

## 包实现

- apiserver: api服务器实践[apiserver](./docs/guide/zh-CN/apiserver/README.md)
- arpcx: rpc实践[arpcx](./docs/guide/zh-CN/arpcx/README.md)