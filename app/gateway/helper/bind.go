package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func BindAndValidate(c *gin.Context, target any) error {
	_ = c.Bind(target)
	
	vzh := zh.New()
	uni := ut.New(vzh)
	translator, _ := uni.GetTranslator("zh")

	validate := validator.New()
	_ = zh_translations.RegisterDefaultTranslations(validate, translator)

	if err := validate.Struct(target); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// fmt.Println(e.Error()) // 原始错误信息
			// fmt.Println(e.Translate(translator)) // 翻译错误信息
			// fmt.Println(e.Namespace())
			// fmt.Println(e.Value())
			// fmt.Println(e.Param())
			// fmt.Println(e.Field())
			// fmt.Println(e.StructNamespace())
			// fmt.Println(e.StructField())
			// fmt.Println(e.Tag())
			// fmt.Println(e.ActualTag())
			// fmt.Println(e.Kind())
			// fmt.Println(e.Type())
			return errors.New(e.Translate(translator))
		}
	}
	
	return nil
}


