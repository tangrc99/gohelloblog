package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
	"net/http"
)

type Redirect struct {
	url string
}

func NewRedirectHandler(url string) Redirect {
	return Redirect{url: url}
}

func (r *Redirect) Create(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	res.SendError(err_code.MethodNotAllowed.WithDetails("应当使用 GET 方法请求"))
}

func (r *Redirect) Delete(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	res.SendError(err_code.MethodNotAllowed.WithDetails("应当使用 GET 方法请求"))
}

func (r *Redirect) Update(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	res.SendError(err_code.MethodNotAllowed.WithDetails("应当使用 GET 方法请求"))
}

func (r *Redirect) Get(ctx *gin.Context) {
	r.redirect(ctx)
}

func (r *Redirect) List(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	res.SendError(err_code.MethodNotAllowed.WithDetails("应当使用 GET 方法请求"))
}

// redirect 会将用户静态资源请求转发到nginx上，实现动静分离
func (r *Redirect) redirect(ctx *gin.Context) {
	ctx.Param("path")
	ctx.Redirect(http.StatusSeeOther, r.url) // 这里表明请求的是静态资源，必须使用 GET 方法来获取
}
