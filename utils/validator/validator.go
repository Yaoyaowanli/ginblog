package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)


//Validate  验证数据合法性，string中记录了翻译后的错误信息（空串为无错误），int为错误码
func Validate (data interface{})(string,int){
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans,_ := uni.GetTranslator("zh_Hans_CN")


	err := zhTrans.RegisterDefaultTranslations(validate,trans)
	if err != nil {
		fmt.Println("err:",err)
	}

	//拿到模型的label翻译
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	//依据结构体tag中validate标签验证公开字段是否合法，
	err = validate.Struct(data)
	if err != nil {
		//对于传入的错误值，它返回InvalidValidationError，否则返回nil或ValidationErrors作为错误。如果不是零，则需要断言错误
		//例如err。（validator.ValidationErrors）访问错误数组。
		for _,v := range err.(validator.ValidationErrors){
			//Translate返回翻译过后的错误信息
			return v.Translate(trans),errmsg.ERROR
		}
	}
	return "",errmsg.SUCCESS
}