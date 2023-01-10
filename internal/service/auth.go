package service

import (
	"github.com/tangrc99/gohelloblog/internal/model"
	"github.com/tangrc99/gohelloblog/pkg/app"
)

type AuthRequest struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterRequest struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
	Again    string `form:"again" binding:"required"`
}

func (auth *AuthRequest) ToModel() model.User {
	return model.User{
		UserName: auth.User,
		Password: auth.Password,
	}
}

func (auth *RegisterRequest) ToModel() model.User {
	return model.User{
		UserName: auth.User,
		Password: auth.Password,
	}
}

func CheckAndAuth(auth *AuthRequest) (string, bool) {

	user := auth.ToModel()
	if !user.IsUsrPwdMatch() {
		return "用户或密码错误", false
	}

	token, err := app.GenerateToken(user.IsAdministrator())

	if err != nil {
		return "生成Token失败", false
	}

	return token, true
}

func CreateNewUser(req *RegisterRequest) error {
	user := req.ToModel()

	if err := user.CreateNewUser(); err != nil {
		return err
	}
	return nil
}
