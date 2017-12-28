package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	rpc "github.com/radu-matei/grpc-gateway/rpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type echoService struct {
}

func (e *echoService) Echo(ctx context.Context, m *rpc.Message) (*rpc.Message, error) {
	return &rpc.Message{
		Value: fmt.Sprintf("Hello from the gRPC server. You said: %s", m.Value),
	}, nil
}

func newServer() *echoService {
	return new(echoService)
}

func main() {
	flag.Parse()

	s := newServer()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to start listening %v", err)
	}
	grpcServer := grpc.NewServer()

	rpc.RegisterEchoServiceServer(grpcServer, s)

	grpcServer.Serve(l)
}
