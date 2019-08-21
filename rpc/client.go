package main

import (
	"context"
	"fmt"
	"log"

	m "github.com/RifeWang/go-study/rpc/grpc"
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
	RPCClient = m.NewRPCServiceClient(conn)
}

func main() {
	req := &m.ReqBody{
		UserId:   "222",
		Page:     1,
		Pagesize: 10,
	}
	res, err := RPCClient.QueryUserOrders(context.Background(), req)
	if err != nil {
		log.Println("rpc call error:", err)
	}
	fmt.Println(res)
}
