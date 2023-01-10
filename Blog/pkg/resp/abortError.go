package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
)

type abortError struct {
	ctx *gin.Context
}

func NewAbortError(ctx *gin.Context) *abortError {
	return &abortError{ctx}
}

func (res *abortError) SendError(err *err_code.Error) {
	body := gin.H{"code": err.Code(), "msg": err.Message()}

	if len(err.Detail()) > 0 {
		body["detail"] = err.Detail()
	}

	res.ctx.AbortWithStatusJSON(err.HttpCode(), body)
}

func (res *abortError) Abort(err *err_code.Error) {
	res.ctx.AbortWithStatus(err.HttpCode())
}
