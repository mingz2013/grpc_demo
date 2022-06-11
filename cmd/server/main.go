package main

import (
	"context"
	"fmt"
	"github.com/mingz2013/grpcdemo/pb"
	"github.com/mingz2013/grpcdemo/servers"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func serveGreeterServer() {
	l, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer(
	//	grpc.KeepaliveParams(keepalive.ServerParameters{
	//	MaxConnectionAge: time.Minute,
	//}),
	)

	pb.RegisterGreeterServer(s, servers.GreeterServer)
	//pb.RegisterEchoServer(s, servers.EchoServer)
	//pb.RegisterGateServer(s, servers.GateServer)

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

func main() {

	hostname := os.Getenv("HOSTNAME")

	log.Println("hostname: ", hostname)

	// 创建etcd客户端
	cli, cerr := clientv3.NewFromURL("http://etcd:2379")
	if cerr != nil {
		log.Fatalln(cerr)
	}

	//session, err := concurrency.NewSession(cli)

	// 创建endpoints管理
	em, err := endpoints.NewManager(cli, "foo/bar/my-service")
	if err != nil {
		log.Fatalln(err)
	}
	// 租约
	resp, err := cli.Grant(context.TODO(), 10)

	// 添加节点, 可设置租约
	err = em.AddEndpoint(
		context.TODO(),
		"foo/bar/my-service/"+hostname,
		endpoints.Endpoint{Addr: hostname + ":8000"},
		clientv3.WithLease(resp.ID),
	)

	if err != nil {
		log.Fatalln(err)
	}

	ch, kaeerr := cli.KeepAlive(context.TODO(), resp.ID)
	if kaeerr != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			ka := <-ch
			fmt.Println("ttl:", ka.TTL, ka.ID)
		}
	}()

	defer func() {
		// 删除节点
		err = em.DeleteEndpoint(context.TODO(), "foo/bar/my-service/"+hostname)
		if err != nil {
			log.Fatalln(err)
		}

		////一次修改多个
		//em.Update(context.TODO(), []*endpoints.UpdateWithOpts{
		//
		//})
	}()

	//  启动服务
	serveGreeterServer()

}
