package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetTracer() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 收到请求，trace 记录

		ctx.Next()

		// 完成请求，trace 记录

	}

}