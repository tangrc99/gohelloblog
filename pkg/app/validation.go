package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(ctx *gin.Context, obj interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := ctx.ShouldBind(obj) // 检查请求是否符合要求，但是不会立刻返回 http:400

	if err != nil {
		v := ctx.Value("translator") // 获得翻译后的错误信息
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs // 如果 validator 发生错误，直接返回
		}
		// 将翻译器翻译后的错误收集起来
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{Key: key, Message: value})
		}

		return false, errs
	}

	return true, nil
}
