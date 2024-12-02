package main

import (
	"context"
	"flag"
	"fmt"
	protos "github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

var (
	port = flag.Int("port", 50060, "Metric Port")
)

type server struct {
	protos.UnimplementedKubeletServer
}

var podsToRun []*protos.Pod

func (s *server) runAPod(ctx context.Context, pod *protos.Pod) (protos.Empty, error) {
	// command
	podsToRun = append(podsToRun, pod)
	return protos.Empty{}, nil
}

func (s *server) Metrics(ctx context.Context, in *protos.Empty) (*protos.NodeMetrics, error) {
	metrics := protos.NodeMetrics{}
	var result []*protos.PodMetrics
	for _, pod := range podsToRun {
		pm := protos.PodMetrics{}
		pm.Name = pod.GetMetadata().GetName()
		pm.Namespace = pod.GetMetadata().GetNamespace()
		CpuUsage, MemoryUsage := 0, 0
		for _, container := range pod.Spec.Containers {
			CCpuUsage, _ := strconv.Atoi(container.GetResources().Requests["cpu"])
			CMemoryUsage, _ := strconv.Atoi(container.GetResources().Requests["memory"])
			CpuUsage += CCpuUsage
			MemoryUsage += CMemoryUsage
		}
		// TODO: it must be the maximum usage and request
		pm.CpuUsage = strconv.Itoa(CpuUsage)
		pm.MemoryUsage = strconv.Itoa(MemoryUsage)
		result = append(result, &pm)
	}
	metrics.PodMetrics = result
	return &metrics, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterKubeletServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
