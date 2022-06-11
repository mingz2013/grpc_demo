package servers

import (
	"context"
	"github.com/mingz2013/grpcdemo/pb"
)

type echoServer struct {
	pb.UnimplementedEchoServer
}

func (e echoServer) SayHello(ctx context.Context, request *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{
		Message: request.Message,
	}, nil
}

var EchoServer = &echoServer{}
