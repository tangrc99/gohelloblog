package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// AccessLog is the data model to be inserted into mongodb.
type AccessLog struct {
	Host     string    // 访问本机所使用的地址
	Url      string    // 访问资源位置
	Tp       time.Time // 访问时间
	Status   int       // 返回状态码
	Method   string    // 用户请求方式
	ClientIP string    // 用户 ip

}

// NewAccessLog create an AccessLog impl using current time.
func NewAccessLog() *AccessLog {
	log := AccessLog{}
	log.Tp = time.Now()
	return &log
}

// InsertTo inserts an AccessLog into mongodb.
func (log *AccessLog) InsertTo(collection *mongo.Collection) {
	_, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return
	}
}
