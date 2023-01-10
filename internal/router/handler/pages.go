package handler

import (
	"github.com/gin-gonic/gin"
)

// ShowIndexPage show users index page.
// @Summary 向用户展示文件上传界面
// @tags Index
// @Produce text/html
// @Success 200 "展示注册界面"
// @Router / [get]
func ShowIndexPage(ctx *gin.Context) {

	ctx.HTML(200, "index.tmpl",
		gin.H{"HEADLINE": "WELCOME",
			"DETAIL": "WELCOME TO GIN!",
			"list1":  "TinySwarm: 最主要的项目",
			"url1":   "http://192.168.1.106:8888/docs/1"})

}

// ShowUploadPage show users upload page.
// @Summary 向用户展示文件上传界面
// @tags File
// @Produce text/html
// @Success 200 "展示注册界面"
// @Router /upload [get]
func ShowUploadPage(ctx *gin.Context) {
	ctx.HTML(200, "upload.html", "")
}

// ShowLoginPage show users login page.
// @Summary 向用户展示注册界面
// @tags Auth
// @Produce text/html
// @Success 200 "展示注册界面"
// @Router /login [get]
func ShowLoginPage(ctx *gin.Context) {
	ctx.HTML(200, "login.html", "")
}

// ShowRegisterPage show users register page.
// @Summary 向用户展示注册界面
// @tags Auth
// @Produce text/html
// @Success 200 "展示登录界面"
// @Router /register [get]
func ShowRegisterPage(ctx *gin.Context) {
	ctx.HTML(200, "register.html", "")
}
