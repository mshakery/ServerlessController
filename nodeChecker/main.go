package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	protos.UnimplementedNodeCheckerServer
}

func ReallocatePods(ctx context.Context, client *clientv3.Client, nodeName string) {
	response, _ := etcd.ReadFromEtcd(client, ctx, "/cluster/resources/pod/", true)
	for _, kv := range response.Kvs {
		if strings.Count(string(kv.Key), "/") == 6 && strings.HasSuffix(string(kv.Key), "/worker") {
			workerProto := protos.Worker{}
			err := proto.Unmarshal(kv.Value, &workerProto)
			if err != nil {
				panic(err)
			}
			if nodeName == workerProto.GetWorker() {
				podNamespace := strings.Split(string(kv.Key), "/")[4]
				podName := strings.Split(string(kv.Key), "/")[5]
				var podProto protos.Pod
				podKey := fmt.Sprintf("/cluster/resources/pod/%s/%s", podNamespace, podName)
				etcd.ReadOneFromEtcdToPb(client, ctx, podKey, &podProto)
				if podProto.Spec.RestartPolicy != "Never" {
					etcd.WriteToEtcdFromPb(client, ctx, string(kv.Key), &protos.Worker{Worker: "-1"})
				}
			}
		}

	}
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

	NodeUnschedulableKey := fmt.Sprintf("/cluster/resources/node/%s/unschedulable", in.GetName())
	if time.Now().Sub(parsedTime).Seconds() > 45 { //TODO: lock
		//etcd.AcquireLock(client, NodeUnschedulableKey, 0, true)
		etcd.WriteToEtcdFromPb(client, ctx, NodeUnschedulableKey, &protos.Unschedulable{Condition: true})
		ReallocatePods(ctx, client, in.GetName())

	} else {
		//etcd.ReleaseLock(client, NodeUnschedulableKey)
		etcd.WriteToEtcdFromPb(client, ctx, NodeUnschedulableKey, &protos.Unschedulable{Condition: false})
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
