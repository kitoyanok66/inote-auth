package grpc

import (
	"log"
	"net"

	pb "github.com/kitoyanok66/inote-protos/proto/auth"
	"google.golang.org/grpc"
)

func RunGRPC(handler *AuthHandler, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterAuthServiceServer(server, handler)

	log.Printf("Auth gRPC server is running on port %s", port)
	return server.Serve(lis)
}
