// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	authentication2 "github.com/mshakery/ServerlessController/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	authentication2.UnimplementedAuthenticationServer
}

func (s *server) Authentication(_ context.Context, in *authentication2.AuthenticationRequest) (authentication2.Response, error) {
	/*
		should reade from ETCD. key is token. checks if value exists.
		if exists, value is uid.
		then makes request for authorization.

		creates Authorizationrequest. uid = uid from etcd.
		then calls authorizationrequest.
	*/
	return authentication2.Response{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	authentication2.RegisterAuthenticationServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
