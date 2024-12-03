package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"slices"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
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
			fmt.Printf("k: %s, v: %s\n", kv.Key, kv.Value)
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
		return &resp, nil
	}
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
		if slices.Contains(rule.Resources, "deployments") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Role:
		if slices.Contains(rule.Resources, "roles") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_RoleBinding:
		if slices.Contains(rule.Resources, "rolebindings") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_User:
		if slices.Contains(rule.Resources, "users") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Node:
		if slices.Contains(rule.Resources, "nodes") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	case *protos.ClientRequest_Pod:
		if slices.Contains(rule.Resources, "pods") && slices.Contains(rule.Verbs, request.Operation) {
			return true
		}
	}
	return false
}
