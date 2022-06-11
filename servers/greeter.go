package servers

import (
	"context"
	"github.com/mingz2013/grpc_demo/pb"
	"log"
	"os"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

var GreeterServer = &greeterServer{}

func (g greeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("sayhello:", request.GetName())

	hostname := os.Getenv("HOSTNAME")
	return &pb.HelloReply{
		Message: "Hello " + request.GetName() + ", My name is: " + hostname,
	}, nil

}
