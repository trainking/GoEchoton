package database

import (
	. "GoEchoton/config"
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB mongoDB结构体
type MongoDB struct {
	uri string
	mu  sync.Mutex
}

var (
	initialized uint32
	instance    *mongo.Client
)

// GetCollection 获取集合
func (m *MongoDB) GetCollection(database, collectionName string) (*mongo.Collection, error) {
	client, err := m.getClient()
	if err != nil {
		return nil, err
	}
	return client.Database(database).Collection(collectionName), nil
}

// getClient 获取client
func (m *MongoDB) getClient() (*mongo.Client, error) {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance, nil
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	var clientOptions = options.Client().ApplyURI(m.uri)
	var err error
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	atomic.StoreUint32(&initialized, 1)
	instance = c
	return instance, nil
}

// Destory 重置
func (m *MongoDB) Destory() {
	m.mu.Lock()
	defer m.mu.Unlock()
	instance = nil
	atomic.StoreUint32(&initialized, 0)
}

// NewMongoDB 新获取一个MongoDB连接
func NewMongoDB() *MongoDB {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", Config.Mongo.User, Config.Mongo.Passwd, Config.Mongo.Host, Config.Mongo.Port)

	return &MongoDB{uri: uri}
}
