package handler

import "github.com/gin-gonic/gin"

type RestFulHandler interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
	Get(ctx *gin.Context)
	List(ctx *gin.Context)
}
