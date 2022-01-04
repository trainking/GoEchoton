package etcdx

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// LeaseAndHeartbeat 创建租约并心跳续租
func LeaseAndHeartbeat(target string, value string, etcdGateway []string, leaseTTL int64, heartT int) error {
	target = strings.TrimRight(target, "/") + "/"

	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdGateway,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}

	kv := clientv3.NewKV(client)
	lease := clientv3.NewLease(client)
	var leaseID clientv3.LeaseID = 0

	go func() {
		defer client.Close()
		for {
			if leaseID == 0 {
				leaseResp, err := lease.Grant(context.TODO(), leaseTTL)
				if err != nil {
					panic(err)
				}
				key := target + fmt.Sprintf("%d", leaseResp.ID)
				if _, err := kv.Put(context.TODO(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
					panic(err)
				}
				leaseID = leaseResp.ID
			} else {
				if _, err := lease.KeepAliveOnce(context.TODO(), leaseID); err == rpctypes.ErrLeaseNotFound {
					leaseID = 0
					continue
				}
			}
			time.Sleep(time.Duration(heartT) * time.Second)
		}
	}()
	return nil
}
