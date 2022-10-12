package user

import (
	"fmt"
	"strings"

	errorModels "employee_manage/constant"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateStructUser(user NewUser) (errValidate []errorModels.ErrorValidate) {
	validate = validator.New()
	err := validate.Struct(user)

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

func validateStructChangePassword(rPassword RequestChangePassword) (errValidate []errorModels.ErrorValidate) {
	validate = validator.New()
	err := validate.Struct(rPassword)

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

func updateValidation(request map[string]interface{}) (errValidate []errorModels.ErrorValidate) {

	for k, v := range request {
		if v == "" {
			errValidate = append(errValidate, errorModels.ErrorValidate{
				Field:   strings.ToLower(k),
				Message: fmt.Sprintf("%s cannot be empty", strings.ToLower(k)),
			})
		}
	}

	validationMap := map[string]string{
		"employee_code": "omitempty,gt=3,lt=100",
		"full_name":     "omitempty,gt=3,lt=100",
		"phone":         "omitempty,gt=3,lt=100",
		"email":         "omitempty,gt=3,lt=100",
		"gender":        "boolean",
		"birthday":      "datetime",
		"address":       "gt=3",
	}

	validate := validator.New()
	err := validate.RegisterValidation("update_validation", func(fl validator.FieldLevel) bool {
		m, ok := fl.Field().Interface().(map[string]interface{})
		if !ok {
			return false
		}

		for k, v := range validationMap {
			errVar := validate.Var(m[k], v)
			if errVar != nil {
				validatorErr := errVar.(validator.ValidationErrors)
				errValidate = append(errValidate, errorModels.ErrorValidate{
					Field:   k,
					Message: fmt.Sprintf("%s a is %s", k, validatorErr[0].Tag()),
				})
			}
		}

		return true
	})

	if err != nil {
		err = errorModels.NewAppError(err, errorModels.UnknownError)
		return
	}

	err = validate.Var(request, "update_validation")

	if err != nil {
		err = errorModels.NewAppError(err, errorModels.UnknownError)
		return
	}

	if errValidate != nil {
		return
	}

	return nil
}
