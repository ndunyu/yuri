package yuri

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func PhoneNumberValidation(fl validator.FieldLevel) bool {
	return true
}

func TranslateErrors(trans ut.Translator, err error) map[int]string {

	errs := err.(validator.ValidationErrors)
	fmt.Println(errs.Translate(trans))
	em := make(map[int]string)
	for m, e := range errs {
		// can translate each error one at a time.
         em[m]=e.Translate(trans)
		///fmt.Println(e.Namespace())
		//fmt.Println(e.StructField())
		//fmt.Println(e.Field())
		//fmt.Println(e.Tag())
		///t = append(t, e.Translate(trans))

		//fmt.Println(e.Translate(trans))
	}

	return em

}
