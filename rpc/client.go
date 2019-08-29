package main

import (
	"context"
	"fmt"
	"log"
	"time"

	m "gomod/rpc/grpc" // 导入 protoc 编译生成的代码包
	"google.golang.org/grpc"
)

// RPCClient ...
var RPCClient m.RPCServiceClient

func init() {
	rpcServerAddr := "localhost:9999"
	conn, err := grpc.Dial(rpcServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	RPCClient = m.NewRPCServiceClient(conn) // 生成 gRPC 客户端
}

func main() {
	// 构造请求数据
	req := &m.ReqBody{
		UserId:   "uuid-222",
		Page:     1,
		Pagesize: 10,
	}

	// 客户端发起请求
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second) // 设置 1s 超时
	defer cancel()
	res, err := RPCClient.QueryUserOrders(ctx, req)
	if err != nil {
		log.Println("rpc call error:", err)
	}
	fmt.Println(res)
}
