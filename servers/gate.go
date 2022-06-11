package servers

import (
	"github.com/mingz2013/grpc_demo/pb"
	"log"
)

type gateServer struct {
	pb.UnimplementedGateServer
}

func (g gateServer) Route(server pb.Gate_RouteServer) error {
	//panic("implement me")
	for {
		message, err := server.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(message)
	}

}

var GateServer = &gateServer{}
