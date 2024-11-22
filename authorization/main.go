package main

import (
	"context"
	"flag"
	"fmt"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	protos.UnimplementedAuthorizationServer
}

func (s *server) Authorization(ctx context.Context, in *protos.AuthorizationRequest) (protos.Response, error) {
	/*
		should check if user has permission to perform action.
	*/
	client, err := etcd.ConnectToEtcd()
	if err != nil {
		panic("Could not connect to etcd. ")
	}
	defer client.Close()

	whatToRead := fmt.Sprintf("%d:read", in.Uid) /* todo <-- what should this be? */
	read, err := etcd.ReadFromEtcd(client, ctx, whatToRead)
	if err != nil {
	}

	resp := protos.Response{}

	if read.Count == 0 {
		resp.Code = -1
		resp.Status = "User does not have permission."
		return resp, nil
	} else {
		for _, kv := range read.Kvs { /* todo test */
			fmt.Printf("k: %s, v: %s\n", kv.Key, kv.Value)
			resp.Code = 0
			resp.Status = "User has permission."
		}
	}

	/* todo the rest */
	return resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterAuthorizationServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
