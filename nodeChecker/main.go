package main

import (
	"context"
	"flag"
	"fmt"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	protos.UnimplementedNodeCheckerServer
}

func (s *server) CheckNode(ctx context.Context, in *protos.NodeNamee) (*protos.Empty, error) {
	client, err := etcd.ConnectToEtcd()
	if err != nil {
		panic("Could not connect to etcd. ")
	}

	NodeLastUpdateKey := fmt.Sprintf("/cluster/resources/node/%s/update/time", in.GetName())
	lastUpdateTime, err := etcd.ReadFromEtcd(client, ctx, NodeLastUpdateKey, false)
	if err != nil || lastUpdateTime.Count < 1 {
		log.Fatalf("node %s does not have update key: %v", in.GetName(), err)
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(lastUpdateTime.Kvs[0].Value))
	if err != nil {
		log.Fatalf("Time parse error node %s: %v", in.GetName(), err)
	}
	if time.Now().Sub(parsedTime).Seconds() > 45 {
		NodeUnschedulableKey := fmt.Sprintf("/cluster/resources/node/%s/unschedulable", in.GetName())
		val := protos.Unschedulable{Condition: true}
		etcd.WriteToEtcdFromPb(client, ctx, NodeUnschedulableKey, &val)
	}

	return &protos.Empty{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterNodeCheckerServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
