package resources

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mshakery/ServerlessController/protos"
)

type NodeResource struct {
	Node protos.Node
}

func (nr *NodeResource) GetKVs() map[string]proto.Message {
	KVs := make(map[string]proto.Message)
	key := fmt.Sprintf("/cluster/resources/node/%s", nr.Node.Metadata.Name)
	KVs[key] = &nr.Node
	key = fmt.Sprintf("/cluster/resources/node/%s/allocatable", nr.Node.Metadata.Name)
	KVs[key] = nr.Node.GetStatus().GetCapacity()
	key = fmt.Sprintf("/cluster/resources/node/%s/condition", nr.Node.Metadata.Name)
	KVs[key] = nr.Node.GetStatus().GetCondition()
	//key = fmt.Sprintf("/cluster/resources/node/%s/pods", nr.Node.Metadata.Name)
	//KVs[key] = nr.Node.GetStatus().GetPods()
	return KVs
}

func (nr *NodeResource) CreatePostHook(ctx context.Context) bool {
	return true
}
