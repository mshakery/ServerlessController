package resources

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mshakery/ServerlessController/protos"
)

type NodeResource struct {
	node protos.Node
}

func (nr *NodeResource) GetKVs() map[string]proto.Message {
	KVs := make(map[string]proto.Message)
	key := fmt.Sprintf("/cluster/resources/node/%s", nr.node.Metadata.Name)
	KVs[key] = &nr.node
	key = fmt.Sprintf("/cluster/resources/node/%s/allocatable", nr.node.Metadata.Name)
	KVs[key] = nr.node.GetStatus().GetCapacity()
	key = fmt.Sprintf("/cluster/resources/node/%s/condition", nr.node.Metadata.Name)
	KVs[key] = nr.node.GetStatus().GetCondition()
	key = fmt.Sprintf("/cluster/resources/node/%s/pods", nr.node.Metadata.Name)
	KVs[key] = nr.node.GetStatus().GetPods()
	return KVs
}

func (nr *NodeResource) CreatePostHook(ctx context.Context) bool {
	return true
}
