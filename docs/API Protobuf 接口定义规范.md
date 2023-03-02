对外的接口 定义我们使用 protobuf ， 总入口在 **api.proto** 文件里面
![image.png](https://cdn.nlark.com/yuque/0/2023/png/85503/1677663041985-6fc07422-643f-4c59-b6ba-0976b0ad2a0a.png#averageHue=%23929292&clientId=u03a8adbb-3482-4&from=paste&height=328&id=u262e45d6&name=image.png&originHeight=328&originWidth=394&originalType=binary&ratio=1&rotation=0&showTitle=false&size=31685&status=done&style=none&taskId=u63b46c43-e12a-4bce-a6f2-0786a1f9927&title=&width=394)
**api.proto 文件里面只定义 对应的模块的 Service rpc 方法**

```protobuf
service apiPostService {
  // 岗位管理 新增一个岗位
  rpc PostAddOne (PostAddRequest) returns (PostAddResponse) {option (api.post) = "/post/PostAddOne";}
  // 岗位管理 删除多个岗位
  rpc PostDel (PostDelRequest) returns (PostDelResponse) {option (api.post) = "/post/PostDel";}
  // 岗位管理 查询岗位列表
  rpc PostQry (PostQryRequest) returns (PostQryResponse) {option (api.post) = "/post/PostQry";}
  // 岗位管理 修改单个岗位信息
  rpc PostModify (PostModifyRequest) returns (PostModifyResponse) {option (api.post) = "/post/PostModify";}
}
```

- service 名称的命名为： Api[模块]service
- rpc 方法 名称定义 一定要语意化
- rpc 对应 方法 统一为的 **api.post**
- api.post 定义的path 规则为：/[模块]/[rpc 方法]

> 示例：参考上面的代码块

**具体的 rpc 方法的参数和响应定义到一个单独的模块文件中**
创建文件的规则为： **api_[模块].proto**

**使用创建的 模块文件**
在 **api.proto**文件中 import 进去

**问题**

1. **import 后 golang 不能智能跳转**

手动配置 导入 idl 文件路径
![image.png](https://cdn.nlark.com/yuque/0/2023/png/85503/1677663763390-b3f25e1f-e939-41b0-b9ef-614ef171df35.png#averageHue=%23cececd&clientId=u03a8adbb-3482-4&from=paste&height=1078&id=u752ea4e2&name=image.png&originHeight=1078&originWidth=1932&originalType=binary&ratio=1&rotation=0&showTitle=false&size=210847&status=done&style=none&taskId=u0e8afafd-854e-4e27-b1d0-4a5bdd0ebaf&title=&width=1932)
