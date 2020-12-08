package yuri

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func PhoneNumberValidation(fl validator.FieldLevel) bool {
	return true
}

func TranslateErrors(trans ut.Translator, err error) {
	errs := err.(validator.ValidationErrors)
	for m, e := range errs {
		fmt.Println(m)
		// can translate each error one at a time.
		fmt.Println(e.Translate(trans))

	}




}
