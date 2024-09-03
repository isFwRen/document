#### gRPC gateway

#### 概述

随着微服务框架的流行，gRPC作为一种高性能、跨语言的远程过程调用（RPC）框架被广泛应用。但是，gRPC并不适用于所有应用场景。例如，**当客户端不支持gRPC协议时，或者需要将gRPC服务暴露给Web应用程序时，需要将一种RESTful API转换为gRPC的方式**。因此，gRPC网关应运而生。

#### gRPC网关在go-zero中的实现

go-zero中的gRPC网关是一个http服务器，它将RESTful API转换为gRPC请求，然后将gRPC响应转换为RESTful API。大致流程如下：

1. 从proto文件中解析出gRPC服务的定义。
2. 从配置文件中解析出gRPC服务的http映射规则。
3. 根据gRPC服务定义的http映射规则，生产gRPC服务的http处理器。
4. 启动http服务器，处理http请求。
5. 将http请求转换为gRPC请求。
6. 将gRPC响应转换为http响应。
7. 返回http响应。

#### Generate ProtoSet files（.pb）

- example command without external imports

```
protoc --descriptor_set_out=hello.pb hello.proto
```

- example command with external imports

```
protoc --include_imports --proto_path=. --descriptor_set_out=hello.pb hello.proto
```

#### GatewayConf

| 名称      | 说明          | 类型       | 是否必填 | 示例               |
| --------- | ------------- | ---------- | -------- | ------------------ |
| RestConf  | rest 服务配置 | RestConf   | 是       | 参考[基础服务配置] |
| Upstreams | gRPC 服务配置 | []Upstream | 是       |                    |
| Timeout   | 超时时间      | duration   | 否       | `5s`               |

#### Upstream

| 名称      | 说明                                  | 类型           | 是否必填 | 示例            |
| --------- | ------------------------------------- | -------------- | -------- | --------------- |
| Name      | 服务名称                              | string         | 否       | `demo1-gateway` |
| Grpc      | gRPC 服务配置                         | RpcClientConf  | 是       | 参考[RPC 配置]  |
| ProtoSets | proto 文件列表                        | []string       | 否       | `["hello.pb"]`  |
| Mappings  | 路由映射,不填则默认映射所有 grpc 路径 | []RouteMapping | 否       |                 |

##### 基础服务配置

在go-zero框架中，rest服务配置主要用于定义一个RESTful API服务的基础配置，如服务地址、超时设置、中间件等。配置文件通常以YAML或JSON格式编写，并且通过在代码中读取配置文件来初始化**rest.Server**

一个典型的 `rest` 服务配置文件可能看起来像这样：

```yaml
Name: example-api
Host: 0.0.0.0
Port: 8080
Timeout: 3000ms

# CORS 配置
Cors:
  allow-origin: "*"
  allow-methods: "GET,POST,PUT,DELETE"
  allow-headers: "Content-Type,Authorization"
  expose-headers: "X-Total-Count"
  allow-credentials: true

# Metrics 配置
Metrics:
  Host: 0.0.0.0
  Port: 9090

# Middlewares
Middlewares:
  - Logger
  - Recover
  - Auth

# JWT 配置
JwtAuth:
  AccessSecret: your-access-secret
  AccessExpire: 3600

# Tracing 配置
Trace:
  Name: example-api
  Endpoint: http://localhost:14268/api/traces

# 其他自定义配置
CustomConfig:
  RedisHost: localhost:6379
  MySQL:
    DataSource: "user:password@tcp(127.0.0.1:3306)/dbname"
```

##### 配置项解释

- **Name：**服务的名称。通常用于废物的标识和日志输出。
- **Host：**服务监听的IP地址。`0.0.0.0`表示监听所用网路接口。
- **Port：**服务监听的端口号。
- **Timeout：**请求超时时间。通常用ms（毫秒）为单位。
- **Cors：**配置跨域资源共享（CORS）。包括允许的域、方法、请求头和是否允许凭证。
- **Metrics：**配置服务监控的端口和地址，用于暴露Prometheus监控数据。
- **Middlewares：**指定使用的中间件列表。**Logger**用于日志记录，**Recover**用于错位恢复，**Auth**用于身份验证。
- **JwtAuth：**配置JWT鉴权。包括**AccessSecret**（用于签名的密钥）和**AccessExpire**（token过期时间）。
- **Trace：**配置分布式追踪系统（如Jeager）的信息。
- **CustomConfig：**用于存放其他自定义配置项，如Redis和Mysql的连接信息。

