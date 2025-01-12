package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"slices"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	protos.UnimplementedAuthorizationServer
}

func (s *server) Authorize(ctx context.Context, in *protos.AuthorizationRequest) (*protos.Response, error) {
	/*
		should check if user has permission to perform action.
	*/
	client, err := etcd.ConnectToEtcd()
	if err != nil {
		panic("Could not connect to etcd. ")
	}
	defer client.Close()

	whatToRead := fmt.Sprintf("/cluster/resources/role_binding/") /* todo verify */
	read, err := etcd.ReadFromEtcd(client, ctx, whatToRead, true)
	if err != nil {
	}

	resp := protos.Response{}

	// var userRoles []*protos.Role
	var hasPermission bool

	if read.Count == 0 {
		resp.Code = -1
		resp.Status = "No role bindings found.."
		return &resp, nil
	} else {
		for _, kv := range read.Kvs { /* todo test */
			fmt.Printf("itering over: key: %s, val: %s\n", kv.Key, kv.Value)
			roleReference := protos.RoleBinding{}
			err := proto.Unmarshal(kv.Value, &roleReference)
			if err != nil {
				return nil, err
			}
			if slices.Contains(roleReference.Subjects, in.Uid) { /* todo refactor add to slice and then read all and check in struct */
				role := protos.Role{}
				whatToRead := fmt.Sprintf("/cluster/resources/role/%s", roleReference.RoleRef)
				err := etcd.ReadOneFromEtcdToPb(client, ctx, whatToRead, &role)
				if err != nil {
					return nil, err
				}
				for _, rule := range role.Rules {
					hasPermission = checkHasPermission(rule, in.ClientRequest)
					if hasPermission {
						break
					}
				}
				// userRoles = append(userRoles, &role)
				if hasPermission {
					break
				}
			}
		}
	}

	if hasPermission {
		/* todo the rest */
		write2Etcd, err := callWrite2Etcd(ctx, in.ClientRequest, in.Uid)
		if err != nil {
			return nil, err
		}
		return write2Etcd, nil
	}
	resp.Code = -1
	resp.Status = "not enough permission"
	return &resp, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterAuthorizationServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func checkHasPermission(rule *protos.PolicyRule, request *protos.ClientRequest) bool {
	switch request.OneofResource.(type) {
	case *protos.ClientRequest_Deployment:
		if slices.Contains(rule.Resources, "deployment") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Role:
		if slices.Contains(rule.Resources, "role") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_RoleBinding:
		if slices.Contains(rule.Resources, "rolebinding") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_User:
		if slices.Contains(rule.Resources, "user") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Node:
		if slices.Contains(rule.Resources, "node") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Pod:
		if slices.Contains(rule.Resources, "pod") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	}
	return false
}

func callWrite2Etcd(ctx context.Context, client_request *protos.ClientRequest, uid string) (*protos.Response, error) {
	conn, err := grpc.NewClient("write-to-etcd.default.10.103.172.226.sslip.io:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewWriteToEtcdClient(conn)

	applyReq := &protos.ApplyRequest{
		ClientRequest: client_request,
	}

	resp, err := c.Apply(ctx, applyReq)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	return resp, nil
}
