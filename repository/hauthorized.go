package repository

import (
	"GoEchoton/database"
	"GoEchoton/model"
	"context"
	"hash/crc32"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Hauthorized_OP 验证操作
type Hauthorized_OP interface {
	Save(username, t string) error
	Check(t string) bool
}

// hauthorized_op op实现
type hauthorized_op struct {
	collection *mongo.Collection
}

// Save 保存数据
func (op *hauthorized_op) Save(username, t string) error {
	// collection := op.mongoClient.Database("local").Collection("hauthorized")
	h := model.Hauthorized{Username: username, Tott: crc32.ChecksumIEEE([]byte(t)), Date: time.Now()}
	_, err := op.collection.InsertOne(context.TODO(), h)
	if err != nil {
		return err
	}
	return nil
}

// Check 检查
func (op *hauthorized_op) Check(t string) bool {
	var result model.Hauthorized
	// collection := op.mongoClient.Database("local").Collection("hauthorized")
	filter := bson.D{primitive.E{Key: "tott", Value: crc32.ChecksumIEEE([]byte(t))}}
	err := op.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	if result == (model.Hauthorized{}) {
		return false
	}
	return true
}

// NewHauthorizedOP 创建操作接口
func NewHauthorizedOP() (Hauthorized_OP, error) {
	m, err := database.NewMongoDB()
	if err != nil {
		return nil, err
	}
	return &hauthorized_op{
		collection: m.GetCollection(model.HauthorizedDatabase, model.HauthorizedCollection),
	}, nil
}
