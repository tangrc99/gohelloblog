package global

import (
	"github.com/tangrc99/gohelloblog/pkg/db"
	"github.com/tangrc99/gohelloblog/pkg/setting"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	_            *log.Logger
	Mongo        *db.Mongo
	MySQL        *db.MySQL
	MongoLog     *mongo.Collection // 用于将 access log 记录到 mongodb 中
	MongoArticle *mongo.Collection
)

var (
	MongoSetting  *setting.MongoSetting
	ServerSetting *setting.ServerSetting
	JWTSetting    *setting.JWTSetting
	MySQLSetting  *setting.MySQLSetting
)
