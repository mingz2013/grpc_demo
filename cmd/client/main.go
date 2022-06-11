package main

import (
	"context"
	"github.com/liyue201/grpc-lb/balancer"
	"github.com/mingz2013/grpcdemo/pb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

func main() {
	balancer.InitConsistentHashBuilder(balancer.DefaultConsistentHashKey) // 注册一致性hash balancer，

	hostname := os.Getenv("HOSTNAME")

	// 创建etcd客户端
	cli, cerr := clientv3.NewFromURL("http://etcd:2379")
	if cerr != nil {
		log.Fatalln(cerr)
	}

	// 创建 resolver
	etcdResolver, err := resolver.NewBuilder(cli)
	if err != nil {
		log.Fatalln(err)
	}

	// 创建grpc连接, 使用etcd resolver，并配置balancer 策略

	//grpcDialCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	//defer cancel()
	conn, gerr := grpc.DialContext(
		context.Background(),
		"etcd:///foo/bar/my-service",
		//grpc.WithKeepaliveParams(keepalive.ClientParameters{
		//	Timeout: time.Minute,
		//}),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy" : "`+balancer.ConsistentHash+`"}`), // 使用一致性hash balancer，
	)
	if gerr != nil {
		log.Fatalln(gerr)
	}

	defer conn.Close()

	// grpc 服务的 greeter 客户端
	greeterClient := pb.NewGreeterClient(conn)
	//echoClient := pb.NewEchoClient(conn)
	//gateClient := pb.NewGateClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	//echoClient.SayHello(context.WithValue(context.Background(), balancer.DefaultConsistentHashKey, hostname), &pb.EchoRequest{
	//	Message: hostname,
	//})

	//gclient, err := gateClient.Route(context.WithValue(context.Background(), balancer.DefaultConsistentHashKey, hostname))
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//go func() {
	//	defer func() {
	//		err := recover()
	//		if err != nil{
	//			log.Println(err)
	//		}
	//	}()
	//
	//
	//	for {
	//		resp, err := gclient.Recv()
	//		if err != nil {
	//			log.Fatalln(err)
	//		}
	//		log.Println(resp.Uri, resp.Payload)
	//
	//	}
	//}()

	for {
		func() {

			defer func() {
				err := recover()
				if err != nil {
					log.Println("recover err:", err)
				}

			}()

			<-time.After(2 * time.Second)

			r, err := greeterClient.SayHello(
				context.WithValue(context.Background(), balancer.DefaultConsistentHashKey, hostname), //  传入一致性hash的值
				&pb.HelloRequest{
					Name: hostname,
				})

			if err != nil {
				log.Println(err)
				return
			}

			log.Println("greeting: ", r.GetMessage())

		}()

	}

}
