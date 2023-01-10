package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ArticleId struct {
	ID primitive.ObjectID
}

func NewArticleId(id string) *ArticleId {
	var objectIDFromHex = func(hex string) primitive.ObjectID {
		objectID, err := primitive.ObjectIDFromHex(hex)
		if err != nil {
			panic(err)
		}
		return objectID
	}
	return &ArticleId{ID: objectIDFromHex(id)}
}

type Article struct {
	Title    string    `bson:"title"`    // 文章名称
	Author   string    `bson:"author"`   // 文章作者
	CTime    time.Time `bson:"CTime"`    // 文章创建时间
	RTime    time.Time `bson:"RTime"`    // 文件修改时间
	FileName string    `bson:"fileName"` // 文章对应文件名
	Content  []byte    `bson:"content"`  // 文章内容
}

func (article *Article) InsertTo(collection *mongo.Collection) string {
	doc, err := collection.InsertOne(context.TODO(), article)
	if err != nil {
		return ""
	}

	return fmt.Sprint(doc.InsertedID)
}

func (id *ArticleId) GetArticle(collection *mongo.Collection) (Article, error) {

	docs, err := collection.Aggregate(context.TODO(), bson.A{
		bson.D{{"$match", bson.D{{"_id", id.ID}}}},
		bson.D{{"$unset", "_id"}},
		bson.D{{"$limit", 1}},
	})

	if err != nil {
		return Article{}, err
	}

	var results []Article

	err = docs.All(context.TODO(), &results)
	if err != nil {
		panic(err)
	}
	return results[0], nil
}
