package app

import (
	"github.com/devhg/go-gin-demo/pkg/e"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

/**
封装github.com/go-playground/validator
翻译返回
*/

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	//注册翻译器
	zh := zh.New()
	uni = ut.New(zh, zh)

	trans, _ = uni.GetTranslator("zh")

	//获取gin的校验器
	//validate = binding.Validator.Engine().(*validator.Validate)
	validate = validator.New()
	//注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int, interface{}) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS,
			e.GetMsg(e.INVALID_PARAMS)
	}

	// 表单验证
	err = validate.Struct(form)

	if err != nil {
		errMsg := translate(err)
		return http.StatusBadRequest, e.INVALID_PARAMS, errMsg
	}
	return http.StatusOK, e.SUCCESS, ""
}

// translate errors to target language
func translate(err error) map[string][]string {
	errors := err.(validator.ValidationErrors)
	result := make(map[string][]string)
	for _, fieldError := range errors {
		result[fieldError.Field()] = append(result[fieldError.Field()], fieldError.Translate(trans))
	}
	return result
}
