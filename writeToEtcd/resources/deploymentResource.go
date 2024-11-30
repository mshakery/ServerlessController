package resources

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/mshakery/ServerlessController/protos"
)

type DeploymentResource struct {
	Deployment protos.Deployment
}

func (dr *DeploymentResource) GetKVs() map[string]proto.Message {
	var podNames []string
	for i := int32(0); i < dr.Deployment.GetSpec().GetReplicas(); i++ {
		newUuid := uuid.New().String()
		newPodName := fmt.Sprintf("%s-%s", dr.Deployment.Metadata.Name, newUuid)
		podNames = append(podNames, newPodName)
	}
	dr.Deployment.Status.Replicas = dr.Deployment.GetSpec().GetReplicas()
	dr.Deployment.Status.PodNames = podNames
	KVs := make(map[string]proto.Message)
	key := fmt.Sprintf("/cluster/resources/Deployment/%s/%s", dr.Deployment.Metadata.Namespace, dr.Deployment.Metadata.Name)
	KVs[key] = &dr.Deployment
	key = fmt.Sprintf("/cluster/resources/Deployment/%s/%s/status", dr.Deployment.Metadata.Namespace, dr.Deployment.Metadata.Name)
	KVs[key] = dr.Deployment.GetStatus()
	return KVs
}

func (dr *DeploymentResource) CreatePostHook(ctx context.Context) bool {
	var pods []protos.Pod
	for _, podName := range dr.Deployment.GetStatus().GetPodNames() {
		newPod := protos.Pod{}
		newPod.Metadata = dr.Deployment.GetMetadata()
		newPod.Metadata.Name = podName
		newPod.Spec = dr.Deployment.GetSpec().GetTemplate().GetSpec()
		pods = append(pods, newPod)
	}
	result := true
	for _, pod := range pods {
		podResource := PodResource{pod: pod}                       //TODO: Debug this
		result = result && CreateResourceInEtcd(ctx, &podResource) //TODO: parallelize it
	}
	return result
}
