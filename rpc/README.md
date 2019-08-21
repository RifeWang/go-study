1. 编写 `.proto` 文件
2. 下载对应系统的 protocol buffers compiler 编译器: https://github.com/protocolbuffers/protobuf/releases , 我的是 osx-x86_64 版本，解压。
   - 将 `/bin` 目录添加到系统 `PATH` 路径，以便于在命令行直接使用 `protoc` 命令。
   - (没必要) 复制 `include` 目录下的所有文件至项目的 `third_party` 目录，以便于编译时使用插件。
3. (没必要，会自动安装) 安装 go 插件: `go get -u github.com/golang/protobuf/protoc-gen-go`。
4. 进入 rpc 目录下，执行 `protoc-gen.sh` 文件中的命令编译生成代码，以供后续使用。
5. 导入上一步编译生成的代码包，以及 `google.golang.org/grpc` ，客户端和服务端代码实现。