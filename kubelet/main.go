package main

import (
	"context"
	"flag"
	"fmt"
	protos "github.com/mshakery/ServerlessController/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50060, "Metric Port")
)

type server struct {
	protos.UnimplementedKubeletServer
}

var podsToRun []string

func (s *server) runAPod(ctx context.Context, pod *protos.Pod) (protos.Empty, error) {
	// command
	podsToRun = append(podsToRun, pod.Metadata.Name)
	return protos.Empty{}, nil
}

func (s *server) Metrics(ctx context.Context, pod *protos.Empty) (*protos.PodMetrics, error) {
	metrics := protos.PodMetrics{
		Name:        "1",
		CpuUsage:    "",
		MemoryUsage: "",
	}
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
