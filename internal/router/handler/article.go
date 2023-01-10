package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/model"
	"github.com/tangrc99/gohelloblog/internal/service"
	"github.com/tangrc99/gohelloblog/pkg/app"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
	"time"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (*Article) Create(ctx *gin.Context) {

	res := resp.NewResponse(ctx)

	// 数据验证阶段
	req := service.ArtReq{}
	ok, errs := app.BindAndValid(ctx, &req)

	if !ok {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 访问数据库进行插入

	doc := model.Article{
		Title:    req.Title,
		Author:   req.Author,
		CTime:    time.Now(),
		RTime:    time.Now(),
		FileName: req.FileName,
		Content:  req.Content,
	}

	id := doc.InsertTo(global.MongoArticle)

	res.SendJson(gin.H{"id": id})
}

func (*Article) Delete(ctx *gin.Context) {

}

func (*Article) Update(ctx *gin.Context) {

}

func (*Article) Get(ctx *gin.Context) {
	res := resp.NewResponse(ctx)

	id := model.NewArticleId(ctx.Param("id"))

	article, err := id.GetArticle(global.MongoArticle)
	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}

	res.SendWithType(article.Content)
}

func (*Article) List(ctx *gin.Context) {

}
