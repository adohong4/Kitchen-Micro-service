package main

import (
	"log"
	"net"

	handler "github.com/adohong4/Kitchen-Microservice/services/orders/handler/orders"
	"github.com/adohong4/Kitchen-Microservice/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)

	log.Println("Starting gRPC Server on", s.addr)

	return grpcServer.Serve(lis)
}
