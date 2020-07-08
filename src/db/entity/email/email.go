package email

import (
	"gopkg.in/go-playground/validator.v8"
)

type Email struct {
	 Email string `form:"email" json:"email" binding:"required,EmailValid"`
}

var EmailValid validator.Func
//=
//	func(fl validator.FieldLevel) bool {
//	return fl.Field().String() == "awesome"
//}

//func EmailValid(
//	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
//	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
//	) bool {
//	log.Fatal("???===========? email !!!W@#@#@", field)
//	if s, ok := field.Interface().(string); ok {
//		if s == "767838865@qq.com" {
//			return  false
//		}
//	}
//	return  true
//}