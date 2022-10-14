package user

import (
	"fmt"
	"regexp"
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

func IsRFC3339Date(fl validator.FieldLevel) bool {
	RFC3339DateRegexString := "^((?:(\\d{4}-\\d{2}-\\d{2})T(\\d{2}:\\d{2}:\\d{2}(?:\\.\\d+)?))(Z|[\\+-]\\d{2}:\\d{2})?)$"
	RFC3339DateRegex := regexp.MustCompile(RFC3339DateRegexString)

	return RFC3339DateRegex.MatchString(fl.Field().String())
}

func updateValidation(request map[string]interface{}) (errValidate []errorModels.ErrorValidate) {
	validationMap := map[string]string{
		"full_name": "required,gt=3,lt=100",
		"phone":     "required,gt=3,lt=100",
		"email":     "required,email,gt=3,lt=100",
		"gender":    "required,boolean",
		"birthday":  "required,RFC3339Date",
		"address":   "required,gt=3",
	}

	validate := validator.New()
	validate.RegisterValidation("RFC3339Date", IsRFC3339Date)
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
