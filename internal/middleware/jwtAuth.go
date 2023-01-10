package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/pkg/app"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
)

// JWTAuth checks cli's request. This function will reject all requests without "token" in cookies. If the request has
// this field, the permission level will be verified.
func JWTAuth() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if ctx.Request.RequestURI == "/auth" || ctx.Request.RequestURI == "/login" {
			return
		}

		var token string

		// 判断 cookie 和 header 中是否具有 token
		if s, err := ctx.Cookie("token"); err == nil {
			token = s
		} else {
			token = ctx.GetHeader("token")
		}

		if token == "" {

			res := resp.NewAbortError(ctx)
			res.SendError(err_code.UnauthorizedAuthNotExist.WithDetails("请先进入登录页面进行登录"))

		} else {
			claim, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					res := resp.NewAbortError(ctx)
					res.SendError(err_code.UnauthorizedTokenTimeout)
				default:
					res := resp.NewAbortError(ctx)
					res.SendError(err_code.UnauthorizedTokenError)
				}
			}

			// 验证后的信息，写入上下文中
			ctx.Set("is_admin", claim.IsAdmin)

			// 目前只允许管理员权限进行资源的删除
			if ctx.Request.Method == "DELETE" && !claim.IsAdmin {
				res := resp.NewAbortError(ctx)
				res.SendError(err_code.PermissionDenied.WithDetails("非管理员权限不可删除资源"))
			}
		}

	}
}
