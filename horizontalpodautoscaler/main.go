package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/uuid"
	etcd "github.com/mshakery/ServerlessController/etcdMiddleware"
	protos "github.com/mshakery/ServerlessController/protos"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	protos.UnimplementedHorizontalPodAutoscalerServer
}

func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}

func CalculateAverageResourceUsage(client *clientv3.Client, ctx context.Context, namespace string, podNames []string, resourceType string) float32 {
	sum := 0
	for _, podName := range podNames {
		key := fmt.Sprintf("/cluster/resources/pod/%s/%s/usage", namespace, podName)
		var ru protos.ResourceUsage
		etcd.ReadOneFromEtcdToPb(client, ctx, key, &ru)
		i, err := strconv.Atoi(ru.GetResourceUsage()[resourceType])
		if err != nil {
			// ... handle error
			panic(err)
		}
		sum += i
	}
	return float32(sum) / float32(len(podNames))
}

func (s *server) Scale(ctx context.Context, in *protos.HpaName) (*protos.Empty, error) {
	conn, err2 := grpc.NewClient("scheduler.default.10.101.174.165.sslip.io:80", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		log.Fatalf("scheduler node: could not connect: %v", err2)
	}
	_ = protos.NewSchedulerClient(conn) // todo <-- why is this not used?

	client, err := etcd.ConnectToEtcd()
	if err != nil {
		panic("Could not connect to etcd. ")
	}
	var hpa *protos.Hpa
	key := fmt.Sprintf("/cluster/resources/hpa/%s/%s", in.GetNamespace(), in.GetName())
	err = etcd.ReadOneFromEtcdToPb(client, ctx, key, hpa)
	if err != nil {
		return nil, err
	}
	deploymentKey := fmt.Sprintf("/cluster/resources/Deployment/%s/%s/status", in.GetNamespace(), hpa.GetSpec().GetTargetResource())
	var deployment *protos.DeploymentStatus
	err = etcd.ReadOneFromEtcdToPb(client, ctx, deploymentKey, deployment)
	if err != nil {
		return nil, err
	}
	resourceAverage := CalculateAverageResourceUsage(client, ctx, in.GetNamespace(), deployment.PodNames, hpa.GetSpec().GetMetrics().GetMetricType())
	targetValue := resourceAverage / hpa.GetSpec().GetMetrics().GetTargetValue()
	if Abs(targetValue/hpa.GetSpec().GetMetrics().GetTargetAverageUtilization()) <= 0.9 || Abs(targetValue/hpa.GetSpec().GetMetrics().GetTargetAverageUtilization()) > 1.1 {
		newReplica := int32(float32(deployment.GetReplicas()) * (targetValue / hpa.GetSpec().GetMetrics().GetTargetAverageUtilization()))
		currentReplica := deployment.Replicas
		currrentPods := deployment.PodNames
		if newReplica > deployment.Replicas {
			for i := int32(0); i < newReplica-deployment.Replicas; i++ {
				currentReplica += 1
				newUuid := uuid.New().String()
				newPodName := fmt.Sprintf("%s-%s", hpa.GetSpec().GetTargetResource(), newUuid)
				currrentPods = append(currrentPods, newPodName)
				err := etcd.WriteToEtcdFromPb(client, ctx, deploymentKey, &protos.DeploymentStatus{Replicas: currentReplica, PodNames: currrentPods})
				if err != nil {
					return nil, err
				}
			}
		} else {
			for i := int32(0); i < deployment.Replicas-newReplica; i++ {
				_ = currrentPods[len(currrentPods)-1]
				currentReplica -= 1
				currrentPods = currrentPods[:len(currrentPods)-1]
				err := etcd.WriteToEtcdFromPb(client, ctx, deploymentKey, &protos.DeploymentStatus{Replicas: currentReplica, PodNames: currrentPods})
				if err != nil {
					return nil, err
				}
			}
		}

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
	protos.RegisterHorizontalPodAutoscalerServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
