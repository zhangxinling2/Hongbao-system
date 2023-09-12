package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	vtzh "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User struct {
	FirstName string `validate:"required"` //非空字段
	LastName  string `validate:"required"`
	Age       uint8  `validate:"gte=0,lte=130"`  //大于等于0，小于等于130
	Email     string `validate:"required,email"` //非空且email字段
}

func main() {
	user := &User{
		FirstName: "firstName",
		LastName:  "lastName",
		Age:       136,
		Email:     "fl163.com",
	}
	validate := validator.New()
	//翻译器
	zh_cn := zh.New()
	uni := ut.New(zh_cn, zh_cn)
	Trans, found := uni.GetTranslator("zh")
	if found {
		err := vtzh.RegisterDefaultTranslations(validate, Trans)
		if err != nil {
			fmt.Println(err)
		}
	}
	err := validate.Struct(user)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			fmt.Println(err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, err := range errs {
				fmt.Println(err.Translate(Trans))
			}
		}
	}
}
