package kubelet

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
	port = flag.Int("port", 50060, "Metric Port")
)

type Pod struct {
	protos.UnimplementedPod
}

var list[] pods_to_run;

func(s *server) run_a_pod(ctx context.Context,  pod *protos.Pod) (Empty, error) {
	// command
	pods_to_run = append(pods_to_run, pod.name);
	return protos.Empty{};
}

func(s *server) Metrics(ctx context.Context,  pod *protos.Empty) (*protos.NodeMetrics, error) {
	
	return protos.PodMetrics {
		name: "1",
		cpu_usage: "",
		memory_usage: ""
	}, nil;
}

