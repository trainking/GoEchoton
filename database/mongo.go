package database

import (
	. "GoEchoton/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", Config.Mongo.User, Config.Mongo.Passwd, Config.Mongo.Host, Config.Mongo.Port)
	var clientOptions = options.Client().ApplyURI(uri)
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
}
