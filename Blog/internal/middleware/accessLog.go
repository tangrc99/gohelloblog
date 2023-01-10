package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/model"
)

// SendAccessLog collects information and insert access log into mongodb
func SendAccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		log := model.NewAccessLog()
		log.Host = ctx.Request.Host      // 访问本机所使用的host
		log.ClientIP = ctx.ClientIP()    // client ip
		log.Url = ctx.Request.RequestURI // uri
		log.Method = ctx.Request.Method  // http method

		ctx.Next()

		log.Status = ctx.Writer.Status() // http status

		go log.InsertTo(global.MongoLog) // 异步插入 mongodb，不过多影响性能

	}
}
