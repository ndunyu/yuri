package yuri

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func PhoneNumberValidation(fl validator.FieldLevel) bool {

	return checkKenyaInternationalPhoneNumber(fl.Field().String())
}





func TranslateErrors(trans ut.Translator, err error) []Field {

	errs := err.(validator.ValidationErrors)
	em := []Field{}
	for _, e := range errs {
		// can translate each error one at a time.
        f:=Field{
			//Field:   m,
			Message:e.Translate(trans) ,
		}


		em = append(em, f)

		//fmt.Println(e.Translate(trans))
	}

	return em

}

type Field struct {
	///Field int `json:"field"`
	Message string `json:"message"`
}