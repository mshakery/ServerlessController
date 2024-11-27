package kubelet

import (
	"context"
	"flag"
	protos "github.com/mshakery/ServerlessController/protos"
)

var (
	port = flag.Int("port", 50060, "Metric Port")
)

type server struct {
	protos.UnimplementedKubeletServer
}

var podsToRun []protos.Pod

func (s *server) runAPod(ctx context.Context, pod *protos.Pod) (protos.Empty, error) {
	// command
	podsToRun = append(podsToRun, *pod)
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
