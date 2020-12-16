package hauthorized

import (
	"context"
	"fmt"
	"hash/crc32"
	"time"

	. "GoEchoton/configs/api/conf"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Hauthorized struct {
	Username string
	Tott     uint32    // 不保存token本体，而是保存crc32校验码便于查询
	Date     time.Time // 通过索引设计超时删除
}

var client *mongo.Client

// 初始化创建连接
func init() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", Conf.Mongo.User, Conf.Mongo.Passwd, Conf.Mongo.Host, Conf.Mongo.Port)
	var clientOptions = options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
}

// 保存一个token标记
func Save(username string, t string) error {
	collection := client.Database("local").Collection("hauthorized")
	h := Hauthorized{username, crc32.ChecksumIEEE([]byte(t)), time.Now()}
	_, err := collection.InsertOne(context.TODO(), h)
	if err != nil {
		return err
	}
	return nil
}

// 煎炒token标记是否存在
func Check(t string) bool {

	var result Hauthorized
	collection := client.Database("local").Collection("hauthorized")
	filter := bson.D{primitive.E{Key: "tott", Value: crc32.ChecksumIEEE([]byte(t))}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return false
	}
	if result == (Hauthorized{}) {
		return false
	}
	return true
}
