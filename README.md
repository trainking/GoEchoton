# GoEchoton - 标准api项目示范

## 概述

构造Api项目开发规范

## 项目结构

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

## 包实现

- apiserver: 定义api服务器实践[apiserver](./docs/guide/zh-CN/apiserver/README.md)