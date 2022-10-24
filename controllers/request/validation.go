package request

import (
	errorModels "employee_manage/constant"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateStructRequest(request NewRequest) (errValidate []errorModels.ErrorValidate) {
	validate = validator.New()
	err := validate.Struct(request)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errValidate = append(errValidate, errorModels.ErrorValidate{
				Field:   err.StructField(),
				Message: fmt.Sprintf("%s a is %s", strings.ToLower(err.StructField()), err.Tag()),
			})
		}
	}

	return
}
