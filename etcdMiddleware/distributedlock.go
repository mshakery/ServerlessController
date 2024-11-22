package etcdMiddleware

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func AcquireLock(cli *clientv3.Client, lockKey string, ttl int64, waitTillRelease bool) (*clientv3.LeaseGrantResponse, error) {
	leaseResp, err := cli.Grant(context.TODO(), ttl)
	if err != nil {
		return nil, fmt.Errorf("failed to create lease: %v", err)
	}

	txn := cli.Txn(context.TODO())
	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "locked", clientv3.WithLease(leaseResp.ID))).
		Else()

	txnResp, err := txn.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit txn: %v", err)
	}

	if !txnResp.Succeeded {
		if waitTillRelease {
			WatchLock(cli, lockKey)
			return AcquireLock(cli, lockKey, ttl, false)
		} else {
			return nil, fmt.Errorf("lock exists from before")
		}
	}

	return leaseResp, nil
}

func ReleaseLock(client *clientv3.Client, leaseID clientv3.LeaseID) error {
	_, err := client.Revoke(context.TODO(), leaseID)
	if err != nil {
		return fmt.Errorf("failed to revoke lease: %v", err)
	}
	return nil
}

func WatchLock(client *clientv3.Client, lockKey string) {
	rch := client.Watch(context.TODO(), lockKey)

	for wresp := range rch {
		for _, ev := range wresp.Events {
			if ev.Type == clientv3.EventTypeDelete {
				return
			}
		}
	}
}
