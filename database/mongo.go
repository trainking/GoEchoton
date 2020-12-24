package database

import (
	. "GoEchoton/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB mongoDB结构体
type MongoDB struct {
	client *mongo.Client
}

// GetCollection 获取集合
func (m *MongoDB) GetCollection(database, collectionName string) *mongo.Collection {
	return m.client.Database(database).Collection(collectionName)
}

// NewMongoDB 新获取一个MongoDB连接
func NewMongoDB() (*MongoDB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", Config.Mongo.User, Config.Mongo.Passwd, Config.Mongo.Host, Config.Mongo.Port)
	var clientOptions = options.Client().ApplyURI(uri)
	var err error
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return &MongoDB{client: c}, nil
}
