package db

import (
	"context"
	"github.com/tangrc99/gohelloblog/pkg/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Mongo struct {
	client *mongo.Client // 连接实例
	db     *mongo.Database
}

func NewMongoFrom(setting *setting.MongoSetting) *Mongo {
	return NewMongo(setting.Url, setting.Db)
}

func NewMongo(mongoUrl, database string) *Mongo {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		panic(err)
	}

	db := client.Database(database)

	return &Mongo{client: client, db: db}
}

var collections = map[string]*mongo.Collection{}

func (db *Mongo) GetCollection(name string) *mongo.Collection {

	if collections[name] == nil {
		coll := db.db.Collection(name)
		collections[name] = coll

	}

	return collections[name]
}
