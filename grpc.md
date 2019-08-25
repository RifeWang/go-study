## protocol buffers
google 成熟的开源机制，用于对结构化数据进行序列化或反序列化。

注意：protocol buffers 不使用于大数据集，如果单条消息超过 1 MB ，则需要考虑其他替代方案。

1、在 `.proto` 后缀的文件中定义数据结构，数据被定义为 `message`：
```
message Person {
    string name = 1;
    int32 id = 2;
    bool has_ponycopter = 3;
}
```

2、定义了数据结构之后，就可以使用 protocol buffer 编辑器 `protoc` 去生成指定编程语言的数据访问类。这些为每个字段提供了简单的访问器（比如 `name()` 和 `set_name()` ），以及序列化或解析为原始 bytes 字节的方法。如果你选择的语言是 C++ ，编译器将生成一个 `Person` 类，在应用程序中就可以使用这个类去填充、序列化、检索 Person 的消息。

通过 `service` 定义 gRPC 服务，`message` 定义 RPC 方法、参数和返回信息：
```
// The greeter service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
    string name = 1;
}

// The response message containing the greetings
message HelloReply {
    string message = 1;
}
```

`protoc` 可以使用 gRPC 插件去生成代码，插件需要下载安装。

当前版本为 `proto3` 。

------
## Service definition

gRPC 允许定义四种类型的 service method：

1、Unary RPCs ，客户端发送单个请求然后接受到服务端的响应，就像一个正常的函数调用：
```
rpc SayHello(HelloRequest) returns (HelloResponse) {

}
```

2、Server streaming RPCs ，客户端发送请求，接受 stream 消息，gRPC 保证单个 RPC 调用中的消息排序：
```
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){

}
```

3、Client streaming RPCs ，客户端发送 stream 流消息，然后接受服务端响应：
```
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {

}
```

4、Bidirectional streaming RPCs ，双向流：
```
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){

}
```

-----
## Using the API surface
通过编译器生成服务端和客户端代码。
- 在服务端，服务器实现服务声明的方法，并运行 gRPC server 来处理客户端调用。gRPC基础设施负责解码传入请求，执行服务方法并对服务响应进行编码。
- 在客户端，客户端有一个本地对象称为 stub ，实现与服务端相同的方法，客户端在本地对象上调用这些方法，传递参数，gRPC 发送请求然后返回服务端响应。


## 同步 vs 异步
查询不同语言具体的文档

--------

## RPC life cycle
客户端调用服务端方法整个过程发生了什么。

### Unary RPC
- Once the client calls the method on the stub/client object, the server is notified that the RPC has been invoked with the client’s metadata for this call, the method name, and the specified deadline if applicable.
- The server can then either send back its own initial metadata (which must be sent before any response) straight away, or wait for the client’s request message - which happens first is application-specific.
- Once the server has the client’s request message, it does whatever work is necessary to create and populate its response. The response is then returned (if successful) to the client together with status details (status code and optional status message) and optional trailing metadata.
- If the status is OK, the client then gets the response, which completes the call on the client side.

### Server streaming RPC

### Client streaming RPC

### Bidirectional streaming RPC

### Deadlines/Timeouts
gRPC 客户端可以指定超时时间，对应错误信息为 `DEADLINE_EXCEEDED`，服务端可以查询是否超时或者还有多久超时。这取决于具体的语言实现。

### RPC termination
客户端和服务端都在本地独立判断调用是否成功，这可能会导致它们的结果并不一致。

### Cancelling RPCs
客户端和服务端可以随时取消 RPC ，取消意味着不再进行进一步处理，而不是回滚。

### Metadata
元数据是一组 RPC 调用信息的键值对，由客户端提供。对元数据的访问取决于具体语言实现。

### Channels
gRPC channel 提供到指定主机和端口号的服务端的连接，并且当创建客户端 stub 时被使用。客户端可以声明 channel 参数去调整 gRPC 的默认行为，比如是否开启消息压缩。channel 有状态，如 `connected` 和 `idle` 。
gRPC 如何处理关闭 channel 取决于具体语言实现，某些语言也提供了对 channel 状态的查询。

-----
# Authentication
TODO

-----
# Error handling and debugging

## standard error model
如果错误发生, gRPC 返回一个错误状态码和可选的错误描述信息。

## Richer error model
官方错误模型受所有客户端服务端支持，独立于 gRPC 数据格式（protocol buffers），但是能力受限。
如果你使用 protocol buffers 数据格式，你可以使用更广泛的错误模型，已支持很多语言。

## Error status codes
参考 https://www.grpc.io/docs/guides/error/