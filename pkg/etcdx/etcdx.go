package etcdx

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
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

// RemotePod 远程节点集合
type RemotePod struct {
	client   *clientv3.Client
	target   string
	onAddd   func(v string)
	onDelete func(v string)

	mu    sync.RWMutex
	nodes map[string]string
}

func NewRemotePod(target string, etcdGateway []string) (*RemotePod, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdGateway,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	r := &RemotePod{client: client, target: target}
	// 获取所有node
	if err := r.getNodes(); err != nil {
		return nil, err
	}
	// 监听变更
	go r.watchUpdate()
	return r, nil
}

// getNodes 获取所有节点
func (r *RemotePod) getNodes() error {
	kv := clientv3.NewKV(r.client)
	rangeRsp, err := kv.Get(context.TODO(), r.target, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	r.mu.Lock()
	for _, kv := range rangeRsp.Kvs {
		r.nodes[string(kv.Key)] = string(kv.Value)
	}
	r.mu.Unlock()
	return nil
}

// watchUpdate 监听节点变更
func (r *RemotePod) watchUpdate() {
	watcher := clientv3.NewWatcher(r.client)

	watchChan := watcher.Watch(context.TODO(), r.target, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			r.mu.Lock()
			switch event.Type {
			case mvccpb.PUT: // 有新节点加入
				v := string(event.Kv.Value)
				r.nodes[string(event.Kv.Key)] = v
				// 增加元素处理
				r.onAddd(v)
			case mvccpb.DELETE: // 有节点删除
				delete(r.nodes, string(event.Kv.Key))
				r.onDelete(string(event.Kv.Value))
			}
			r.mu.Unlock()
		}
	}
}

// GetNodes 获取所有节点
func (r *RemotePod) GetNodes() map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.nodes
}

// SetOnAdd 设置新增后处理
func (r *RemotePod) SetOnAdd(f func(v string)) {
	r.onAddd = f
}

// SetOnDelete 设置删除后处理
func (r *RemotePod) SetOnDelete(f func(v string)) {
	r.onDelete = f
}
