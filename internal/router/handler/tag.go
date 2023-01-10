package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/pkg/app"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
	"time"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

type Test struct {
	Name string `form:"name" binding:"max=20"`
}

func (tag Tag) Create(ctx *gin.Context) {
	res := resp.NewResponse(ctx)

	var test Test
	ok, errs := app.BindAndValid(ctx, &test)
	fmt.Printf(test.Name)
	if !ok {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	res.SendJson("OK")
}

func (tag Tag) Delete(ctx *gin.Context) {

}

func (tag Tag) Update(ctx *gin.Context) {

}

func (tag Tag) Get(ctx *gin.Context) {
	time.Sleep(5 * time.Second)

	ctx.HTML(200, "out.html.html", "dsfsdfdsfsd")

}

func (tag Tag) List(ctx *gin.Context) {

}
