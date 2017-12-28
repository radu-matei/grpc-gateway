package main

import (
	"context"
	"flag"
	"log"

	rpc "github.com/radu-matei/grpc-gateway/rpc"
	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := rpc.NewEchoServiceClient(conn)

	m, err := client.Echo(context.Background(), &rpc.Message{Value: "Testing from a gRPC client"})
	if err != nil {
		log.Fatalf("failed to get echo from grpc server: %v", err)
	}

	log.Printf("response from grpc server: %v", m.Value)
}
