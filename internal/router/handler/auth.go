package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tangrc99/gohelloblog/internal/service"
	"github.com/tangrc99/gohelloblog/pkg/app"
	"github.com/tangrc99/gohelloblog/pkg/err_code"
	"github.com/tangrc99/gohelloblog/pkg/resp"
	"net/http"
)

// HandleAuth handles user's login and authority request.
// @Summary 根据用户所填写的信息来对api接口进行授权
// @tags Auth
// @Accept application/x-www-form-urlencoded,application/json
// @Produce json
// @Param user body string true "用户名称" maxlength(20)
// @Param password body string true "用户密码" maxlength(20)
// @Success 302
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error
// @Router /auth [post]
func HandleAuth(ctx *gin.Context) {

	res := resp.NewResponse(ctx)
	param := service.AuthRequest{}

	valid, errs := app.BindAndValid(ctx, &param)
	if !valid {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	token, ok := service.CheckAndAuth(&param)
	if !ok {
		res.SendError(err_code.UnauthorizedTokenGenerate.WithDetails(token))
		return
	}

	ctx.SetCookie("token", token, 3600, "/", "", false, true)

	ctx.Redirect(http.StatusFound, "/")
}

// HandleRegister handles user's register request.
// @Summary 根据用户所填写的信息进行注册操作
// @tags Auth
// @Accept application/x-www-form-urlencoded,application/json
// @Produce json
// @Param user body string true "用户名称" maxlength(20)
// @Param password body string true "用户密码" maxlength(20)
// @Success 302
// @Failure 400 {object} err_code.Error
// @Failure 401 {object} err_code.Error
// @Router /register [post]
func HandleRegister(ctx *gin.Context) {
	res := resp.NewResponse(ctx)
	req := service.RegisterRequest{}

	valid, errs := app.BindAndValid(ctx, &req)
	if !valid {
		res.SendError(err_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	if req.Password != req.Again {
		res.SendError(err_code.InvalidParams.WithDetails("两次密码输入不一致"))
		return
	}

	err := service.CreateNewUser(&req)
	if err != nil {
		res.SendError(err_code.ServerError.WithDetails(err.Error()))
		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}
