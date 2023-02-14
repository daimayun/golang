package gin

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
	zht "github.com/go-playground/validator/v10/translations/zh"
)

// Translator 翻译器
var Translator ut.Translator

// InitTrans 初始化翻译
func InitTrans(locale ...string) (err error) {
	l := "zh"
	if len(locale) > 0 {
		l = locale[0]
	}
	// 修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		Translator, ok = uni.GetTranslator(l)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", l)
		}
		switch l {
		case "en":
			err = ent.RegisterDefaultTranslations(v, Translator)
		case "zh":
			err = zht.RegisterDefaultTranslations(v, Translator)
		default:
			err = zht.RegisterDefaultTranslations(v, Translator)
		}
		return
	}
	return
}

// TransValidateError 翻译验证错误
func TransValidateError(err validator.ValidationErrors) map[string]string {
	return removeTopStruct(err.Translate(Translator))
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
