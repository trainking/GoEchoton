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

type Hauthorized_OP interface {
	Save(username, t string) error
	Check(t string) bool
}

type hauthorized_op struct {
	mongoClient *mongo.Client
}

// 保存
func (op *hauthorized_op) Save(username, t string) error {
	collection := op.mongoClient.Database("local").Collection("hauthorized")
	h := model.Hauthorized{Username: username, Tott: crc32.ChecksumIEEE([]byte(t)), Date: time.Now()}
	_, err := collection.InsertOne(context.TODO(), h)
	if err != nil {
		return err
	}
	return nil
}

// 检查
func (op *hauthorized_op) Check(t string) bool {
	var result model.Hauthorized
	collection := op.mongoClient.Database("local").Collection("hauthorized")
	filter := bson.D{primitive.E{Key: "tott", Value: crc32.ChecksumIEEE([]byte(t))}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	if result == (model.Hauthorized{}) {
		return false
	}
	return true
}

func NewHauthorizedOP() Hauthorized_OP {
	return &hauthorized_op{
		mongoClient: database.MongoClient,
	}
}
