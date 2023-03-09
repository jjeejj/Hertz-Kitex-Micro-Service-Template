# Hertz-Kitex-Micro-Service-Template

Hertz ， Kitex 微服务模版

## 使用文档

* [如何开发一个模块接口](./docs/如何开发一个模块接口.md)
* [日志模块怎么使用](./docs/日志模块怎么使用.md)
* [API Protobuf 接口定义规范](./docs/API%20Protobuf%20接口定义规范.md)
* [Protobuf生成TypeScript 声明](./docs/Protobuf生成TypeScript%20声明.md)


## 说明

1. 当项目比较少的时候，IDL 声明文件可以和仓库放在一起；项目较大的时候 可以把 IDL 单独进行管理

## 服务说明

### API 服务

提供对外部的接口

### OSS 对象服务

提供媒体资源上传获取的服务，只提供 rpc 接口供其他内部 RPC 服务调用

支持的存储类型：
1. Minio  ✅
2. 阿里云  ❌

## 代码生成

客户端：`kitex  -thrift frugal_tag -module github.com/jjeejj/Hertz-Kitex-Micro-Service-Template ./idl/oss.thrift`

服务端：`kitex  -thrift frugal_tag -service oss -module github.com/jjeejj/Hertz-Kitex-Micro-Service-Template  -use github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/kitex_gen ../../../idl/oss.thrift`
