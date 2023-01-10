package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tangrc99/gohelloblog/docs"
	"github.com/tangrc99/gohelloblog/global"
	"github.com/tangrc99/gohelloblog/internal/middleware"
	"github.com/tangrc99/gohelloblog/internal/router/handler"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
)

func AddHandler(r *gin.Engine, routePath string, name string, handler handler.RestFulHandler) *gin.RouterGroup {
	group := r.Group(routePath)
	group.POST(name, handler.Create)
	group.DELETE(name, handler.Delete)
	group.PUT(name, handler.Update)
	group.GET(name, handler.Get)
	return group
}

func New() *gin.Engine {
	r := gin.New()

	//r.SetTrustedProxies()		这里需要添加信任的代理，目前打算是前面使用 nginx 作为代理

	// 这里添加 gin 的中间件，每一个包被提高到 router 之前都会在这个调用链中传递
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.GetTranslator()) // 中间件增加翻译选项
	//r.Use(middleware.SendAccessLog()) // 中间件加入 access log选项
	//r.Use(middleware.Timeout(global.ServerSetting.HandlerTimeout)) // 中间件加入超时选项，这里认证界面是不能加超时的
	r.LoadHTMLGlob("/Users/tangrenchu/GolandProjects/Blog/html/*") // 载入所有的 HTML 模板

	// 以下为 gin 的路由表，通过解析 url 将包转发给不同的处理器
	r.GET("/", handler.ShowIndexPage)
	r.GET("/login", handler.ShowLoginPage)
	r.POST("/auth", handler.HandleAuth)

	r.GET("/register", handler.ShowRegisterPage) // GET 请求返回注册页面
	r.POST("/register", handler.HandleRegister)  // POST 请求返回登录函数
	r.GET("/upload", handler.ShowUploadPage)     // GET 请求返回上传文件界面

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.Use(middleware.Timeout(global.ServerSetting.HandlerTimeout))
		api.Use(middleware.JWTAuth()) // 要求 api 路径下的所有请求需要进行授权
		var tag = handler.NewTag()
		api.POST("/tags", tag.Create)
		api.DELETE("/tags/:id", tag.Delete)
		api.PUT("/tags/:id", tag.Update)
		api.GET("/tags/:id", tag.Get)
		api.GET("/tags", tag.Get)

		var article = handler.NewArticle()

		api.POST("/articles", article.Create)
		api.GET("/articles/:id", article.Get)

		// 用来处理留言功能
		var note = handler.Note{}
		api.POST("/notes", note.Create)
		api.GET("/notes", note.List)
		api.PUT("/notes/:id", note.Update)
		api.GET("/notes/:id", note.Get)
		api.DELETE("/notes/:id", note.Delete)

		var file = handler.File{}

		api.POST("/files", file.Create)
	}

	r.NoRoute(func(ctx *gin.Context) {
		res := resp.NewResponse(ctx)
		res.SendError(err_code.NotFound.WithDetails("资源路径无法匹配"))
	})

	return r
}
