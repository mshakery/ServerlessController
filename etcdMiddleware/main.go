package etcdMiddleware

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var endpoints = []string{"localhost:2379"}

func connectToEtcd() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("Failed to connect to etcd:", err)
		return nil, err
	}
	return cli, nil
	/* dont forget to:
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {
		}
	}(cli)
	*/
}

func writeToEtcd(cli *clientv3.Client, ctx context.Context, key string, val string) error {
	_, err := cli.Put(ctx, key, val)
	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Printf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			fmt.Printf("client-side error: %v\n", err)
		default:
			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
		return err
	}
	return nil
}

func readFromEtcd(cli *clientv3.Client, ctx context.Context, key string, val string) (*clientv3.GetResponse, error) {
	resp, err := cli.Get(ctx, key)
	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Printf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			fmt.Printf("client-side error: %v\n", err)
		default:
			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
		return nil, err
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
	return resp, nil
}
