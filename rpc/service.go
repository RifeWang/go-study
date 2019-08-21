package main

import (
	"context"
	"fmt"
	"log"
	"net"

	m "github.com/RifeWang/go-study/rpc/grpc"
	"google.golang.org/grpc"
)

func main() {
	port := ":9999"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("listen on port", port)

	server := grpc.NewServer()
	m.RegisterRPCServiceServer(server, &srv{})
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("serve error: %v", err)
	}
}

type srv struct{}

func (s *srv) QueryUserOrders(ctx context.Context, req *m.ReqBody) (*m.UserOrders, error) {
	fmt.Println("receive req:", req)

	result := &m.UserOrders{}
	result.Id = 1
	result.Username = "wang"
	result.Email = "123456@gmail.com"
	result.Phone = "188xxxxxxxx"

	order := &m.UserOrders_Order{}
	order.OrderId = 11111
	order.Info = "info"
	order.Status = m.UserOrders_Order_SATAUS_SUCCESS

	result.Orders = append(result.Orders, order)
	result.Orders = append(result.Orders, order)
	result.Orders = append(result.Orders, order)

	return result, nil
}
