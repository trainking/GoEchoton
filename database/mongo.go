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
	client *mongo.Client
}

var initialized uint32

var instance *MongoDB

var mu sync.Mutex

// GetCollection 获取集合
func (m *MongoDB) GetCollection(database, collectionName string) *mongo.Collection {
	return m.client.Database(database).Collection(collectionName)
}

// Destory 重置
func (m *MongoDB) Destory() {
	mu.Lock()
	defer mu.Unlock()
	instance = nil
	atomic.StoreUint32(&initialized, 0)
}

// NewMongoDB 新获取一个MongoDB连接
func NewMongoDB() (*MongoDB, error) {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance, nil
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", Config.Mongo.User, Config.Mongo.Passwd, Config.Mongo.Host, Config.Mongo.Port)
		var clientOptions = options.Client().ApplyURI(uri)
		var err error
		c, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return nil, err
		}
		instance = &MongoDB{client: c}
		atomic.StoreUint32(&initialized, 1)
	}
	return instance, nil
}
