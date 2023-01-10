package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

// GetTranslator provides translation of errors.
func GetTranslator() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer c.Next() // 将上下文传递给下一个 middleware

		// 声明可以翻译的语言
		uni := ut.New(en.New(), zh.New())
		// 获得用户的地区
		locale := c.GetHeader("locale")
		// 根据用户地区选择翻译器
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zhTrans.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = enTrans.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zhTrans.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("translator", trans) // 将翻译器放入到上下文中
		}

	}

}
