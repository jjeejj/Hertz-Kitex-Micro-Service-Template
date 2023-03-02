
### 背景

1. 后端使用 protobuf 声明接口文件
2. 前端使用 TypeScript 进行类型声明

**所以 这里前端进行接口调用的时候 直接使用 Protobuf 生成的 TypeScript 声明文件**

### 好处

1. 前端不需要在写接口请求参数 和 返回参数
2. 前端接口请求参数 和返回 参数可以做到智能提醒
3. 后端 不用单独维护接口文档

### 生成方法-新

1. 安装 protobuf

> [protobuf](https://github.com/protocolbuffers/protobuf/releases)

2. 项目中配置 ts-proto

> [ts-proto](https://github.com/stephenh/ts-proto#esm)

`protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_opt=forceLong=bigint --ts_proto_opt=esModuleInterop=true --ts_proto_out=./src/types/api -I /idl/  api.proto`
> 执行 命令 **pnpm pc** 即可以生成
> 下面的命令 只需要修改  -I idl 文件路径
> 还需要创建目标  api 文件夹路径

1. 生成之后 就可以看到对应的什么文件

**api.ts**是所有请求的总入口
会按照 protobuf 中定义的 Service 生成对应的 interface 和 对应的 class 实现

4. 发送请求使用

- 实例化对应模块的 interface Service 的实现

```typescript
export interface ApiService {
}

export class ApiServiceClientImpl implements ApiService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "api.apiService";
    this.rpc = rpc;
  }
}

let apiServiceClientImpl = new ApiServiceClientImpl(rpcImpl, {service: api})
```

> 这里实例化的时候 传入的2个参数
> rpc: 发送请求的 rpc 接口实现，用 axios
> service: 服务模块，为 api_auth.ts  文件名称中 api 后面的 部门： **auth**

- 实现生成定义的 RPC 接口 用来实际发送请求

```typescript
interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
}
```

> 实现 rpc 发送请求的时候 要把 service 和 method 参数拼接出 路径 /service/method
> 如果没有特殊说明，请求方法一般 为 post 请求
