package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"net/http"
)

type response struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *response {
	return &response{ctx}
}

func (res *response) SendJson(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res.ctx.JSON(http.StatusOK, data)
}

func (res *response) SendWithType(data []byte) {

	res.ctx.Data(200, "text/html; charset=utf-8", data)
}

func (res *response) SendError(err *err_code.Error) {
	body := gin.H{"http_code": err.HttpCode(), "server_code": err.Code(), "msg": err.Message()}
	if len(err.Detail()) > 0 {
		body["detail"] = err.Detail()
	}

	res.ctx.HTML(err.HttpCode(), "error.tmpl", body)

	//res.ctx.JSON(err.HttpCode(), body)
}
