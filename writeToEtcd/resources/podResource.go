package resources

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mshakery/ServerlessController/protos"
)

type PodResource struct {
	pod protos.Pod
}

func (pr *PodResource) GetKVs() map[string]proto.Message {
	KVs := make(map[string]proto.Message)
	key := fmt.Sprintf("/cluster/resources/pod/%s/%s", pr.pod.Metadata.Namespace, pr.pod.Metadata.Name)
	KVs[key] = &pr.pod
	key = fmt.Sprintf("/cluster/resources/pod/%s/%s/usage", pr.pod.Metadata.Namespace, pr.pod.Metadata.Name)
	KVs[key] = pr.pod.GetStatus().GetResourceUsage()
	key = fmt.Sprintf("/cluster/resources/pod/%s/%s/worker", pr.pod.Metadata.Namespace, pr.pod.Metadata.Name)
	KVs[key] = pr.pod.GetStatus().GetWorker()
	return KVs
}

func (pr *PodResource) CreatePostHook(ctx context.Context) bool {
	return true
	// TODO: call scheduler
}
