package resources

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mshakery/ServerlessController/protos"
)

type PodResource struct {
	Pod protos.Pod
}

func (pr *PodResource) GetKVs() map[string]proto.Message {
	KVs := make(map[string]proto.Message)
	key := fmt.Sprintf("/cluster/resources/pod/%s/%s", pr.Pod.Metadata.Namespace, pr.Pod.Metadata.Name)
	KVs[key] = &pr.Pod
	key = fmt.Sprintf("/cluster/resources/pod/%s/%s/usage", pr.Pod.Metadata.Namespace, pr.Pod.Metadata.Name)
	KVs[key] = pr.Pod.GetStatus().GetResourceUsage()
	key = fmt.Sprintf("/cluster/resources/pod/%s/%s/worker", pr.Pod.Metadata.Namespace, pr.Pod.Metadata.Name)
	KVs[key] = pr.Pod.GetStatus().GetWorker()
	return KVs
}

func (pr *PodResource) CreatePostHook(ctx context.Context) bool {
	return true
	// TODO: call scheduler
}
