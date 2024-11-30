package main

import (
	"context"
	"fmt"
	"github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var pods []*protos.Pod // todo: where are we fetching this from?
	responses := make(map[*protos.Pod]*protos.NodeMetrics)

	for _, pod := range pods {
		host := fmt.Sprintf("%s:50051", pod.Metadata.Name) // todo: what is the address going to be?
		conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("could not connect: %v", err)
		}
		defer conn.Close()
		c := protos.NewKubeletClient(conn)

		in := protos.Empty{}
		resp, err := c.Metric(context.Background(), &in)

		if err != nil {
			log.Fatalf("metrics: could not connect: %v", err)
		}

		responses[pod] = resp
	}

	fmt.Printf("Collected metrics. What shall we do with them?") // todo
}
