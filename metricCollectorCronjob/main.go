package main

import (
	"context"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	"github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"path"
	"strings"
)

func main() {
	client, err := etcd.ConnectToEtcd()
	if err != nil {
		panic("Could not connect to etcd. ")
	}
	defer client.Close()
	ctx := context.TODO()
	read, err := etcd.ReadFromEtcd(client, ctx, "/cluster/resources/node/", true)
	if err != nil {
	}
	for _, kv := range read.Kvs {
		if strings.Count(string(kv.Key), "/") == 4 {
			NodeName := path.Base(string(kv.Key))
			conn, err2 := grpc.NewClient("metric-collector:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err2 != nil {
				log.Fatalf("metrics: could not connect: %v", err2)
			}
			c := protos.NewMetricCollectorClient(conn)

			in := protos.NodeName{Name: NodeName}
			_, err3 := c.GatherMetric(context.Background(), &in)

			if err3 != nil {
				log.Fatalf("metrics: could not connect: %v", err2)
			}
			log.Printf("metrics are collected for node %s\n", NodeName)
		}
	}
}
